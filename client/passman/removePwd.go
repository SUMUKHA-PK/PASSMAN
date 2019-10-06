package passman

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/SUMUKHA-PK/PASSMAN/client/crypto"
	"github.com/SUMUKHA-PK/PASSMAN/client/redis"
	"github.com/gookit/color"
)

// RemovePwd removes passwords from the DB
// 1. Auth the user
// 2. Get the vault from the DB and decrypt the vault.
// 3. Unmarshal onto map and remove the entry from the map.
// 4. Marshall the vault and encrypt the data.
// 5. Add back to the DB.
func RemovePwd() {
	color.Info.Printf("\nPASSMAN Password removal sequence.\n\n")

	username, err := getUsername()
	if err != nil {
		color.Error.Println(err)
		return
	}
	vault, err := redis.Retrieve(username)
	if err != nil {
		color.Error.Printf("You've not registered to PASSMAN! Please register by choosing option 1.\n\n")
		return
	}
	fmt.Printf("Hello %s!\nPlease enter your master password: ", username)
	masterPwd, err := getMasterPwd()
	if err != nil {
		color.Error.Println(err)
		return
	}

	vaultPwd := crypto.SHA256(username + masterPwd)
	decryptedVault, err := decryptVault([]byte(vault.Vault), vaultPwd)
	if err != nil {
		color.Error.Println("You entered a wrong password! Please try again.")
		return
	}

	host, err := getHostToRemove()
	if err != nil {
		return
	}
	var vaultMap map[string]Vault
	err = json.Unmarshal([]byte(decryptedVault), &vaultMap)
	if err != nil {
		color.Error.Println(err)
		return
	}

	if _, ok := vaultMap[host]; ok {
		delete(vaultMap, host)
	} else {
		color.Error.Println("Host doesn't exist in your vault!")
		return
	}

	byteVault, err := json.Marshal(vaultMap)

	byteEncryptedVault, err := encryptVault(byteVault, vaultPwd)
	if err != nil {
		return
	}

	err = redis.Update(username, string(byteEncryptedVault))
	if err != nil {
		log.Fatalf("Can't add data to Redis DB: %v", err)
	}

	fmt.Printf("\n\n")

	color.Success.Println("Password removal complete!")
}

func getHostToRemove() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the host you want to remove: ")
	host, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error in reading input : %v", err)
		return "", err
	}
	host = strings.TrimSpace(host)

	return host, nil
}
