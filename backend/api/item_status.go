package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/git-adithyanair/cs130-group-project/util"

	"github.com/gin-gonic/gin"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	api_error "github.com/git-adithyanair/cs130-group-project/errors"
)

type UpdateItemStatusRequest struct {
	ID    int64 `json:"id" binding:"required,min=1"`
	Found bool  `json:"found"`
}

func (server *Server) UpdateItemStatus(ctx *gin.Context) {
	var req UpdateItemStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, unknownErrorResponse(err))
		return
	}

	arg := db.UpdateItemFoundParams{
		ID:    req.ID,
		Found: sql.NullBool{Bool: req.Found, Valid: true},
	}

	item, err := server.queries.UpdateItemFound(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoItem, err))
		} else {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		}
		return
	}

	user, err := server.queries.GetUser(ctx, item.RequestedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoUser, err))
		} else {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		}
		return
	}

	msg := fmt.Sprintf("Your item, %s, from your request has been found.", item.Name)
	if !req.Found {
		msg = fmt.Sprintf("Your item, %s, from your request cannot be found by the buyer.", item.Name)
	}
	util.NotifyUser(user.PhoneNumber, msg)

	ctx.JSON(http.StatusOK, item)
}
