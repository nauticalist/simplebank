package main

import (
	"database/sql"
	"github.com/nauticalist/simplebank/util"
	"log"

	_ "github.com/lib/pq"
	"github.com/nauticalist/simplebank/api"
	db "github.com/nauticalist/simplebank/db/sqlc"
)

func main()  {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatalln("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
