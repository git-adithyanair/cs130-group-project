package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	"github.com/git-adithyanair/cs130-group-project/token"

	"github.com/gin-gonic/gin"
)

type CreateItemRequest struct {
	Name           string              `json:"name" binding:"required,min=1"`
	RequestID      int64               `json:"request_id" binding:"required,min=1"`
	QuantityType   db.ItemQuantityType `json:"quantity_type" binding:"required,min=2"`
	Quantity       float64             `json:"quantity" binding:"required,min=1"`
	PreferredBrand string              `json:"preferred_brand"`
	Image          string              `json:"image"`
	ExtraNotes     string              `json:"extra_notes"`
}

func (server *Server) CreateItem(ctx *gin.Context) {
	var req CreateItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if _, err := server.queries.GetRequest(ctx, req.RequestID); err != nil {
		if err == sql.ErrNoRows {
			err := errors.New("No request exists with given ID.")
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateItemParams{
		Name:           req.Name,
		RequestedBy:    authPayload.UserID,
		RequestID:      req.RequestID,
		QuantityType:   req.QuantityType,
		Quantity:       req.Quantity,
		PreferredBrand: sql.NullString{String: req.PreferredBrand, Valid: len(req.PreferredBrand) > 0},
		Image:          sql.NullString{String: req.Image, Valid: len(req.Image) > 0},
		ExtraNotes:     sql.NullString{String: req.ExtraNotes, Valid: len(req.ExtraNotes) > 0},
	}

	item, err := server.queries.CreateItem(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, item)
}
