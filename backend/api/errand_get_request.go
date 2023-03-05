//API not in use

package api

import (
	"database/sql"
	"errors"
	"net/http"

	api_error "github.com/git-adithyanair/cs130-group-project/errors"

	"github.com/git-adithyanair/cs130-group-project/token"

	"github.com/gin-gonic/gin"
)

type GetRequestByErrandRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetRequestsByErrand(ctx *gin.Context) {
	var req GetRequestByErrandRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, unknownErrorResponse(err))
		return
	}

	errand, err := server.queries.GetErrand(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoErrand, err))
		} else {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		}
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.UserID != errand.UserID {
		err = errors.New("incorrect errand_id: user does not own errand")
		ctx.JSON(http.StatusUnauthorized, errorResponse(api_error.ErrInvalidUserForErrand, err))
		return
	}

	id := sql.NullInt64{Int64: req.ID, Valid: true}
	requests, err := server.queries.GetRequestsByErrandId(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, requests)
}
