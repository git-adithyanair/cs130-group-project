package main

import (
	"log"

	"github.com/git-adithyanair/cs130-group-project/api"
	"github.com/git-adithyanair/cs130-group-project/util"
)

func main() {

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("could not load config: ", err)
	}

	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("could not create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("could not start server: ", err)
	}

}
