package api

import (
	"database/sql"
	"net/http"

	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/git-adithyanair/cs130-group-project/token"
)

func (server *Server) GetUserAcceptedRequest(ctx *gin.Context) {

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.GetRequestsForUserByStatusParams{
		UserID: authPayload.UserID,
		Status: db.RequestStatusInProgress,
	}

	requests, err := server.queries.GetRequestsForUserByStatus(ctx, arg)
	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
			return
		}
	}

	ctx.JSON(http.StatusOK, requests)

}
