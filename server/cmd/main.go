package main

import (
	"log"

	"github.com/gorilla/mux"

	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"
	"github.com/SUMUKHA-PK/PASSMAN/server/routing"
)

func main() {

	r := mux.NewRouter()
	m := make(map[string]int)
	r = routing.SetupRouting(r)
	counter := 0
	routing.ServerData = server.Data{
		Router:        r,
		Port:          "6666",
		HTTPS:         false,
		ConnectionMap: m,
		Count:         counter,
	}

	err := server.Server(&routing.ServerData)
	if err != nil {
		log.Fatalf("Could not start server : %v", err)
	}
}
