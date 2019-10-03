package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/SUMUKHA-PK/Password-Manager/client/redis"

	"github.com/SUMUKHA-PK/Password-Manager/client/crypto"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	welcome := `
|------ |----| ------ ------ |\    /| |----| |\    |
|     |	|    | |      |      | \  / | |    | | \   |
|------	|----| |----- |----- |  \/  | |----| |  \  |
|       |    |      |      | |      | |    | |   \ |
|       |    | -----| -----| |      | |    | |    \|

Welcome to PASSMAN, your locally hosted Password Manager!

`
	fmt.Println(welcome)

	for {
		fmt.Printf("\nWhat do you want to do?\n\n")
		fmt.Printf("1. Register on PASSMAN.\n2. View saved passwords.\n3. Add password.\n4. Change master password.\n5. Exit PASSMAN :(\n\n")
		reader := bufio.NewReader(os.Stdin)

		option, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error in reading input: %v", err)
		}
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			register()
		case "2":
			access()
		case "3":
			addPwd()
		case "4":
			changeMasterPwd()
		default:
			fmt.Println("Please enter a valid option in the given list!")
		case "5":
			break

		}

		if option == "5" {
			fmt.Println("Exiting PASSMAN. Bye :)")
			break
		}
	}

}

// register control flow is as follows:
// 1. Takes username from STDIN and checks
// for existance of given user.
// 2. On confirming that the user is new,
// proceeds to obtain master password from
// STDIN,creates a vault password.
// 3. Creates a dummy vault entry with the username
// of the user, encrypts using the vault password.
// 4. Registration complete!
func register() {

	fmt.Printf("\nPASSMAN Registration sequence.\n")
	username, err := getUsername()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = redis.Retrieve(username)
	if err == nil {
		fmt.Printf("You've already registered to PASSMAN! Try something else.\n\n")
		return
	}
	fmt.Printf("Hello %s!\nPlease enter your master password: ", username)
	masterPwd, err := getMasterPwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\nGenerating vault key....\n")

	vaultPwd := crypto.SHA256(username + masterPwd)
	fmt.Printf("Your vault password is: %s\n\n", vaultPwd)

	authPwd := crypto.SHA256(masterPwd + vaultPwd)
	fmt.Printf("Your auth password is: %s\n\n", authPwd)

	m := make(map[string]string)
	// A first dummy entry
	m[username] = username

	byteMap, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error in marshalling : %v", err)
	}

	byteEncryptedVault := encryptVault(byteMap, vaultPwd)

	err = redis.Update(username, vaultPwd, string(byteEncryptedVault))
	if err != nil {
		log.Fatalf("Can't add data to Redis DB: %v", err)
	}

	fmt.Println("Registration complete!")
}

// access gets all the data from the vault of the current person
func access() {
	fmt.Printf("\nPASSMAN Vault view sequence.\n")
	vault, _, _, err := verifyAndGetDecryptedVaultData()
	if err != nil {
		return
	}
	fmt.Println("Your vault currently looks like this : ")
	fmt.Println(vault)
	fmt.Printf("\n\n")

}

// addPwd enables the user to add more passwords to the existing vault
func addPwd() {
	fmt.Printf("\nPASSMAN Password addition sequence.\n")

	vault, username, vaultPwd, err := verifyAndGetDecryptedVaultData()
	if err != nil {
		return
	}

	var vaultMap map[string]string
	err = json.Unmarshal([]byte(vault), &vaultMap)
	fmt.Println(vaultMap)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the host of the password: ")
	hostname, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error in reading input : %v", err)
	}
	hostname = strings.TrimSpace(hostname)

	fmt.Println("\nEnter the password of the host: ")
	password, err := getMasterPwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	vaultMap[hostname] = password
	if _, ok := vaultMap[username]; ok {
		delete(vaultMap, username)
	}
	byteMap, err := json.Marshal(vaultMap)
	if err != nil {
		log.Fatalf("Error in marshalling : %v", err)
	}

	byteEncryptedVault := encryptVault(byteMap, vaultPwd)

	err = redis.Update(username, vaultPwd, string(byteEncryptedVault))
	if err != nil {
		log.Fatalf("Can't add data to Redis DB: %v", err)
	}

	fmt.Printf("\nPassword addition complete!\n\n")

}

