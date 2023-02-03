package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	"github.com/git-adithyanair/cs130-group-project/util"
)

type RegisterUserRequest struct {
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=6"`
	FullName     string `json:"full_name" binding:"required"`
	PhoneNumber  string `json:"phone_number" binding:"required,numeric"`
	AddressLine1 string `json:"address_line_1" binding:"required"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city" binding:"required"`
	State        string `json:"state" binding:"required,alpha,len=2"`
	ZipCode      string `json:"zip_code" binding:"required,numeric,len=5"`
}

func (server *Server) RegisterUser(ctx *gin.Context) {
	var req RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	addressLine2 := sql.NullString{String: req.AddressLine2, Valid: true}
	if req.AddressLine2 == "" {
		addressLine2 = sql.NullString{String: "", Valid: false}
	}

	arg := db.CreateUserParams{
		Email:          req.Email,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		PhoneNumber:    req.PhoneNumber,
		AddressLine1:   req.AddressLine1,
		AddressLine2:   addressLine2,
		City:           req.City,
		State:          req.State,
		ZipCode:        req.ZipCode,
	}

	user, err := server.queries.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := newUserResponse(user)
	ctx.JSON(http.StatusOK, res)

}
