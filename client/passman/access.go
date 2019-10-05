package passman

import (
	"fmt"

	"github.com/gookit/color"
)

// Access gets all the data from the vault of the current person
func Access() {
	color.Info.Printf("\nPASSMAN Vault view sequence.\n\n")
	vault, _, _, err := verifyAndGetDecryptedVaultData()
	if err != nil {
		return
	}
	fmt.Println("Your vault currently looks like this : ")
	fmt.Println(vault)
	fmt.Printf("\n\n")

	color.Success.Println("Vault access complete!")
}
