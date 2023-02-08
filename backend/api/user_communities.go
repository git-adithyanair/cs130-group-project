package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/git-adithyanair/cs130-group-project/token"
)

func (server *Server) GetUserCommunities(ctx *gin.Context) {

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	communities, err := server.queries.GetUserCommunities(ctx, authPayload.UserID)
	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, communities)

}
