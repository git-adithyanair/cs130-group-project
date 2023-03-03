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
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
			return
		}
	}

	res := make([]userCommunityResponse, len(communities))
	for i, community := range communities {
		memberCount, err := server.queries.GetMemberCountInCommunity(ctx, community.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
			return
		}

		res[i] = userCommunityResponse{
			Community:   community,
			MemberCount: memberCount,
		}
	}

	ctx.JSON(http.StatusOK, res)

}
