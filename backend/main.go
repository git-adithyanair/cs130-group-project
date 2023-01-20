package main

import (
	"log"

	"github.com/git-adithyanair/cs130-group-project/api"
)

func main() {

	server, err := api.NewServer()
	if err != nil {
		log.Fatal("could not create server: ", err)
	}

	err = server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("could not start server: ", err)
	}

}
