package api

import (
	"inventory/main/db"
	"inventory/main/util"
	"log"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

var testStore *db.Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("..")
	if err != nil {
		log.Fatal("failed to load config file")
	}
	path := "../" + config.DBName
	conn, err := util.Connect(path)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	testStore = db.NewStore(conn)

	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

func newTestServer(store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:  util.RandomText(32),
		AccesTokenDuration: time.Minute,
	}
	server := NewServer(config, store)
	return server
}
