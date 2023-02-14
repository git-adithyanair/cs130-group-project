package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	api_error "github.com/git-adithyanair/cs130-group-project/errors"
)

type UpdateItemStatusRequest struct {
	ID    int64 `json:"id" binding:"required,min=1"`
	Found bool  `json:"found"`
}

func (server *Server) UpdateItemStatus(ctx *gin.Context) {
	var req UpdateItemStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, unknownErrorResponse(err))
		return
	}

	arg := db.UpdateItemFoundParams{
		ID:    req.ID,
		Found: sql.NullBool{Bool: req.Found, Valid: true},
	}

	item, err := server.queries.UpdateItemFound(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoItem, err))
		} else {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, item)
}
