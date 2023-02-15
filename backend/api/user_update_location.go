package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	api_error "github.com/git-adithyanair/cs130-group-project/errors"
	"github.com/git-adithyanair/cs130-group-project/token"
)

type UpdateUserLocationRequest struct {
	Address string  `json:"address" binding:"required"`
	PlaceID string  `json:"place_id" binding:"required"`
	XCoord  float64 `json:"x_coord" binding:"required,numeric"`
	YCoord  float64 `json:"y_coord" binding:"required,numeric"`
}

func (server *Server) UpdateUserLocation(ctx *gin.Context) {
	var req UpdateUserLocationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, unknownErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.UpdateUserLocationParams{
		ID:      authPayload.UserID,
		Address: req.Address,
		PlaceID: req.PlaceID,
		XCoord:  req.XCoord,
		YCoord:  req.YCoord,
	}

	err := server.queries.UpdateUserLocation(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoUser, err))
		} else {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
