package main

import (
	"database/sql"
	"log"

	"github.com/git-adithyanair/cs130-group-project/api"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	"github.com/git-adithyanair/cs130-group-project/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("could not load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("could not create server: ", err)
	}

	if config.Env == "dev" {
		db.PopulateWithData(store)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("could not start server: ", err)
	}

}
