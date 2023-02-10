package api

import (
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	api_error "github.com/git-adithyanair/cs130-group-project/errors"
)

// ========================================================================
// Function to convert errors into a form that can be sent as a response.
func errorResponse(errCode string, err error) gin.H {
	errMessage := api_error.GetErrorMessage[api_error.ErrUnknown]
	val, ok := api_error.GetErrorMessage[errCode]
	if ok {
		errMessage = val
	}
	return gin.H{
		"id":    errCode,
		"error": errMessage,
		"raw":   err.Error(),
	}
}

func unknownErrorResponse(err error) gin.H {
	return errorResponse(api_error.ErrUnknown, err)
}

func authErrorResponse(err error) gin.H {
	return errorResponse(api_error.ErrAuthFail, err)
}

// ========================================================================

// ========================================================================
// Type and function to return only unprotected user information.
type userResponse struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

// ========================================================================

// ========================================================================
// Type and function to return auth information.
type authUserResponse struct {
	Token string       `json:"token"`
	User  userResponse `json:"user"`
}

func newAuthUserResponse(token string, user db.User) authUserResponse {
	return authUserResponse{
		Token: token,
		User:  newUserResponse(user),
	}
}

// ========================================================================
