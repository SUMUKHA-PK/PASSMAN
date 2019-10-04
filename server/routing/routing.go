package routing

import (
	"github.com/gorilla/mux"
)

// SetupRouting adds all the routes
func SetupRouting(r mux.Router) mux.Router {
	// r.HandleFunc("/gaWebhook", GoogleAssistantWebHook).Methods(http.MethodPost, http.MethodOptions)
	// r.HandleFunc("/imageRecognition", GetImageResponse).Methods(http.MethodPost, http.MethodOptions)
	return r
}
