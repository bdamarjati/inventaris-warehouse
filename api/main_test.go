package api

import (
	"inventory/main/db"
	"inventory/main/util"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

var testServer *Server
var testUser *loginUserResponse

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

	testStore := db.NewStore(conn)
	testServer = newTestServer(*testStore)

	user, err := createRandomUser()
	if err != nil {
		log.Fatal("failed to create test user: ", err)
	}
	loggedUser, err := getAccessToken(loginUserRequest{
		Username: user.Username,
		Password: user.Password,
	})
	if err != nil {
		log.Fatal("failed to get test user access token: ", err)
	}
	testUser = &loggedUser

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

func addAuthorization(req *http.Request) error {
	auth := "Bearer " + testUser.AccessToken
	req.Header.Add("Authorization", auth)
	return nil
}
