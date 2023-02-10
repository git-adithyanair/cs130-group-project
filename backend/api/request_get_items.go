package api

import (
	"database/sql"
	"errors"
	"net/http"

	api_error "github.com/git-adithyanair/cs130-group-project/errors"

	"github.com/git-adithyanair/cs130-group-project/token"

	"github.com/gin-gonic/gin"
)

type GetItemsByRequestRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetItemsByRequest(ctx *gin.Context) {
	var req GetItemsByRequestRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, unknownErrorResponse(err))
		return
	}

	request, err := server.queries.GetRequest(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoRequest, err))
		} else {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		}
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.UserID != request.UserID {
		err = errors.New("incorrect request_id: user does not own request")
		ctx.JSON(http.StatusUnauthorized, errorResponse(api_error.ErrInvalidUserForRequest, err))
		return
	}

	items, err := server.queries.GetItemsByRequest(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, items)
}
