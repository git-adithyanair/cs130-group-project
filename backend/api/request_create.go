package api

import (
	"database/sql"
	"errors"
	"net/http"

	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	"github.com/git-adithyanair/cs130-group-project/token"

	"github.com/gin-gonic/gin"
)

type CreateRequestRequestItem struct {
	Name           string              `json:"name" binding:"required,min=1"`
	QuantityType   db.ItemQuantityType `json:"quantity_type" binding:"required,min=2"`
	Quantity       float64             `json:"quantity" binding:"required,min=1"`
	PreferredBrand string              `json:"preferred_brand"`
	Image          string              `json:"image"`
	ExtraNotes     string              `json:"extra_notes"`
}

type CreateRequestRequest struct {
	CommunityID int64                      `json:"community_id" binding:"required,min=1"`
	StoreID     *int64                     `json:"store_id"`
	Items       []CreateRequestRequestItem `json:"items" binding:"required"`
}

func (server *Server) CreateRequest(ctx *gin.Context) {
	var req CreateRequestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if _, err := server.queries.GetCommunity(ctx, req.CommunityID); err != nil {
		if err == sql.ErrNoRows {
			err := errors.New("No community exists with given ID.")
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	arg := db.CreateRequestParams{
		UserID:      authPayload.UserID,
		CommunityID: sql.NullInt64{Int64: req.CommunityID, Valid: true},
	}
	if req.StoreID != nil {
		if _, err := server.queries.GetStore(ctx, *req.StoreID); err != nil {
			if err == sql.ErrNoRows {
				err := errors.New("No store exists with given ID.")
				ctx.JSON(http.StatusNotFound, errorResponse(err))
			} else {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			}
			return
		}

		arg = db.CreateRequestParams{
			UserID:      authPayload.UserID,
			CommunityID: sql.NullInt64{Int64: req.CommunityID, Valid: true},
			StoreID:     sql.NullInt64{Int64: *req.StoreID, Valid: true},
		}
	}

	request, err := server.queries.CreateRequest(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	for _, item := range req.Items {
		arg := db.CreateItemParams{
			Name:           item.Name,
			RequestedBy:    authPayload.UserID,
			RequestID:      request.ID,
			QuantityType:   item.QuantityType,
			Quantity:       item.Quantity,
			PreferredBrand: sql.NullString{String: item.PreferredBrand, Valid: len(item.PreferredBrand) > 0},
			Image:          sql.NullString{String: item.Image, Valid: len(item.Image) > 0},
			ExtraNotes:     sql.NullString{String: item.ExtraNotes, Valid: len(item.ExtraNotes) > 0},
		}

		_, err := server.queries.CreateItem(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	ctx.JSON(http.StatusCreated, request)
}
