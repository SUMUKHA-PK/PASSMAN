package routing

import (
	"net/http"

	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"
	"github.com/gorilla/mux"
)

var ServerData server.Data

// SetupRouting adds all the routes
func SetupRouting(r *mux.Router) *mux.Router {
	r.HandleFunc("/removeDataFromServer", RemoveDataFromServer).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/getDataFromServer", GetDataFromServer).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/putDataToServer", PutDataToServer).Methods(http.MethodPost, http.MethodOptions)
	return r
}
