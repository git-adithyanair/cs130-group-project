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

func TestJoinCommunity(t *testing.T) {

	user, _ := createRandomUser(t)
	community := createRandomCommunity(t, user.ID)
	member := createRandomMember(t, user.ID, community.ID)

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
			name: "OK",
			body: gin.H{
				"id": community.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), gomock.Eq(community.ID)).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					CreateMember(gomock.Any(), db.CreateMemberParams{
						UserID:      user.ID,
						CommunityID: community.ID,
					}).
					Times(1).
					Return(member, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCommunity(t, recorder.Body, community)
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
			name: "NotFound",
			body: gin.H{
				"id": community.ID + 100,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID+100).
					Times(1).
					Return(db.Community{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalServerErrorOnGetCommunity",
			body: gin.H{
				"id": community.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(1).
					Return(db.Community{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InternalServerErrorOnCreateMember",
			body: gin.H{
				"id": community.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), gomock.Eq(community.ID)).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					CreateMember(gomock.Any(), db.CreateMemberParams{
						UserID:      user.ID,
						CommunityID: community.ID,
					}).
					Times(1).
					Return(db.Member{}, sql.ErrConnDone)
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
			url := "/community/join"
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
