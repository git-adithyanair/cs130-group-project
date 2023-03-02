package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	api_error "github.com/git-adithyanair/cs130-group-project/errors"
	"github.com/git-adithyanair/cs130-group-project/util"
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

	if req.IsComplete {
		requests, err := server.queries.GetRequestsByErrandId(ctx, sql.NullInt64{Int64: errand.ID, Valid: true})
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoRequest, err))
			} else {
				ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
			}
			return
		}

		for _, request := range requests {
			arg := db.UpdateRequestStatusParams{
				ID:     request.ID,
				Status: db.RequestStatusCompleted,
			}
			request, _ := server.queries.UpdateRequestStatus(ctx, arg)

			user, _ := server.queries.GetUser(ctx, request.UserID)
			util.NotifyUser(user.PhoneNumber, "Your request has been completed! Thank you for using GoodGrocer!")
		}
	}

	ctx.JSON(http.StatusOK, errand)
}
