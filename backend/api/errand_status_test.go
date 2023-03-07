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

func TestErrandStatus(t *testing.T) {

	shopper, _ := createRandomUser(t)
	buyer1, _ := createRandomUser(t)
	buyer2, _ := createRandomUser(t)
	community := createRandomCommunity(t, shopper.ID)
	groceryStore := createRandomStore(t)
	createRandomCommunityStore(t, community.ID, groceryStore.ID)

	request1 := createRandomRequest(t, buyer1.ID, community.ID, groceryStore.ID)
	request2 := createRandomRequest(t, buyer2.ID, community.ID, groceryStore.ID)

	errand := createRandomErrand(t, shopper.ID, community.ID)
	request1.ErrandID = sql.NullInt64{Int64: errand.ID, Valid: false}
	request2.ErrandID = sql.NullInt64{Int64: errand.ID, Valid: false}

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
			name: "OKStatusTrue",
			body: gin.H{
				"id":          errand.ID,
				"is_complete": true,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().UpdateErrandStatus(gomock.Any(), db.UpdateErrandStatusParams{
					ID:         errand.ID,
					IsComplete: true,
				}).Times(1).Return(errand, nil)
				store.EXPECT().GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).
					Times(1).Return([]db.Request{request1, request2}, nil)
				store.EXPECT().UpdateRequestStatus(gomock.Any(), db.UpdateRequestStatusParams{
					ID:     request1.ID,
					Status: db.RequestStatusCompleted,
				}).Times(1).Return(request1, nil)
				store.EXPECT().UpdateRequestStatus(gomock.Any(), db.UpdateRequestStatusParams{
					ID:     request2.ID,
					Status: db.RequestStatusCompleted,
				}).Times(1).Return(request2, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(1).Return(buyer1, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(1).Return(buyer2, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "OKStatusFalse",
			body: gin.H{
				"id":          errand.ID,
				"is_complete": false,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().UpdateErrandStatus(gomock.Any(), db.UpdateErrandStatusParams{
					ID:         errand.ID,
					IsComplete: false,
				}).Times(1).Return(errand, nil)
				store.EXPECT().GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).
					Times(0)
				store.EXPECT().UpdateRequestStatus(gomock.Any(), db.UpdateRequestStatusParams{
					ID:     request1.ID,
					Status: db.RequestStatusCompleted,
				}).Times(0)
				store.EXPECT().UpdateRequestStatus(gomock.Any(), db.UpdateRequestStatusParams{
					ID:     request2.ID,
					Status: db.RequestStatusCompleted,
				}).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			body: gin.H{
				"is_complete": true,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().UpdateErrandStatus(gomock.Any(), db.UpdateErrandStatusParams{
					ID:         errand.ID,
					IsComplete: true,
				}).Times(0)
				store.EXPECT().GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).
					Times(0)
				store.EXPECT().UpdateRequestStatus(gomock.Any(), db.UpdateRequestStatusParams{
					ID:     request1.ID,
					Status: db.RequestStatusCompleted,
				}).Times(0)
				store.EXPECT().UpdateRequestStatus(gomock.Any(), db.UpdateRequestStatusParams{
					ID:     request2.ID,
					Status: db.RequestStatusCompleted,
				}).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "ErrandNotFound",
			body: gin.H{
				"id":          errand.ID,
				"is_complete": true,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().UpdateErrandStatus(gomock.Any(), db.UpdateErrandStatusParams{
					ID:         errand.ID,
					IsComplete: true,
				}).Times(1).Return(db.Errand{}, sql.ErrNoRows)
				store.EXPECT().GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).
					Times(0)
				store.EXPECT().UpdateRequestStatus(gomock.Any(), db.UpdateRequestStatusParams{
					ID:     request1.ID,
					Status: db.RequestStatusCompleted,
				}).Times(0)
				store.EXPECT().UpdateRequestStatus(gomock.Any(), db.UpdateRequestStatusParams{
					ID:     request2.ID,
					Status: db.RequestStatusCompleted,
				}).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "FailToUpdateErrand",
			body: gin.H{
				"id":          errand.ID,
				"is_complete": true,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().UpdateErrandStatus(gomock.Any(), db.UpdateErrandStatusParams{
					ID:         errand.ID,
					IsComplete: true,
				}).Times(1).Return(db.Errand{}, sql.ErrConnDone)
				store.EXPECT().GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).
					Times(0)
				store.EXPECT().UpdateRequestStatus(gomock.Any(), db.UpdateRequestStatusParams{
					ID:     request1.ID,
					Status: db.RequestStatusCompleted,
				}).Times(0)
				store.EXPECT().UpdateRequestStatus(gomock.Any(), db.UpdateRequestStatusParams{
					ID:     request2.ID,
					Status: db.RequestStatusCompleted,
				}).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "ErrandRequestsNotFound",
			body: gin.H{
				"id":          errand.ID,
				"is_complete": true,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().UpdateErrandStatus(gomock.Any(), db.UpdateErrandStatusParams{
					ID:         errand.ID,
					IsComplete: true,
				}).Times(1).Return(errand, nil)
				store.EXPECT().GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).
					Times(1).Return([]db.Request{}, sql.ErrNoRows)
				store.EXPECT().UpdateRequestStatus(gomock.Any(), db.UpdateRequestStatusParams{
					ID:     request1.ID,
					Status: db.RequestStatusCompleted,
				}).Times(0)
				store.EXPECT().UpdateRequestStatus(gomock.Any(), db.UpdateRequestStatusParams{
					ID:     request2.ID,
					Status: db.RequestStatusCompleted,
				}).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "FailToGetErrandRequests",
			body: gin.H{
				"id":          errand.ID,
				"is_complete": true,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().UpdateErrandStatus(gomock.Any(), db.UpdateErrandStatusParams{
					ID:         errand.ID,
					IsComplete: true,
				}).Times(1).Return(errand, nil)
				store.EXPECT().GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).
					Times(1).Return([]db.Request{request1, request2}, sql.ErrConnDone)
				store.EXPECT().UpdateRequestStatus(gomock.Any(), db.UpdateRequestStatusParams{
					ID:     request1.ID,
					Status: db.RequestStatusCompleted,
				}).Times(0)
				store.EXPECT().UpdateRequestStatus(gomock.Any(), db.UpdateRequestStatusParams{
					ID:     request2.ID,
					Status: db.RequestStatusCompleted,
				}).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
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
			url := "/errand/update-status"
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
