package routing

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/PASSMAN/server/redis"
)

// PutDataToServer syncs data with the client and server
func PutDataToServer(w http.ResponseWriter, r *http.Request) {
	log.Println("Putting data to server")

	// Parse the incoming request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newReq PutDataReq
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Couldn't Unmarshal data in routing/getDataFromServer.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = redis.Update(newReq.AuthPwd, newReq.Vault)
	if err != nil {
		log.Printf("Couldn't add data to REDIS\n")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	outData := &PutDataRes{200, "Successfully added data"}
	outJSON, err := json.Marshal(outData)
	if err != nil {
		log.Printf("Can't Marshall to JSON in routing/pastebin.go")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(outJSON)
}
