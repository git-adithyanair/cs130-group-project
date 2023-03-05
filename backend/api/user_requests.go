package api

import (
	"database/sql"
	"net/http"

	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"

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

func getRequestsForStatus(ctx *gin.Context, server *Server, userID int64, status db.RequestStatus) []db.Request {
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

	return requests
}
