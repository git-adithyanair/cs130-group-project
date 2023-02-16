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

	queries := db.New(conn)
	server, err := api.NewServer(config, queries)
	if err != nil {
		log.Fatal("could not create server: ", err)
	}

	if config.Env == "dev" {
		db.PopulateWithData(queries)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("could not start server: ", err)
	}

}
