package passman

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/SUMUKHA-PK/PASSMAN/client/redis"
	"github.com/gookit/color"
)

// AddPwd enables the user to add more passwords to the existing vault
func AddPwd() {
	color.Info.Printf("\nPASSMAN Password addition sequence.\n\n")

	vault, username, vaultPwd, err := verifyAndGetDecryptedVaultData()
	if err != nil {
		return
	}

	var vaultMap map[string]Vault
	err = json.Unmarshal([]byte(vault), &vaultMap)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the host of the password: ")
	hostname, err := reader.ReadString('\n')
	if err != nil {
		color.Error.Printf("Error in reading input : %v", err)
	}
	hostname = strings.TrimSpace(hostname)

	fmt.Println("\nEnter the password of the host: ")
	password, err := getMasterPwd()
	if err != nil {
		color.Error.Println(err)
		return
	}

	vaultMap[hostname] = Vault{password, time.Now()}
	if _, ok := vaultMap[username]; ok {
		delete(vaultMap, username)
	}
	byteMap, err := json.Marshal(vaultMap)
	if err != nil {
		color.Error.Printf("Error in marshalling : %v", err)
	}

	byteEncryptedVault, err := encryptVault(byteMap, vaultPwd)
	if err != nil {
		return
	}

	err = redis.Update(username, string(byteEncryptedVault))
	if err != nil {
		color.Error.Printf("Can't add data to Redis DB: %v", err)
	}
	fmt.Printf("\n\n")

	color.Success.Println("Password addition complete!")

}
