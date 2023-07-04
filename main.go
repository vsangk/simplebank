package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vsangk/simplebank/api"
	db "github.com/vsangk/simplebank/db/sqlc"
	"github.com/vsangk/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot connect config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.DBAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
