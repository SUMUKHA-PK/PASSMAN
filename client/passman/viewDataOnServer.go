package passman

import (
	"fmt"

	"github.com/SUMUKHA-PK/PASSMAN/client/crypto"
	"github.com/gookit/color"
)

// ViewDataOnServer allows to view vault on server
func ViewDataOnServer() {
	color.Info.Printf("PASSMAN Server data view sequence\n\n")

	username, _, vaultPwd, err := verifyAndGetVaultData()
	if err != nil {
		color.Error.Println(err)
		return
	}

	authPwd := crypto.SHA256(username + vaultPwd)
	vaultServer, err := getDataFromServer(authPwd)
	if err != nil {
		color.Error.Println("No data available on server.")
	}

	decryptedVaultServer, err := decryptVault([]byte(vaultServer), vaultPwd)
	if err != nil {
		color.Error.Println(err)
		return
	}
	fmt.Println("Data on the server:")
	fmt.Println(decryptedVaultServer)

	fmt.Printf("\n\n")

	color.Success.Println("Data read from server.")
}
