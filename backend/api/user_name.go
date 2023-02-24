package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	api_error "github.com/git-adithyanair/cs130-group-project/errors"
	"github.com/git-adithyanair/cs130-group-project/token"
)

type UpdateUserNameRequest struct {
	Name string `json:"name" binding:"required,min=1"`
}

func (server *Server) UpdateUserName(ctx *gin.Context) {
	var req UpdateUserNameRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, unknownErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.UpdateUserNameParams{
		ID:       authPayload.UserID,
		FullName: req.Name,
	}

	err := server.queries.UpdateUserName(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoUser, err))
		} else {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
