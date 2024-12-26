package api

import (
	"inventory/main/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store  db.Store
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}

	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
