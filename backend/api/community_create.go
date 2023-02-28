package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	api_error "github.com/git-adithyanair/cs130-group-project/errors"
	"github.com/git-adithyanair/cs130-group-project/token"
)

type CreateCommunityRequestStore struct {
	Name    string  `json:"name" binding:"required"`
	XCoord  float64 `json:"x_coord" binding:"required,numeric"`
	YCoord  float64 `json:"y_coord" binding:"required,numeric"`
	PlaceID string  `json:"place_id" binding:"required"`
	Address string  `json:"address" binding:"required"`
}

type CreateCommunityRequest struct {
	Name         string                        `json:"name" binding:"required"`
	CenterXCoord float64                       `json:"center_x_coord" binding:"required,numeric"`
	CenterYCoord float64                       `json:"center_y_coord" binding:"required,numeric"`
	Range        int32                         `json:"range" binding:"required,numeric"`
	PlaceID      string                        `json:"place_id" binding:"required"`
	Address      string                        `json:"address" binding:"required"`
	Stores       []CreateCommunityRequestStore `json:"stores" binding:"required"`
}

func (server *Server) CreateCommunity(ctx *gin.Context) {
	var req CreateCommunityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, unknownErrorResponse(err))
		return
	}

	_, err := server.queries.GetCommunityByPlaceID(ctx, req.PlaceID)
	if err != nil {
		if err == sql.ErrNoRows {

			authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

			arg := db.CreateCommunityParams{
				Name:         req.Name,
				Admin:        authPayload.UserID,
				CenterXCoord: req.CenterXCoord,
				CenterYCoord: req.CenterYCoord,
				Range:        req.Range,
				PlaceID:      req.PlaceID,
				Address:      req.Address,
			}

			community, err := server.queries.CreateCommunity(ctx, arg)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(api_error.ErrCommunityCreateFail, err))
				return
			}

			_, err = server.queries.CreateMember(ctx, db.CreateMemberParams{
				CommunityID: community.ID,
				UserID:      authPayload.UserID,
			})
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
				return
			}

			if len(req.Stores) == 0 {
				ctx.JSON(http.StatusBadRequest, errorResponse(api_error.ErrCommunityCreateFail, errors.New("communities need to have at least one store")))
				return
			}

			for _, store := range req.Stores {

				existingStore, err := server.queries.GetStoreByPlaceId(ctx, store.PlaceID)

				if err != nil {
					if err == sql.ErrNoRows {
						arg := db.CreateStoreParams{
							Name:    store.Name,
							XCoord:  store.XCoord,
							YCoord:  store.YCoord,
							PlaceID: store.PlaceID,
							Address: store.Address,
						}
						newStore, err := server.queries.CreateStore(ctx, arg)
						if err == nil {
							addCommunityStore(ctx, server, community.ID, newStore.ID)
						}
					}
				} else {
					addCommunityStore(ctx, server, community.ID, existingStore.ID)
				}

			}

			ctx.JSON(http.StatusOK, community)

		} else {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		}
	}

	ctx.JSON(http.StatusBadRequest, errorResponse(api_error.ErrCommunityExists, errors.New("community already exists")))
}

func addCommunityStore(ctx *gin.Context, server *Server, communityID int64, storeID int64) {
	arg := db.CreateCommunityStoreParams{
		CommunityID: communityID,
		StoreID:     storeID,
	}
	server.queries.CreateCommunityStore(ctx, arg)
}
