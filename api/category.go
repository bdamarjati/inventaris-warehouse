package api

import (
	"database/sql"
	"inventory/main/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createCategory(ctx *gin.Context) {
	var req createCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	category, err := server.store.CreateCategory(req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

type getCategoryRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getCategory(ctx *gin.Context) {
	var req getCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	category, err := server.store.GetCategory(req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

type listCategoryRequest struct {
	Page int `uri:"page" binding:"required,min=1"`
	Size int `uri:"size" binding:"required,min=2,max=10"`
}

func (server *Server) listCategories(ctx *gin.Context) {
	var req listCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCategoriesParams{
		Limit:  req.Size,
		Offset: (req.Page - 1) * req.Size,
	}

	categories, err := server.store.ListCategories(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

type updateCategoryRequest struct {
	ID   int64  `json:"id" binding:"required,min=1"`
	Name string `json:"name" binding:"required"`
}

func (server *Server) updateCategory(ctx *gin.Context) {
	var req updateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCategoryParams{
		ID:   req.ID,
		Name: req.Name,
	}

	_, err := server.store.UpdateCategory(arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

type deleteCategoryRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteCategory(ctx *gin.Context) {
	var req deleteCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := server.store.DeleteCategory(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
