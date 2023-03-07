package api

import (
	"database/sql"
	"net/http"

	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	api_error "github.com/git-adithyanair/cs130-group-project/errors"

	"github.com/gin-gonic/gin"
	"github.com/git-adithyanair/cs130-group-project/token"
)

func (server *Server) GetUserRequest(ctx *gin.Context) {

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	pendingRequests := getRequestsForStatus(ctx, server, authPayload.UserID, db.RequestStatusPending)
	if pendingRequests == nil {
		return
	}
	inProgressRequests := getRequestsForStatus(ctx, server, authPayload.UserID, db.RequestStatusInProgress)
	if inProgressRequests == nil {
		return
	}
	completeRequests := getRequestsForStatus(ctx, server, authPayload.UserID, db.RequestStatusCompleted)
	if completeRequests == nil {
		return
	}

	response := userRequestsResponse{
		Pending:    pendingRequests,
		InProgress: inProgressRequests,
		Complete:   completeRequests,
	}

	ctx.JSON(http.StatusOK, response)

}

func getRequestsForStatus(ctx *gin.Context, server *Server, userID int64, status db.RequestStatus) []userRequestsDetailResponse {
	arg := db.GetRequestsForUserByStatusParams{
		UserID: userID,
		Status: status,
	}

	requests, err := server.queries.GetRequestsForUserByStatus(ctx, arg)
	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
			return nil
		}
	}

	requestRes := make([]userRequestsDetailResponse, len(requests))
	for i, request := range requests {
		items, err := server.queries.GetItemsByRequest(ctx, request.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoItem, err))
			} else {
				ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
			}
			return nil
		}

		community, err := server.queries.GetCommunity(ctx, request.CommunityID.Int64)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoCommunity, err))
			} else {
				ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
			}
			return nil
		}

		requestRes[i] = userRequestsDetailResponse{
			Request:       request,
			Items:         items,
			CommunityName: community.Name,
		}

		if request.StoreID.Valid {
			store, err := server.queries.GetStore(ctx, request.StoreID.Int64)
			if err != nil {
				if err == sql.ErrNoRows {
					ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoStore, err))
				} else {
					ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
				}
				return nil
			}

			requestRes[i].Store = &store
		}
	}

	return requestRes
}
