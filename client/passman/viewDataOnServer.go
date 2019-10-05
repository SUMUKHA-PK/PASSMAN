package passman

import (
	"fmt"

	"github.com/SUMUKHA-PK/PASSMAN/client/crypto"
)

// ViewDataOnServer allows to view vault on server
func ViewDataOnServer() {
	username, _, vaultPwd, err := verifyAndGetVaultData()
	if err != nil {
		fmt.Println(err)
		return
	}

	authPwd := crypto.SHA256(username + vaultPwd)
	vaultServer, err := getDataFromServer(authPwd)
	if err != nil {
		fmt.Println("No data available on server.")
	}

	fmt.Printf("%x", vaultServer)
	decryptedVaultServer, err := decryptVault([]byte(vaultServer), vaultPwd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data on the server:")
	fmt.Println(decryptedVaultServer)
}
