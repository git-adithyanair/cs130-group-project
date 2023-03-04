package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
)

func (server *Server) GetAllCommunities(ctx *gin.Context) {

	arg := db.ListCommunitiesParams{
		Limit:  100,
		Offset: 0,
	}

	communities, err := server.queries.ListCommunities(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		return
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
