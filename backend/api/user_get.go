package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	api_error "github.com/git-adithyanair/cs130-group-project/errors"
	"github.com/git-adithyanair/cs130-group-project/token"
)

func (server *Server) GetCurrentUser(ctx *gin.Context) {

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	user, err := server.queries.GetUser(ctx, authPayload.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoUser, err))
		} else {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		}
		return
	}

	communityCount, err := server.queries.GetNumberOfUserCommunities(ctx, authPayload.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		return
	}

	response := newGetUserResponse(user, communityCount)

	ctx.JSON(http.StatusOK, response)

}
