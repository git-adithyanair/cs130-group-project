package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	mock_db "github.com/git-adithyanair/cs130-group-project/db/mock"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestLoginUser(t *testing.T) {

	user, rawPassword := randomUser(t)

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
				"email":    user.Email,
				"password": rawPassword,
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Eq(user.Email)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name: "InvalidEmail",
			body: gin.H{
				"email":    "invalid_email",
				"password": rawPassword,
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "PasswordTooShort",
			body: gin.H{
				"email":    user.Email,
				"password": "short",
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "UserNotFound",
			body: gin.H{
				"email":    user.Email,
				"password": rawPassword,
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Eq(user.Email)).
					Times(1).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalServerError",
			body: gin.H{
				"email":    user.Email,
				"password": rawPassword,
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Eq(user.Email)).
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
			url := "/user/login"
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
