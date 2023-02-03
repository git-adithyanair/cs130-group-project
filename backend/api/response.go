package api

import (
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
)

// ========================================================================
// Function to convert errors into a form that can be sent as a response.
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
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
// Type to return only unprotected user information.
type registerUserResponse struct {
	Token string       `json:"token"`
	User  userResponse `json:"user"`
}

func newRegisterUserResponse(token string, user db.User) registerUserResponse {
	return registerUserResponse{
		Token: token,
		User:  newUserResponse(user),
	}
}

// ========================================================================
