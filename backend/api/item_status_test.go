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
	"github.com/git-adithyanair/cs130-group-project/token"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUpdateItemStatus(t *testing.T) {

	user, _ := createRandomUser(t)
	item := createRandomItemWithUser(t, user.ID)
	item_found, item_not_found := item, item
	item_found.Found = sql.NullBool{Bool: true, Valid: true}
	item_not_found.Found = sql.NullBool{Bool: false, Valid: true}

	item_invalid_user := item
	item_invalid_user.RequestedBy = user.ID + 100

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_db.NewMockDBStore(ctrl)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mock_db.MockDBStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OKFound",
			body: gin.H{
				"id":    item.ID,
				"found": true,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					UpdateItemFound(gomock.Any(), db.UpdateItemFoundParams{
						ID:    item.ID,
						Found: sql.NullBool{Bool: true, Valid: true},
					}).
					Times(1).
					Return(item_found, nil)
				store.EXPECT().
					GetUser(gomock.Any(), user.ID).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchItemWithStatus(t, recorder.Body, item_found, true)
			},
		},
		{
			name: "OKFound",
			body: gin.H{
				"id":    item.ID,
				"found": false,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					UpdateItemFound(gomock.Any(), db.UpdateItemFoundParams{
						ID:    item.ID,
						Found: sql.NullBool{Bool: false, Valid: true},
					}).
					Times(1).
					Return(item_not_found, nil)
				store.EXPECT().
					GetUser(gomock.Any(), user.ID).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchItemWithStatus(t, recorder.Body, item_not_found, false)
			},
		},
		{
			name: "BadRequest",
			body: gin.H{},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "ItemNotFound",
			body: gin.H{
				"id":    item.ID + 100,
				"found": false,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					UpdateItemFound(gomock.Any(), db.UpdateItemFoundParams{
						ID:    item.ID + 100,
						Found: sql.NullBool{Bool: false, Valid: true},
					}).
					Times(1).
					Return(db.Item{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "ItemInternalServerError",
			body: gin.H{
				"id":    item.ID,
				"found": false,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					UpdateItemFound(gomock.Any(), db.UpdateItemFoundParams{
						ID:    item.ID,
						Found: sql.NullBool{Bool: false, Valid: true},
					}).
					Times(1).
					Return(db.Item{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "UserNotFound",
			body: gin.H{
				"id":    item_invalid_user.ID,
				"found": false,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					UpdateItemFound(gomock.Any(), db.UpdateItemFoundParams{
						ID:    item.ID,
						Found: sql.NullBool{Bool: false, Valid: true},
					}).
					Times(1).
					Return(item_invalid_user, nil)
				store.EXPECT().
					GetUser(gomock.Any(), user.ID+100).
					Times(1).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "UserInternalServerError",
			body: gin.H{
				"id":    item.ID,
				"found": false,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					UpdateItemFound(gomock.Any(), db.UpdateItemFoundParams{
						ID:    item.ID,
						Found: sql.NullBool{Bool: false, Valid: true},
					}).
					Times(1).
					Return(item_not_found, nil)
				store.EXPECT().
					GetUser(gomock.Any(), user.ID).
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
			url := "/item/update-status"
			jsonBody, err := json.Marshal(testCase.body)
			require.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))
			require.NoError(t, err)

			// Set up auth header.
			testCase.setupAuth(t, request, server.tokenMaker)

			// Send the request and record result in recorder.
			server.router.ServeHTTP(recorder, request)

			// Check the response.
			testCase.checkResponse(t, recorder)
		})

	}

}
