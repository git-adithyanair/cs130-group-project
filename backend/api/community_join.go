package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	"github.com/git-adithyanair/cs130-group-project/token"

	"github.com/gin-gonic/gin"
)

type JoinCommunityRequest struct {
	ID int64 `json:"id" binding:"required,min=1"`
}

func (server *Server) JoinCommunity(ctx *gin.Context) {
	var req JoinCommunityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if _, err := server.queries.GetCommunity(ctx, req.ID); err != nil {
		if err == sql.ErrNoRows {
			err := errors.New("No community exists with given ID.")
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	arg := db.CreateMemberParams{
		UserID:      authPayload.UserID,
		CommunityID: req.ID,
	}

	member, err := server.queries.CreateMember(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, member)
}
