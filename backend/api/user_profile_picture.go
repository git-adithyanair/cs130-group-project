package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	api_error "github.com/git-adithyanair/cs130-group-project/errors"
	"github.com/git-adithyanair/cs130-group-project/token"
)

type UpdateUserProfilePicRequest struct {
	Image string `json:"image"`
}

func (server *Server) UpdateUserProfilePic(ctx *gin.Context) {
	var req UpdateUserProfilePicRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, unknownErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.UpdateUserProfilePictureParams{
		ID:             authPayload.UserID,
		ProfilePicture: req.Image,
	}

	user, err := server.queries.UpdateUserProfilePicture(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoUser, err))
		} else {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, user)
}
