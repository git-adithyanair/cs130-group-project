package api

import (
	"database/sql"
	"net/http"

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

	res := activeErrandResponse{
		Errand:   activeErrand,
		Requests: requests,
	}
	ctx.JSON(http.StatusOK, res)
}
