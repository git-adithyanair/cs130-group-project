package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	api_error "github.com/git-adithyanair/cs130-group-project/errors"
)

type UpdateErrandStatusRequest struct {
	ID         int64 `json:"id" binding:"required,min=1"`
	IsComplete bool  `json:"is_complete"`
}

func (server *Server) UpdateErrandStatus(ctx *gin.Context) {
	var req UpdateErrandStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, unknownErrorResponse(err))
		return
	}

	arg := db.UpdateErrandStatusParams{
		ID:         req.ID,
		IsComplete: req.IsComplete,
	}

	errand, err := server.queries.UpdateErrandStatus(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoErrand, err))
		} else {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, errand)
}
