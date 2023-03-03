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
	ID             int64     `json:"id"`
	Email          string    `json:"email"`
	FullName       string    `json:"full_name"`
	CreatedAt      time.Time `json:"created_at"`
	ProfilePicture string    `json:"profile_picture"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		ID:             user.ID,
		FullName:       user.FullName,
		Email:          user.Email,
		CreatedAt:      user.CreatedAt,
		ProfilePicture: user.ProfilePicture,
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

// ========================================================================
// Type and function to return users requests stored by status.
type userRequestsResponse struct {
	Pending    []db.Request `json:"pending"`
	InProgress []db.Request `json:"in_progress"`
	Complete   []db.Request `json:"complete"`
}

// ========================================================================

// ========================================================================
// Type and function to return only detailed unprotected user information.
// To be used during an errand for shopper to contact user.
type userDetailedResponse struct {
	ID             int64   `json:"id"`
	Email          string  `json:"email"`
	FullName       string  `json:"full_name"`
	PhoneNumber    string  `json:"phone_number"`
	XCoord         float64 `json:"x_coord"`
	YCoord         float64 `json:"y_cord"`
	Address        string  `json:"address"`
	ProfilePicture string  `json:"profile_picture"`
}

func newUserDetailedResponse(user db.User) userDetailedResponse {
	return userDetailedResponse{
		ID:             user.ID,
		FullName:       user.FullName,
		Email:          user.Email,
		PhoneNumber:    user.PhoneNumber,
		XCoord:         user.XCoord,
		YCoord:         user.YCoord,
		Address:        user.Address,
		ProfilePicture: user.ProfilePicture,
	}
}

// ========================================================================

// ========================================================================
// Type and function to return users active errand with its requests.
type activeErrandResponse struct {
	Errand   db.Errand                     `json:"errand"`
	Requests []activeErrandRequestResponse `json:"requests"`
}

type activeErrandRequestResponse struct {
	Request db.Request           `json:"request"`
	Items   []db.Item            `json:"items"`
	User    userDetailedResponse `json:"user"`
	Store   *db.Store            `json:"store"`
}

// ========================================================================

// ========================================================================
// Type and function to return info for a request and it's associated user and store
type communityRequestsResponse struct {
	Request db.Request   `json:"request"`
	User    userResponse `json:"user"`
	Store   *db.Store    `json:"store"`
}

// ========================================================================
