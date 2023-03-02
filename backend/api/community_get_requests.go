package api

import (
	"database/sql"
	"net/http"

	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"

	api_error "github.com/git-adithyanair/cs130-group-project/errors"
	"github.com/git-adithyanair/cs130-group-project/token"

	"github.com/gin-gonic/gin"
)

type GetRequestsByCommunityRequest struct {
	ID     int64  `form:"id" binding:"required,min=1"`
	Limit  *int32 `form:"limit"`
	Offset *int32 `form:"offset"`
}

func (server *Server) GetRequestsByCommunity(ctx *gin.Context) {
	var req GetRequestsByCommunityRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, unknownErrorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	community, err := server.queries.GetCommunity(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoCommunity, err))
		} else {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		}
		return
	}

	getMemberArg := db.GetMemberParams{
		UserID:      authPayload.UserID,
		CommunityID: community.ID,
	}

	_, err = server.queries.GetMember(ctx, getMemberArg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoMember, err))
		} else {
			ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		}
		return
	}

	arg := db.GetRequestsByCommunityIdParams{
		CommunityID: sql.NullInt64{Int64: req.ID, Valid: true},
		Limit:       10,
		Offset:      0,
	}
	if req.Offset != nil {
		arg.Offset = *req.Offset
	}
	if req.Limit != nil {
		arg.Limit = *req.Limit
	}

	requests, err := server.queries.GetRequestsByCommunityId(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		return
	}

	requestRes := make([]communityRequestsResponse, len(requests))
	for i, request := range requests {
		user, err := server.queries.GetUser(ctx, request.UserID)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoUser, err))
			} else {
				ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
			}
			return
		}

		requestRes[i] = communityRequestsResponse{
			Request: request,
			User:    newUserResponse(user),
		}

		if request.StoreID.Valid {
			store, err := server.queries.GetStore(ctx, request.StoreID.Int64)
			if err != nil {
				if err == sql.ErrNoRows {
					ctx.JSON(http.StatusNotFound, errorResponse(api_error.ErrNoStore, err))
				} else {
					ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
				}
				return
			}

			requestRes[i].Store = &store
		}
	}

	ctx.JSON(http.StatusOK, requestRes)
}
