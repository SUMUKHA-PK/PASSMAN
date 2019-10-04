package main

import (
	"log"

	"github.com/SUMUKHA-PK/Basic-Golang-Server/server"
	"github.com/SUMUKHA-PK/PASSMAN/server/routing"
)

func main() {

	m := make(map[string]int)
	r = routing.SetupRouting(r)
	serverData = server.Data{
		Router:        r,
		Port:          "66666",
		HTTPS:         false,
		ConnectionMap: m,
	}

	err := server.Server(serverData)
	if err != nil {
		log.Fatalf("Could not start server : %v", err)
	}
}
