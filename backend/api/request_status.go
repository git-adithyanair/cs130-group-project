//API not in use

package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	api_error "github.com/git-adithyanair/cs130-group-project/errors"
)

type UpdateRequestStatusRequest struct {
	ID     int64            `json:"id" binding:"required,min=1"`
	Status db.RequestStatus `json:"status" binding:"required"`
}

func (server *Server) UpdateRequestStatus(ctx *gin.Context) {
	var req UpdateRequestStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, unknownErrorResponse(err))
		return
	}

	arg := db.UpdateRequestStatusParams{
		ID:     req.ID,
		Status: req.Status,
	}

	request, err := server.queries.UpdateRequestStatus(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoRequest, err))
		} else {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, request)
}
