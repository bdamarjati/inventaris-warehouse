package db_test

import (
	"inventory/main/db"
	"inventory/main/util"
	"log"
	"os"
	"testing"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	conn, err := util.Connect("../../inventory.db")
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	testQueries = db.New(*conn)

	os.Exit(m.Run())
}
