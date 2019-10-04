package passman

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/SUMUKHA-PK/PASSMAN/client/crypto"
	"github.com/SUMUKHA-PK/PASSMAN/client/redis"
	"golang.org/x/crypto/ssh/terminal"
)

// Vault is the entire vault data that gets encrypted
type Vault struct {
	HostPwd   string
	TimeStamp time.Time
}

// verifyAndGetDecryptedVaultData is a helper function that
// checks auth and gets data from the REDIS server
func verifyAndGetDecryptedVaultData() (string, string, string, error) {
	username, vault, err := verifyAndGetVaultData()
	if err != nil {
		return "", "", "", err
	}

	fmt.Printf("\nYour vault password is: %s\n\n", vault.VaultPwd)

	decryptedVault := decryptVault([]byte(vault.Vault), vault.VaultPwd)

	return decryptedVault, username, vault.VaultPwd, err
}

func verifyAndGetVaultData() (string, redis.VaultData, error) {
	username, err := getUsername()
	if err != nil {
		fmt.Println(err)
		return "", redis.VaultData{}, err
	}
	vault, err := redis.Retrieve(username)
	if err != nil {
		fmt.Printf("You've not registered to PASSMAN! Please register by choosing option 1.\n\n")
		return "", redis.VaultData{}, err
	}
	fmt.Printf("Hello %s!\nPlease enter your master password: ", username)
	masterPwd, err := getMasterPwd()
	if err != nil {
		fmt.Println(err)
		return "", redis.VaultData{}, err
	}

	vaultPwd := crypto.SHA256(username + masterPwd)
	if vaultPwd != vault.VaultPwd {
		fmt.Println("You entered a wrong password! Please try again.")
		return "", redis.VaultData{}, err
	}
	return username, vault, nil
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
