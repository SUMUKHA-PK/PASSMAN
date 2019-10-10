package passman

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/SUMUKHA-PK/PASSMAN/client/redis"
	"github.com/SUMUKHA-PK/PASSMAN/server/routing"
	"github.com/gookit/color"
	"golang.org/x/crypto/argon2"
)

// SyncDataWithServer pushes all the data to the server
// It ensures no data is lost while syncing, by the given steps:
// 1. Auth with the users credentials.
// 2. Request the server for the vault data.
// 3. Decrypt both vaults.
// 4. Update the client vault, timestamp based.
// 5. Encrypt the vault, save the vault in REDIS and send to server.
func SyncDataWithServer() {
	color.Info.Printf("PASSMAN Server sync sequence\n\n")

	err := serverCommunication("sync")
	if err != nil {
		return
	}
	fmt.Printf("\n\n")

	color.Success.Println("Vault data synced with server.")
}

func serverCommunication(syncOrRemove string) error {
	username, vault, vaultPwd, err := verifyAndGetVaultData()
	if err != nil {
		color.Error.Println(err)
		return err
	}

	authPwd := argon2.IDKey(vaultPwd, []byte(username), 1, 64*1024, 4, 32)
	vaultServer, err := getDataFromServer(string(authPwd))
	if err != nil {
		color.Error.Println("No data available on server.")
	}

	byteEncryptedVault := []byte(vault.Vault)
	// Case where there is no data in the server skips this block
	if vaultServer != "" {
		decryptedVault, err := decryptVault([]byte(vault.Vault), vaultPwd)
		if err != nil {
			color.Error.Println("Error in decryption: %v", err)
			return err
		}
		decryptedVaultServer, err := decryptVault([]byte(vaultServer), vaultPwd)
		if err != nil {
			color.Error.Printf("Error in decryption: %v", err)
			return err
		}

		var vaultMap, vaultMapServer map[string]Vault
		err = json.Unmarshal([]byte(decryptedVault), &vaultMap)
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = json.Unmarshal([]byte(decryptedVaultServer), &vaultMapServer)
		if err != nil {
			fmt.Println(err)
			return err
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

		var byteMap []byte
		if syncOrRemove == "sync" {
			byteMap, err = json.Marshal(vaultMap)
			if err != nil {
				color.Error.Println(err)
				return err
			}
			byteEncryptedVault, err = encryptVault(byteMap, vaultPwd)
			if err != nil {
				return err
			}
			err = redis.Update(username, string(byteEncryptedVault))
			if err != nil {
				color.Error.Printf("Can't add data to Redis DB: %v", err)
				return err
			}
		} else if syncOrRemove == "remove" {
			byteMap, err = json.Marshal(getInitMap(username))
			if err != nil {
				color.Error.Println(err)
				return err
			}
		}

		byteEncryptedVault, err = encryptVault(byteMap, vaultPwd)
		if err != nil {
			return err
		}
	}

	err = putDataToServer(string(authPwd), byteEncryptedVault)
	return err
}

// getDataFromServer contacts the server to get vault data
func getDataFromServer(authPwd string) (string, error) {

	URL := "http://127.0.0.1:6666/getDataFromServer"
	outData := &routing.GetDataReq{authPwd}
	payload, err := json.Marshal(outData)
	if err != nil {
		log.Printf("Can't Marshall to JSON in routing/startExp.go : %v\n", err)
		return "", err
	}

	req, err := http.NewRequest("POST", URL, strings.NewReader(string(payload)))
	if err != nil {
		log.Printf("Bad request in routing/startExp.go : %v\n", err)
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Bad response in serverSync.go: %v\n", err)
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Printf("Bad request in serverSync.go. Wanted 200, received : %v\n\n", res.StatusCode)
		return "", errors.New("Bad Request")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Bad request in serverSync.go")
		log.Println(err)
		return "", err
	}

	var newReq routing.GetDataRes
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Couldn't Unmarshal data in serverSync.go : %v\n", err)
		return "", err
	}

	return string(newReq.Vault), nil
}

func putDataToServer(authPwd string, vault []byte) error {
	URL := "http://127.0.0.1:6666/putDataToServer"
	outData := &routing.PutDataReq{authPwd, vault}
	payload, err := json.Marshal(outData)
	if err != nil {
		color.Error.Printf("Can't Marshall to JSON in routing/startExp.go : %v\n", err)
		return err
	}
	req, err := http.NewRequest("POST", URL, strings.NewReader(string(payload)))
	if err != nil {
		color.Error.Printf("Bad request in routing/startExp.go : %v\n", err)
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		color.Error.Printf("Bad response in serverSync.go: %v\n", err)
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Printf("Bad request in serverSync.go. Wanted 200, received : %v\n", res.StatusCode)
		return errors.New("Bad request")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		color.Error.Printf("Bad request in serverSync.go: %v", err)
		return err
	}

	var newReq routing.PutDataRes
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		color.Error.Printf("Couldn't Unmarshal data in serverSync.go : %v\n", err)
		return err
	}

	return nil
}

// getInitMap returns a dummy map to add to server on data removal
func getInitMap(username string) map[string]Vault {
	m := make(map[string]Vault)
	m[username] = Vault{username, time.Now()}
	return m
}
