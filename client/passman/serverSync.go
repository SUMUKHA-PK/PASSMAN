package passman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/SUMUKHA-PK/PASSMAN/client/crypto"
	"github.com/SUMUKHA-PK/PASSMAN/client/redis"
	"github.com/SUMUKHA-PK/PASSMAN/server/routing"
)

// SyncDataWithServer pushes all the data to the server
// It ensures no data is lost while syncing, by the given steps:
// 1. Auth with the users credentials.
// 2. Request the server for the vault data.
// 3. Decrypt both vaults.
// 4. Update the client vault, timestamp based.
// 5. Encrypt the vault, save the vault in REDIS and send to server.
func SyncDataWithServer() {
	fmt.Println("PASSMAN Server sync sequence")

	username, vault, err := verifyAndGetVaultData()
	if err != nil {
		fmt.Println(err)
		return
	}

	// send a req to a server and obtain the vault in the server
	//getDataFromServer

	authPass := crypto.SHA256(username + vault.VaultPwd)
	vaultServer := getDataFromServer(authPass)

	decryptedVault := decryptVault([]byte(vault.Vault), vault.VaultPwd)
	decryptedVaultServer := decryptVault([]byte(vaultServer), vault.VaultPwd)

	var vaultMap, vaultMapServer map[string]Vault
	err = json.Unmarshal([]byte(decryptedVault), &vaultMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal([]byte(decryptedVaultServer), &vaultMapServer)
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range vaultMapServer {
		// If the value mapping doesnt exist, add the data into client map
		if _, ok := vaultMap[k]; !ok {
			vaultMap[k] = v
		} else {
			// If the mapping exists, use the most updated data
			if vaultMap[k].TimeStamp.Before(vaultMapServer[k].TimeStamp) {
				vaultMap[k] = v
			}
		}
	}

	byteMap, err := json.Marshal(vaultMap)
	if err != nil {
		fmt.Println(err)
		return
	}

	byteEncryptedVault := encryptVault(byteMap, vault.VaultPwd)

	err = redis.Update(username, vault.VaultPwd, string(byteEncryptedVault))
	if err != nil {
		fmt.Printf("Can't add data to Redis DB: %v", err)
		return
	}

	// send the encryted vault to the server

}

func getDataFromServer(authPass string) string {

	URL := "127.0.0.1:6666"
	outData := &routing.GetDataReq{authPass}
	payload, err := json.Marshal(outData)
	if err != nil {
		log.Printf("Can't Marshall to JSON in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", URL, strings.NewReader(string(payload)))
	if err != nil {
		log.Printf("Bad request in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req.Header.Add("Content-Type", "application/json")

	log.Println("Sent data to server")

	// Reading the response from the RPi
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Bad response in routing/startExp.go: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Printf("Bad request in routing/startExp.go. Wanted 200, received : %v\n", res.StatusCode)
		return
	}

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Bad request in routing/RPiForward.go")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newReq routing.GetDataRes
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Couldn't Unmarshal data in routing/getDataFromServer.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return newReq.Vault
}
