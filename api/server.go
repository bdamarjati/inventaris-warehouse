package api

import (
	"inventory/main/db"
	"inventory/main/token"
	"inventory/main/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router     *gin.Engine
	store      db.Store
	config     util.Config
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) *Server {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil
	}

	server := &Server{
		store:      store,
		config:     config,
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

	authRoute := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoute.GET("/auth", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "authorized"})
	})

	authRoute.POST("/category", server.createCategory)
	authRoute.GET("/category/:id", server.getCategory)
	authRoute.GET("/categories/:size/:page", server.listCategories)
	authRoute.PUT("/category", server.updateCategory)
	authRoute.DELETE("/category/:id", server.deleteCategory)

	authRoute.POST("/status", server.createStatus)
	authRoute.GET("/status/:id", server.getStatus)
	authRoute.GET("/statuses/:size/:page", server.listStatus)
	authRoute.PUT("/status", server.updateStatus)
	authRoute.DELETE("/status/:id", server.deleteStatus)

	authRoute.POST("/inventory", server.createInventory)
	authRoute.GET("/inventory/:id", server.getInventory)
	authRoute.GET("/inventories/:size/:page", server.listInventories)
	authRoute.PUT("/inventory", server.updateInventory)
	authRoute.DELETE("/inventory/:id", server.deleteInventory)

	authRoute.POST("/consumable", server.createConsumable)
	authRoute.GET("/consumable/:id", server.getConsumable)
	authRoute.GET("/consumables/:size/:page", server.listConsumables)
	authRoute.PUT("/consumable", server.updateConsumable)
	authRoute.DELETE("/consumable/:id", server.deleteConsumable)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
