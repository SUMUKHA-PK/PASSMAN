package passman

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/SUMUKHA-PK/PASSMAN/client/crypto"
	"github.com/SUMUKHA-PK/PASSMAN/client/redis"
	"github.com/gookit/color"
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
// The vaultPwd is not stored in the DB as it
// might compromise the security. The correctness
// of decryption is handled by the message auth codes.
func Register() {

	color.Info.Printf("\nPASSMAN Registration sequence.\n\n")
	username, err := getUsername()
	if err != nil {
		color.Error.Printf("Can't read from STDIN: %v", err)
		return
	}
	_, err = redis.Retrieve(username)
	if err == nil {
		fmt.Printf("\n")
		color.Error.Printf("You've already registered to PASSMAN! Try something else.")
		return
	}
	fmt.Printf("Hello %s!\nPlease enter your master password: ", username)
	masterPwd, err := getMasterPwd()
	if err != nil {
		color.Error.Printf("Can't read from STDIN: %v", err)
		return
	}

	fmt.Printf("\nGenerating vault key....\n")

	vaultPwd := crypto.SHA256(username + masterPwd)
	fmt.Printf("\nYour vault password is: %s\n", vaultPwd)

	m := make(map[string]Vault)
	// A first dummy entry
	m[username] = Vault{username, time.Now()}

	byteMap, err := json.Marshal(m)
	if err != nil {
		color.Error.Printf("Error in marshalling the map: %v", err)
		return
	}

	byteEncryptedVault, err := encryptVault(byteMap, vaultPwd)
	if err != nil {
		return
	}
	err = redis.Update(username, string(byteEncryptedVault))
	if err != nil {
		color.Error.Printf("Can't add data to Redis DB: %v", err)
		return
	}
	fmt.Printf("\n")
	color.Success.Println("Registration complete!")
}
