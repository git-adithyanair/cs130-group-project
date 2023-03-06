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

func TestCreateErrand(t *testing.T) {

	shopper, _ := createRandomUser(t)
	buyer1, _ := createRandomUser(t)
	buyer2, _ := createRandomUser(t)
	community := createRandomCommunity(t, shopper.ID)
	groceryStore := createRandomStore(t)
	createRandomCommunityStore(t, community.ID, groceryStore.ID)

	request1 := createRandomRequest(t, buyer1.ID, community.ID, groceryStore.ID)
	request2 := createRandomRequest(t, buyer2.ID, community.ID, groceryStore.ID)
	requestInErrand := createRandomRequest(t, buyer2.ID, community.ID, groceryStore.ID)
	requestByShopper := createRandomRequest(t, shopper.ID, community.ID, groceryStore.ID)

	errand := createRandomErrand(t, shopper.ID, community.ID)
	request1.ErrandID = sql.NullInt64{Int64: errand.ID, Valid: false}
	request2.ErrandID = sql.NullInt64{Int64: errand.ID, Valid: false}
	requestInErrand.ErrandID = sql.NullInt64{Int64: errand.ID, Valid: true}

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
				"request_ids":  []int64{request1.ID, request2.ID},
				"community_id": community.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetUser(gomock.Any(), shopper.ID).Times(1).Return(shopper, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(1).Return(buyer1, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(1).Return(buyer2, nil)
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(db.Errand{}, sql.ErrNoRows)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(1).Return(community, nil)
				store.EXPECT().GetRequest(gomock.Any(), request1.ID).Times(2).Return(request1, nil)
				store.EXPECT().GetRequest(gomock.Any(), request2.ID).Times(2).Return(request2, nil)
				store.EXPECT().CreateErrand(gomock.Any(), db.CreateErrandParams{
					CommunityID: community.ID,
					UserID:      shopper.ID,
				}).Times(1).Return(errand, nil)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request1.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(1)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request2.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(1)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
			},
		},
		{
			name: "MissingCommunityID",
			body: gin.H{
				"request_ids": []int64{request1.ID, request2.ID},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetUser(gomock.Any(), shopper.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(0)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(0)
				store.EXPECT().GetRequest(gomock.Any(), request1.ID).Times(0)
				store.EXPECT().GetRequest(gomock.Any(), request2.ID).Times(0)
				store.EXPECT().CreateErrand(gomock.Any(), db.CreateErrandParams{
					CommunityID: community.ID,
					UserID:      shopper.ID,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request1.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request2.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "NoShopperWithGivenID",
			body: gin.H{
				"request_ids":  []int64{request1.ID, request2.ID},
				"community_id": community.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetUser(gomock.Any(), shopper.ID).Times(1).Return(db.User{}, sql.ErrConnDone)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(0)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(0)
				store.EXPECT().GetRequest(gomock.Any(), request1.ID).Times(0)
				store.EXPECT().GetRequest(gomock.Any(), request2.ID).Times(0)
				store.EXPECT().CreateErrand(gomock.Any(), db.CreateErrandParams{
					CommunityID: community.ID,
					UserID:      shopper.ID,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request1.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request2.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "ShopperHasActiveErrand",
			body: gin.H{
				"request_ids":  []int64{request1.ID, request2.ID},
				"community_id": community.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetUser(gomock.Any(), shopper.ID).Times(1).Return(shopper, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(errand, nil)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(0)
				store.EXPECT().GetRequest(gomock.Any(), request1.ID).Times(0)
				store.EXPECT().GetRequest(gomock.Any(), request2.ID).Times(0)
				store.EXPECT().CreateErrand(gomock.Any(), db.CreateErrandParams{
					CommunityID: community.ID,
					UserID:      shopper.ID,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request1.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request2.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusExpectationFailed, recorder.Code)
			},
		},
		{
			name: "FailToGetShopperActiveErrand",
			body: gin.H{
				"request_ids":  []int64{request1.ID, request2.ID},
				"community_id": community.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetUser(gomock.Any(), shopper.ID).Times(1).Return(shopper, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(db.Errand{}, sql.ErrConnDone)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(0)
				store.EXPECT().GetRequest(gomock.Any(), request1.ID).Times(0)
				store.EXPECT().GetRequest(gomock.Any(), request2.ID).Times(0)
				store.EXPECT().CreateErrand(gomock.Any(), db.CreateErrandParams{
					CommunityID: community.ID,
					UserID:      shopper.ID,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request1.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request2.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "NoCommunityWithGivenID",
			body: gin.H{
				"request_ids":  []int64{request1.ID, request2.ID},
				"community_id": community.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetUser(gomock.Any(), shopper.ID).Times(1).Return(shopper, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(db.Errand{}, sql.ErrNoRows)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(1).Return(db.Community{}, sql.ErrNoRows)
				store.EXPECT().GetRequest(gomock.Any(), request1.ID).Times(0)
				store.EXPECT().GetRequest(gomock.Any(), request2.ID).Times(0)
				store.EXPECT().CreateErrand(gomock.Any(), db.CreateErrandParams{
					CommunityID: community.ID,
					UserID:      shopper.ID,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request1.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request2.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "FailToGetCommunity",
			body: gin.H{
				"request_ids":  []int64{request1.ID, request2.ID},
				"community_id": community.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetUser(gomock.Any(), shopper.ID).Times(1).Return(shopper, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(db.Errand{}, sql.ErrNoRows)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(1).Return(db.Community{}, sql.ErrConnDone)
				store.EXPECT().GetRequest(gomock.Any(), request1.ID).Times(0)
				store.EXPECT().GetRequest(gomock.Any(), request2.ID).Times(0)
				store.EXPECT().CreateErrand(gomock.Any(), db.CreateErrandParams{
					CommunityID: community.ID,
					UserID:      shopper.ID,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request1.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request2.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "NoRequestIDs",
			body: gin.H{
				"request_ids":  []int64{},
				"community_id": community.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetUser(gomock.Any(), shopper.ID).Times(1).Return(shopper, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(db.Errand{}, sql.ErrNoRows)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(1).Return(community, nil)
				store.EXPECT().GetRequest(gomock.Any(), request1.ID).Times(0)
				store.EXPECT().GetRequest(gomock.Any(), request2.ID).Times(0)
				store.EXPECT().CreateErrand(gomock.Any(), db.CreateErrandParams{
					CommunityID: community.ID,
					UserID:      shopper.ID,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request1.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request2.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "RequestNotFound",
			body: gin.H{
				"request_ids":  []int64{request1.ID, request2.ID},
				"community_id": community.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetUser(gomock.Any(), shopper.ID).Times(1).Return(shopper, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(db.Errand{}, sql.ErrNoRows)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(1).Return(community, nil)
				store.EXPECT().GetRequest(gomock.Any(), request1.ID).Times(1).Return(db.Request{}, sql.ErrNoRows)
				store.EXPECT().GetRequest(gomock.Any(), request2.ID).Times(0)
				store.EXPECT().CreateErrand(gomock.Any(), db.CreateErrandParams{
					CommunityID: community.ID,
					UserID:      shopper.ID,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request1.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request2.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "FailToGetRequest",
			body: gin.H{
				"request_ids":  []int64{request1.ID, request2.ID},
				"community_id": community.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetUser(gomock.Any(), shopper.ID).Times(1).Return(shopper, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(db.Errand{}, sql.ErrNoRows)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(1).Return(community, nil)
				store.EXPECT().GetRequest(gomock.Any(), request1.ID).Times(1).Return(db.Request{}, sql.ErrConnDone)
				store.EXPECT().GetRequest(gomock.Any(), request2.ID).Times(0)
				store.EXPECT().CreateErrand(gomock.Any(), db.CreateErrandParams{
					CommunityID: community.ID,
					UserID:      shopper.ID,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request1.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request2.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "RequestAlreadyInErrand",
			body: gin.H{
				"request_ids":  []int64{request1.ID, request2.ID, requestInErrand.ID},
				"community_id": community.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetUser(gomock.Any(), shopper.ID).Times(1).Return(shopper, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(db.Errand{}, sql.ErrNoRows)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(1).Return(community, nil)
				store.EXPECT().GetRequest(gomock.Any(), request1.ID).Times(1).Return(request1, nil)
				store.EXPECT().GetRequest(gomock.Any(), request2.ID).Times(1).Return(request2, nil)
				store.EXPECT().GetRequest(gomock.Any(), requestInErrand.ID).Times(1).Return(requestInErrand, nil)
				store.EXPECT().CreateErrand(gomock.Any(), db.CreateErrandParams{
					CommunityID: community.ID,
					UserID:      shopper.ID,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request1.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request2.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusExpectationFailed, recorder.Code)
			},
		},
		{
			name: "RequestBelongsToShopper",
			body: gin.H{
				"request_ids":  []int64{request1.ID, request2.ID, requestByShopper.ID},
				"community_id": community.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetUser(gomock.Any(), shopper.ID).Times(1).Return(shopper, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(db.Errand{}, sql.ErrNoRows)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(1).Return(community, nil)
				store.EXPECT().GetRequest(gomock.Any(), request1.ID).Times(1).Return(request1, nil)
				store.EXPECT().GetRequest(gomock.Any(), request2.ID).Times(1).Return(request2, nil)
				store.EXPECT().GetRequest(gomock.Any(), requestByShopper.ID).Times(1).Return(requestByShopper, nil)
				store.EXPECT().CreateErrand(gomock.Any(), db.CreateErrandParams{
					CommunityID: community.ID,
					UserID:      shopper.ID,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request1.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request2.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusExpectationFailed, recorder.Code)
			},
		},
		{
			name: "FailedToCreateErrand",
			body: gin.H{
				"request_ids":  []int64{request1.ID, request2.ID},
				"community_id": community.ID,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetUser(gomock.Any(), shopper.ID).Times(1).Return(shopper, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(db.Errand{}, sql.ErrNoRows)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(1).Return(community, nil)
				store.EXPECT().GetRequest(gomock.Any(), request1.ID).Times(1).Return(request1, nil)
				store.EXPECT().GetRequest(gomock.Any(), request2.ID).Times(1).Return(request2, nil)
				store.EXPECT().CreateErrand(gomock.Any(), db.CreateErrandParams{
					CommunityID: community.ID,
					UserID:      shopper.ID,
				}).Times(1).Return(db.Errand{}, sql.ErrConnDone)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request1.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
				store.EXPECT().UpdateRequestErrandAndStatus(gomock.Any(), db.UpdateRequestErrandAndStatusParams{
					ID:       request2.ID,
					ErrandID: sql.NullInt64{Int64: errand.ID, Valid: true},
					Status:   db.RequestStatusInProgress,
				}).Times(0)
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
			url := "/errand"
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
