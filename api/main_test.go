package api

import (
	"inventory/main/db"
	"inventory/main/util"
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var testStore *db.Store

func TestMain(m *testing.M) {
	conn, err := util.Connect("../inventory.db")
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	testStore = db.NewStore(conn)

	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

func newTestServer(store db.Store) *Server {
	server := NewServer(store)
	return server
}
