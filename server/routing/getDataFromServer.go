package routing

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/SUMUKHA-PK/PASSMAN/server/redis"
)

// GetDataFromServer syncs data with the client and server
func GetDataFromServer(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting data from server")

	// Parse the incoming request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Bad request in routing/startExp.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newReq GetDataReq
	err = json.Unmarshal(body, &newReq)
	if err != nil {
		log.Printf("Couldn't Unmarshal data in routing/getDataFromServer.go : %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vault, err := redis.Retrieve(newReq.AuthPwd)
	if err != nil {
		log.Printf("Couldn't get data from REDIS\n")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	outData := &GetDataRes{200, "Successfully queried for requested data", vault}
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
