package api

import (
	"database/sql"
	"inventory/main/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createStatusRequest struct {
	Description string `json:"description" binding:"required"`
}

func (server *Server) createStatus(ctx *gin.Context) {
	var req createStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, errorResponse(err))
		return
	}

	status, err := server.store.CreateStatus(req.Description)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, status)
}

type getStatusRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getStatus(ctx *gin.Context) {
	var req getStatusRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	status, err := server.store.GetStatus(req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, status)
}

type listStatusRequest struct {
	Page int `uri:"page" binding:"required,min=1"`
	Size int `uri:"size" binding:"required,min=2,max=10"`
}

func (server *Server) listStatus(ctx *gin.Context) {
	var req listStatusRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListStatusParams{
		Limit:  req.Size,
		Offset: (req.Page - 1) * req.Size,
	}

	statuses, err := server.store.ListStatus(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, statuses)
}

type updateStatusRequest struct {
	ID          int64  `json:"id" binding:"required,min=1"`
	Description string `json:"description" binding:"required"`
}

func (server *Server) updateStatus(ctx *gin.Context) {
	var req updateStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateStatusParams{
		ID:          req.ID,
		Description: req.Description,
	}

	_, err := server.store.UpdateStatus(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

type deleteStatusRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteStatus(ctx *gin.Context) {
	var req deleteStatusRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := server.store.DeleteStatus(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
