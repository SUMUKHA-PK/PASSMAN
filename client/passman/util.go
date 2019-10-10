package passman

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/SUMUKHA-PK/PASSMAN/client/crypto"
	"github.com/SUMUKHA-PK/PASSMAN/client/redis"
	"github.com/gookit/color"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/ssh/terminal"
)

// Vault is the entire vault data that gets encrypted
type Vault struct {
	HostPwd   string
	TimeStamp time.Time
}

// verifyAndGetDecryptedVaultData is a helper function that
// checks auth and gets data from the REDIS server
func verifyAndGetDecryptedVaultData() (string, string, []byte, error) {
	username, vault, vaultPwd, err := verifyAndGetVaultData()
	if err != nil {
		return "", "", []byte{}, err
	}

	fmt.Printf("\nYour vault password is: %s\n\n", vaultPwd)

	decryptedVault, err := decryptVault([]byte(vault.Vault), vaultPwd)
	if err != nil {
		fmt.Println("Error in decrypting vault: %v", err)
	}
	return decryptedVault, username, vaultPwd, err
}

func verifyAndGetVaultData() (string, redis.VaultData, []byte, error) {
	username, err := getUsername()
	if err != nil {
		fmt.Println(err)
		return "", redis.VaultData{}, []byte{}, err
	}
	vault, err := redis.Retrieve(username)
	if err != nil {
		fmt.Printf("You've not registered to PASSMAN! Please register by choosing option 1.\n\n")
		return "", redis.VaultData{}, []byte{}, err
	}
	fmt.Printf("Hello %s!\nPlease enter your master password: ", username)
	masterPwd, err := getMasterPwd()
	if err != nil {
		fmt.Println(err)
		return "", redis.VaultData{}, []byte{}, err
	}

	vaultPwd := argon2.IDKey([]byte(masterPwd), []byte(username), 1, 64*1024, 4, 32)

	_, err = decryptVault([]byte(vault.Vault), vaultPwd)
	if err != nil {
		fmt.Printf("\n")
		color.Error.Println("You entered a wrong password! Please try again.")
		return "", redis.VaultData{}, []byte{}, err
	}

	return username, vault, vaultPwd, nil
}

// getUsername gets the username from the STDIN
func getUsername() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your email id: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error in reading input : %v", err)
		return "", err
	}
	username = strings.TrimSpace(username)
	return username, nil
}

// getMasterPwd gets the master password from STDIN
func getMasterPwd() (string, error) {

	masterPw1, err := terminal.ReadPassword(0)
	if err != nil {
		fmt.Printf("Error in reading input : %v", err)
		return "", err
	}

	fmt.Println("\nRe-enter password to confirm: ")
	masterPw2, err := terminal.ReadPassword(0)
	if err != nil {
		fmt.Printf("Error in reading input : %v", err)
		return "", err
	}

	if string(masterPw1) != string(masterPw2) {
		return "", errors.New("Passwords dont match!")
	}

	return string(masterPw1), nil
}

// encryptVault encrypts the vault and returns the encrypted string
func encryptVault(byteMap []byte, vaultPwd []byte) ([]byte, error) {

	// double encryption
	encryptedData, err := crypto.AESEncrypt(byteMap, vaultPwd)
	if err != nil {
		fmt.Printf("Error in encrypting : %v", err)
		return []byte{}, err
	}

	return encryptedData, nil
}

// decryptVault decrypts the encrypted Vault and returns the PT string
func decryptVault(encryptedData []byte, vaultPwd []byte) (string, error) {

	// equivalent decryption
	plainText, err := crypto.AESDecrypt(encryptedData, vaultPwd)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
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
