package api

import (
	"database/sql"
	"net/http"

	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	api_error "github.com/git-adithyanair/cs130-group-project/errors"
	"github.com/git-adithyanair/cs130-group-project/token"

	"github.com/gin-gonic/gin"
)

type JoinCommunityRequest struct {
	ID int64 `json:"id" binding:"required,min=1"`
}

func (server *Server) JoinCommunity(ctx *gin.Context) {
	var req JoinCommunityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, unknownErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	community, err := server.queries.GetCommunity(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoCommunity, err))
		} else {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		}
		return
	}

	arg := db.CreateMemberParams{
		UserID:      authPayload.UserID,
		CommunityID: req.ID,
	}

	_, err = server.queries.CreateMember(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, community)
}
