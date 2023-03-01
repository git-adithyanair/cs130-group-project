package api

import (
	"database/sql"
	"net/http"

	api_error "github.com/git-adithyanair/cs130-group-project/errors"

	"github.com/git-adithyanair/cs130-group-project/token"

	"github.com/gin-gonic/gin"
)

func (server *Server) GetActiveErrand(ctx *gin.Context) {

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	activeErrand, err := server.queries.GetActiveErrand(ctx, authPayload.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, gin.H{})
		} else {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		}
		return
	}

	requests, err := server.queries.GetRequestsByErrandId(ctx, sql.NullInt64{Int64: activeErrand.ID, Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		return
	}

	requestRes := make([]activeErrandRequestResponse, len(requests))
	for i, request := range requests {
		items, err := server.queries.GetItemsByRequest(ctx, request.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoItem, err))
			} else {
				ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
			}
			return
		}

		user, err := server.queries.GetUser(ctx, request.UserID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoUser, err))
			} else {
				ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
			}
			return
		}

		requestRes[i] = activeErrandRequestResponse{
			Request: request,
			Items:   items,
			User:    newUserResponse(user),
		}

		if request.StoreID.Valid {
			store, err := server.queries.GetStore(ctx, request.StoreID.Int64)
			if err != nil {
				if err == sql.ErrNoRows {
					ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoStore, err))
				} else {
					ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
				}
				return
			}

			requestRes[i].Store = &store
		}
	}

	res := activeErrandResponse{
		Errand:   activeErrand,
		Requests: requestRes,
	}

	ctx.JSON(http.StatusOK, res)
}
