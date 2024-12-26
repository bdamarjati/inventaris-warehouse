package api

import (
	"inventory/main/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createInventoryRequest struct {
	Name       string `json:"name"`
	Quantity   int64  `json:"quantity"`
	CategoryId int64  `json:"category_id"`
	Condition  int    `json:"condition"`
	Status     int64  `json:"status"`
}

func (server *Server) createInventory(ctx *gin.Context) {
	var req createInventoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateInventoryParams{
		Name:       req.Name,
		Quantity:   req.Quantity,
		CategoryId: req.CategoryId,
		Condition:  req.Condition,
		Status:     req.Status,
	}

	inventory, err := server.store.CreateInventory(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, inventory)
}

type getInventoryRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getInventory(ctx *gin.Context) {
	var req getInventoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	inventory, err := server.store.GetInventory(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, inventory)
}

type listInventoryRequest struct {
	Page int `uri:"page" binding:"required,min=1"`
	Size int `uri:"size" binding:"required,min=2,max=10"`
}

func (server *Server) listInventories(ctx *gin.Context) {
	var req listInventoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListInventoryParams{
		Limit:  req.Size,
		Offset: (req.Page - 1) * req.Size,
	}

	inventories, err := server.store.ListInventories(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, inventories)
}

type updateInventoryRequest struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Quantity   int64  `json:"quantity"`
	CategoryId int64  `json:"category_id"`
	Condition  int    `json:"condition"`
	Status     int64  `json:"status"`
}

func (server *Server) updateInventory(ctx *gin.Context) {
	var req updateInventoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateInventoryParams{
		ID:         req.ID,
		Name:       req.Name,
		Quantity:   req.Quantity,
		CategoryId: req.CategoryId,
		Condition:  req.Condition,
		Status:     req.Status,
	}

	_, err := server.store.UpdateInventory(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

type deleteInventoryRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteInventory(ctx *gin.Context) {
	var req deleteInventoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := server.store.DeleteInventory(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
