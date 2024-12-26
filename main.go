package main

import (
	"inventory/main/api"
	"inventory/main/db"
	"inventory/main/util"
	"log"
)

func main() {
	conn, err := util.Connect("inventory.db")
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(*store)

	err = server.Start("localhost:8080")
	if err != nil {
		log.Fatal("fail starting server: ", err)
	}
}
