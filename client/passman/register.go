package passman

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/SUMUKHA-PK/Password-Manager/client/crypto"
	"github.com/SUMUKHA-PK/Password-Manager/client/redis"
)

// Register control flow is as follows:
// 1. Takes username from STDIN and checks
// for existance of given user.
// 2. On confirming that the user is new,
// proceeds to obtain master password from
// STDIN,creates a vault password.
// 3. Creates a dummy vault entry with the username
// of the user, encrypts using the vault password.
// 4. Registration complete!
func Register() {

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
