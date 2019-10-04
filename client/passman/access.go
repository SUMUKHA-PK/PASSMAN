package passman

import "fmt"

// Access gets all the data from the vault of the current person
func Access() {
	fmt.Printf("\nPASSMAN Vault view sequence.\n")
	vault, _, _, err := verifyAndGetDecryptedVaultData()
	if err != nil {
		return
	}
	fmt.Println("Your vault currently looks like this : ")
	fmt.Println(vault)
	fmt.Printf("\n\n")

}
