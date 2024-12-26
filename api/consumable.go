package api

import (
	"inventory/main/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createConsumableRequest struct {
	Name       string `json:"name"`
	Quantity   int64  `json:"quantity"`
	CategoryId int64  `json:"category_id"`
	Status     int64  `json:"status"`
}

func (server *Server) createConsumable(ctx *gin.Context) {
	var req createConsumableRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateConsumableParams{
		Name:       req.Name,
		Quantity:   req.Quantity,
		CategoryId: req.CategoryId,
		Status:     req.Status,
	}

	consumable, err := server.store.CreateConsumable(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, consumable)
}

type getConsumableRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getConsumable(ctx *gin.Context) {
	var req getConsumableRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	consumable, err := server.store.GetConsumable(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, consumable)
}

type listConsumableRequest struct {
	Page int `uri:"page" binding:"required,min=1"`
	Size int `uri:"size" binding:"required,min=2,max=10"`
}

func (server *Server) listConsumables(ctx *gin.Context) {
	var req listConsumableRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListConsumableParams{
		Limit:  req.Size,
		Offset: (req.Page - 1) * req.Size,
	}

	consumables, err := server.store.ListConsumables(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, consumables)
}

type updateConsumableRequest struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Quantity   int64  `json:"quantity"`
	CategoryId int64  `json:"category_id"`
	Status     int64  `json:"status"`
}

func (server *Server) updateConsumable(ctx *gin.Context) {
	var req updateConsumableRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateConsumableParams{
		ID:         req.ID,
		Name:       req.Name,
		Quantity:   req.Quantity,
		CategoryId: req.CategoryId,
		Status:     req.Status,
	}

	_, err := server.store.UpdateConsumable(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

type deleteConsumableRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteConsumable(ctx *gin.Context) {
	var req deleteConsumableRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := server.store.DeleteConsumable(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
