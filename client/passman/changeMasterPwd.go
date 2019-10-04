package passman

import (
	"fmt"
	"log"

	"github.com/SUMUKHA-PK/PASSMAN/client/crypto"
	"github.com/SUMUKHA-PK/PASSMAN/client/redis"
)

// ChangeMasterPwd allows the user to change the master password
// Following are the steps involved.
// 1. Authenticate the user as usual.
// 2. Ask for the new master password on authentication.
// 3. Decrypt the existing vault data using the old password.
// 4. Calculate the new vault password using the new master password
// 5. Encrypt the decrypted data using the new vault password.
// 6. Write all the data back to Redis.
func ChangeMasterPwd() {
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
