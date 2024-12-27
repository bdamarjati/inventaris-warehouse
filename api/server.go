package api

import (
	"inventory/main/db"
	"inventory/main/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

const secretKey string = "12345678901234567890123456789012"

type Server struct {
	router     *gin.Engine
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(store db.Store) *Server {
	tokenMaker, err := token.NewJWTMaker(secretKey)
	if err != nil {
		return nil
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	router.POST("/user", server.createUser)
	router.POST("/login", server.loginUser)

	router.POST("/category", server.createCategory)
	router.GET("/category/:id", server.getCategory)
	router.GET("/categories/:size/:page", server.listCategories)
	router.PUT("/category", server.updateCategory)
	router.DELETE("/category/:id", server.deleteCategory)

	router.POST("/status", server.createStatus)
	router.GET("/status/:id", server.getStatus)
	router.GET("/statuses/:size/:page", server.listStatus)
	router.PUT("/status", server.updateStatus)
	router.DELETE("/status/:id", server.deleteStatus)

	router.POST("/inventory", server.createInventory)
	router.GET("/inventory/:id", server.getInventory)
	router.GET("/inventories/:size/:page", server.listInventories)
	router.PUT("/inventory", server.updateInventory)
	router.DELETE("/inventory/:id", server.deleteInventory)

	router.POST("/consumable", server.createConsumable)
	router.GET("/consumable/:id", server.getConsumable)
	router.GET("/consumables/:size/:page", server.listConsumables)
	router.PUT("/consumable", server.updateConsumable)
	router.DELETE("/consumable/:id", server.deleteConsumable)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
