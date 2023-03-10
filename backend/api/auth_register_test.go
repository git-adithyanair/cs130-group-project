package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	mock_db "github.com/git-adithyanair/cs130-group-project/db/mock"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	"github.com/git-adithyanair/cs130-group-project/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserParams
	password string
}

// Create custom matcher to verify passwords.
func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {

	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	err := util.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}

	e.arg.HashedPassword = arg.HashedPassword

	return reflect.DeepEqual(e.arg, arg)

}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func TestRegisterUser(t *testing.T) {

	user, rawPassword := createRandomUser(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_db.NewMockDBStore(ctrl)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mock_db.MockDBStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"email":           user.Email,
				"password":        rawPassword,
				"full_name":       user.FullName,
				"phone_number":    user.PhoneNumber,
				"address":         user.Address,
				"place_id":        user.PlaceID,
				"x_coord":         user.XCoord,
				"y_coord":         user.YCoord,
				"profile_picture": "PROFILE_PICTURE",
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				arg := db.CreateUserParams{
					Email:       user.Email,
					FullName:    user.FullName,
					PhoneNumber: user.PhoneNumber,
					PlaceID:     user.PlaceID,
					XCoord:      user.XCoord,
					YCoord:      user.YCoord,
					Address:     user.Address,
				}
				store.EXPECT().
					CreateUser(gomock.Any(), EqCreateUserParams(arg, rawPassword)).
					Times(1).
					Return(user, nil)
				store.EXPECT().
					UpdateUserProfilePicture(gomock.Any(), db.UpdateUserProfilePictureParams{
						ID:             user.ID,
						ProfilePicture: "PROFILE_PICTURE",
					}).Times(1).Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)

			},
		},
		{
			name: "InvalidEmail",
			body: gin.H{
				"email":           "invalid_email",
				"password":        rawPassword,
				"full_name":       user.FullName,
				"phone_number":    user.PhoneNumber,
				"address":         user.Address,
				"place_id":        user.PlaceID,
				"x_coord":         user.XCoord,
				"y_coord":         user.YCoord,
				"profile_picture": "DEFAULT",
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "PasswordTooShort",
			body: gin.H{
				"email":           user.Email,
				"password":        "short",
				"full_name":       user.FullName,
				"phone_number":    user.PhoneNumber,
				"address":         user.Address,
				"place_id":        user.PlaceID,
				"x_coord":         user.XCoord,
				"y_coord":         user.YCoord,
				"profile_picture": "DEFAULT",
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalServerError",
			body: gin.H{
				"email":           user.Email,
				"password":        rawPassword,
				"full_name":       user.FullName,
				"phone_number":    user.PhoneNumber,
				"address":         user.Address,
				"place_id":        user.PlaceID,
				"x_coord":         user.XCoord,
				"y_coord":         user.YCoord,
				"profile_picture": "DEFAULT",
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				arg := db.CreateUserParams{
					Email:       user.Email,
					FullName:    user.FullName,
					PhoneNumber: user.PhoneNumber,
					PlaceID:     user.PlaceID,
					XCoord:      user.XCoord,
					YCoord:      user.YCoord,
					Address:     user.Address,
				}
				store.EXPECT().
					CreateUser(gomock.Any(), EqCreateUserParams(arg, rawPassword)).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			// Build stubs.
			testCase.buildStubs(store)

			// Start the test server and send request.
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Construct the request.
			url := "/user"
			jsonBody, err := json.Marshal(testCase.body)
			require.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
			require.NoError(t, err)

			// Send the request and record result in recorder.
			server.router.ServeHTTP(recorder, request)

			// Check the response.
			testCase.checkResponse(t, recorder)
		})

	}

}
