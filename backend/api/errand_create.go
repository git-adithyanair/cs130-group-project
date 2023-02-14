package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	api_error "github.com/git-adithyanair/cs130-group-project/errors"
	"github.com/git-adithyanair/cs130-group-project/token"

	"github.com/gin-gonic/gin"
)

type CreateErrandRequest struct {
	CommunityID int64   `json:"community_id" binding:"required,min=1"`
	RequestIDs  []int64 `json:"request_ids" binding:"required"`
}

func (server *Server) CreateErrand(ctx *gin.Context) {
	var req CreateErrandRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, unknownErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if _, err := server.queries.GetCommunity(ctx, req.CommunityID); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoCommunity, err))
		} else {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		}
		return
	}

	if len(req.RequestIDs) == 0 {
		ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrErrandNoRequests, errors.New("invalid errand, missing requests")))
		return
	}

	for _, requestID := range req.RequestIDs {
		if _, err := server.queries.GetRequest(ctx, requestID); err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoRequest, err))
			} else {
				ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
			}
			return
		}
	}

	arg := db.CreateErrandParams{
		UserID:      authPayload.UserID,
		CommunityID: req.CommunityID,
	}

	errand, err := server.queries.CreateErrand(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		return
	}

	for _, requestID := range req.RequestIDs {
		arg := db.UpdateRequestErrandAndStatusParams{
			ID:       requestID,
			ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
			Status:   db.RequestStatusInProgress,
		}
		server.queries.UpdateRequestErrandAndStatus(ctx, arg)
	}

	ctx.JSON(http.StatusCreated, errand)
}
