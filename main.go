package main

import (
	"inventory/main/api"
	"inventory/main/db"
	"inventory/main/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config file: ", err)
	}

	conn, err := util.Connect(config.DBName)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(config, *store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("fail starting server: ", err)
	}
}
