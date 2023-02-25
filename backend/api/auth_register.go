package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	api_error "github.com/git-adithyanair/cs130-group-project/errors"
	"github.com/git-adithyanair/cs130-group-project/util"
)

type RegisterUserRequest struct {
	Email       string  `json:"email" binding:"required,email"`
	Password    string  `json:"password" binding:"required,min=6"`
	FullName    string  `json:"full_name" binding:"required"`
	PhoneNumber string  `json:"phone_number" binding:"required,numeric"`
	Address     string  `json:"address" binding:"required"`
	PlaceID     string  `json:"place_id" binding:"required"`
	XCoord      float64 `json:"x_coord" binding:"required,numeric"`
	YCoord      float64 `json:"y_coord" binding:"required,numeric"`
}

func (server *Server) RegisterUser(ctx *gin.Context) {
	var req RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, unknownErrorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Email:          req.Email,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		PhoneNumber:    req.PhoneNumber,
		Address:        req.Address,
		PlaceID:        req.PlaceID,
		XCoord:         req.XCoord,
		YCoord:         req.YCoord,
	}

	user, err := server.queries.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(api_error.ErrRegisterFail, err))
		return
	}

	if err = util.NotifyUser(user.PhoneNumber, "Thank you for registering for GoodGrocer! You will recieve updates about your requests and errands through this number as you use the app!"); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(api_error.ErrTwilio, err))
		return
	}

	token, err := server.tokenMaker.CreateToken(user.ID, user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, unknownErrorResponse(err))
		return
	}

	res := newAuthUserResponse(token, user)
	ctx.JSON(http.StatusOK, res)

}