// changeMasterPwd allows the user to change the master password
// Following are the steps involved.
// 1. Authenticate the user as usual.
// 2. Ask for the new master password on authentication.
// 3. Decrypt the existing vault data using the old password.
// 4. Calculate the new vault password using the new master password
// 5. Encrypt the decrypted data using the new vault password.
// 6. Write all the data back to Redis.
func changeMasterPwd() {
	fmt.Printf("\nPASSMAN Master Password changing sequence.\n")

	username, err := getUsername()
	if err != nil {
		fmt.Println(err)
		return
	}
	vault, err := redis.Retrieve(username)
	if err != nil {
		fmt.Printf("You've not registered to PASSMAN! Please register by choosing option 1.\n\n")
		return
	}
	fmt.Printf("Hello %s!\nPlease enter your master password: ", username)
	masterPwd, err := getMasterPwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	vaultPwd := crypto.SHA256(username + masterPwd)
	if vaultPwd != vault.VaultPwd {
		fmt.Println("You entered a wrong password! Please try again.")
		return
	}

	fmt.Println("\nEnter the new master password: ")
	masterPwd, err = getMasterPwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	decryptedVault := decryptVault([]byte(vault.Vault), vault.VaultPwd)

	vaultPwd = crypto.SHA256(username + masterPwd)
	fmt.Printf("Your vault password is: %s\n\n", vaultPwd)

	byteEncryptedVault := encryptVault([]byte(decryptedVault), vaultPwd)

	err = redis.Update(username, vaultPwd, string(byteEncryptedVault))
	if err != nil {
		log.Fatalf("Can't add data to Redis DB: %v", err)
	}
	fmt.Println("Password change complete!")
}

func verifyAndGetDecryptedVaultData() (string, string, string, error) {
	username, err := getUsername()
	if err != nil {
		fmt.Println(err)
		return "", "", "", err
	}
	vault, err := redis.Retrieve(username)
	if err != nil {
		fmt.Printf("You've not registered to PASSMAN! Please register by choosing option 1.\n\n")
		return "", "", "", err
	}
	fmt.Printf("Hello %s!\nPlease enter your master password: ", username)
	masterPwd, err := getMasterPwd()
	if err != nil {
		fmt.Println(err)
		return "", "", "", err
	}

	vaultPwd := crypto.SHA256(username + masterPwd)
	if vaultPwd != vault.VaultPwd {
		fmt.Println("You entered a wrong password! Please try again.")
		return "", "", "", err
	}

	fmt.Printf("\nYour vault password is: %s\n\n", vaultPwd)

	decryptedVault := decryptVault([]byte(vault.Vault), vault.VaultPwd)

	return decryptedVault, username, vaultPwd, err
}

// getUsername gets the username from the STDIN
func getUsername() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your email id: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error in reading input : %v", err)
		return "", err
	}
	username = strings.TrimSpace(username)
	return username, nil
}

// getMasterPwd gets the master password from STDIN
func getMasterPwd() (string, error) {

	masterPw1, err := terminal.ReadPassword(0)
	if err != nil {
		log.Fatalf("Error in reading input : %v", err)
	}

	fmt.Println("\nRe-enter password to confirm: ")
	masterPw2, err := terminal.ReadPassword(0)
	if err != nil {
		log.Fatalf("Error in reading input : %v", err)
		return "", err
	}

	if string(masterPw1) != string(masterPw2) {
		return "", errors.New("Passwords dont match!")
	}

	return string(masterPw1), nil
}

// encryptVault encrypts the vault and returns the encrypted string
func encryptVault(byteMap []byte, vaultPwd string) []byte {
	vaultPwdArr := splitString(vaultPwd)

	// double encryption
	encryptedData, err := crypto.AESEncrypt(byteMap, []byte(vaultPwdArr[0]))
	if err != nil {
		log.Fatalf("Error in encrypting : %v", err)
	}

	encryptedData, err = crypto.AESEncrypt(encryptedData, []byte(vaultPwdArr[1]))
	if err != nil {
		log.Fatalf("Error in encrypting : %v", err)
	}
	return encryptedData
}

// decryptVault decrypts the encrypted Vault and returns the PT string
func decryptVault(encryptedData []byte, vaultPwd string) string {
	vaultPwdArr := splitString(vaultPwd)

	// equivalent decryption
	plainText, err := crypto.AESDecrypt(encryptedData, []byte(vaultPwdArr[1]))
	if err != nil {
		log.Fatalf("Error in decrypting 1: %v", err)
	}

	plainText, err = crypto.AESDecrypt(plainText, []byte(vaultPwdArr[0]))
	if err != nil {
		log.Fatalf("Error in decrypting : %v", err)
	}

	return string(plainText)
}

// splitString is a utility function to split a string into 2 equal parts
func splitString(pwd string) []string {
	var arr []string
	arr = append(arr, "")
	arr = append(arr, "")
	for i := 0; i < len(pwd); i++ {
		if i < len(pwd)/2 {
			arr[0] += string(pwd[i])
		} else {
			arr[1] += string(pwd[i])
		}
	}
	return arr
}
