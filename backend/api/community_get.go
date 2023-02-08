package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetCommunityRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) GetCommunity(ctx *gin.Context) {
	var req GetCommunityRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	community, err := server.queries.GetCommunity(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, community)
}
