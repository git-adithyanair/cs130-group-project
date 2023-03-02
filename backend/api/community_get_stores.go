package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	api_error "github.com/git-adithyanair/cs130-group-project/errors"
)

type GetCommunityStoresRequest struct {
	ID int64 `uri:"id" binding:"min=0"`
}

func (server *Server) GetCommunityStores(ctx *gin.Context) {
	var req GetCommunityStoresRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, unknownErrorResponse(err))
		return
	}

	stores, err := server.queries.GetStoresByCommunity(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoStore, err))
		} else {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, stores)
}
