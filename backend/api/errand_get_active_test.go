package api

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_db "github.com/git-adithyanair/cs130-group-project/db/mock"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	"github.com/git-adithyanair/cs130-group-project/token"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetActiveErrand(t *testing.T) {

	shopper, _ := createRandomUser(t)
	buyer1, _ := createRandomUser(t)
	buyer2, _ := createRandomUser(t)
	community := createRandomCommunity(t, shopper.ID)
	groceryStore := createRandomStore(t)

	request1 := createRandomRequest(t, buyer1.ID, community.ID, groceryStore.ID)
	request2 := createRandomRequestWithNoStore(t, buyer2.ID, community.ID)

	request1Items := []db.Item{
		createRandomItem(t, buyer1.ID, request1.ID, "oz", 1, false, true, true),
		createRandomItem(t, buyer1.ID, request1.ID, "fl_oz", 1, true, false, true),
	}
	request2Items := []db.Item{
		createRandomItem(t, buyer2.ID, request2.ID, "numerical", 1, false, true, true),
		createRandomItem(t, buyer2.ID, request2.ID, "lbs", 1, true, false, true),
	}

	errand := createRandomErrand(t, shopper.ID, community.ID)
	request1.ErrandID = sql.NullInt64{Int64: errand.ID, Valid: true}
	request2.ErrandID = sql.NullInt64{Int64: errand.ID, Valid: true}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_db.NewMockDBStore(ctrl)

	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mock_db.MockDBStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(errand, nil)
				store.EXPECT().GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).
					Times(1).Return([]db.Request{request1, request2}, nil)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request1.ID).Times(1).Return(request1Items, nil)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request2.ID).Times(1).Return(request2Items, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(1).Return(buyer1, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(1).Return(buyer2, nil)
				store.EXPECT().GetStore(gomock.Any(), request1.StoreID.Int64).Times(1).Return(groceryStore, nil)
				store.EXPECT().GetStore(gomock.Any(), request2.StoreID.Int64).Times(0)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(1).Return(community, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "NoActiveErrand",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(db.Errand{}, sql.ErrNoRows)
				store.EXPECT().GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).Times(0)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request1.ID).Times(0)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request2.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetStore(gomock.Any(), request1.StoreID.Int64).Times(0)
				store.EXPECT().GetStore(gomock.Any(), request2.StoreID.Int64).Times(0)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				data, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)
				var response struct{}
				err = json.Unmarshal(data, &response)
				require.NoError(t, err)
				require.Empty(t, response)
			},
		},
		{
			name: "FailToGetActiveErrand",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(db.Errand{}, sql.ErrConnDone)
				store.EXPECT().GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).Times(0)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request1.ID).Times(0)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request2.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetStore(gomock.Any(), request1.StoreID.Int64).Times(0)
				store.EXPECT().GetStore(gomock.Any(), request2.StoreID.Int64).Times(0)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "FailToGetRequests",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(errand, nil)
				store.EXPECT().
					GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).
					Times(1).
					Return(nil, sql.ErrConnDone)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request1.ID).Times(0)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request2.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetStore(gomock.Any(), request1.StoreID.Int64).Times(0)
				store.EXPECT().GetStore(gomock.Any(), request2.StoreID.Int64).Times(0)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "NoItemsInRequest",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(errand, nil)
				store.EXPECT().
					GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).
					Times(1).
					Return([]db.Request{request1, request2}, nil)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request1.ID).Times(1).Return([]db.Item{}, sql.ErrNoRows)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request2.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetStore(gomock.Any(), request1.StoreID.Int64).Times(0)
				store.EXPECT().GetStore(gomock.Any(), request2.StoreID.Int64).Times(0)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "FailToGetItemsInRequest",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(errand, nil)
				store.EXPECT().
					GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).
					Times(1).
					Return([]db.Request{request1, request2}, nil)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request1.ID).Times(1).Return([]db.Item{}, sql.ErrConnDone)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request2.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetStore(gomock.Any(), request1.StoreID.Int64).Times(0)
				store.EXPECT().GetStore(gomock.Any(), request2.StoreID.Int64).Times(0)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "BuyerNotFound",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(errand, nil)
				store.EXPECT().
					GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).
					Times(1).
					Return([]db.Request{request1, request2}, nil)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request1.ID).Times(1).Return(request1Items, nil)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request2.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(1).Return(db.User{}, sql.ErrNoRows)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetStore(gomock.Any(), request1.StoreID.Int64).Times(0)
				store.EXPECT().GetStore(gomock.Any(), request2.StoreID.Int64).Times(0)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "FailToGetBuyer",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(errand, nil)
				store.EXPECT().
					GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).
					Times(1).
					Return([]db.Request{request1, request2}, nil)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request1.ID).Times(1).Return(request1Items, nil)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request2.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(1).Return(db.User{}, sql.ErrConnDone)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetStore(gomock.Any(), request1.StoreID.Int64).Times(0)
				store.EXPECT().GetStore(gomock.Any(), request2.StoreID.Int64).Times(0)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "StoreNotFound",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(errand, nil)
				store.EXPECT().
					GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).
					Times(1).
					Return([]db.Request{request1, request2}, nil)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request1.ID).Times(1).Return(request1Items, nil)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request2.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(1).Return(buyer1, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetStore(gomock.Any(), request1.StoreID.Int64).Times(1).Return(db.Store{}, sql.ErrNoRows)
				store.EXPECT().GetStore(gomock.Any(), request2.StoreID.Int64).Times(0)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "FailToGetStore",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(errand, nil)
				store.EXPECT().
					GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).
					Times(1).
					Return([]db.Request{request1, request2}, nil)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request1.ID).Times(1).Return(request1Items, nil)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request2.ID).Times(0)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(1).Return(buyer1, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(0)
				store.EXPECT().GetStore(gomock.Any(), request1.StoreID.Int64).Times(1).Return(db.Store{}, sql.ErrConnDone)
				store.EXPECT().GetStore(gomock.Any(), request2.StoreID.Int64).Times(0)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "CommunityNotFound",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(errand, nil)
				store.EXPECT().
					GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).
					Times(1).
					Return([]db.Request{request1, request2}, nil)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request1.ID).Times(1).Return(request1Items, nil)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request2.ID).Times(1).Return(request2Items, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(1).Return(buyer1, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(1).Return(buyer2, nil)
				store.EXPECT().GetStore(gomock.Any(), request1.StoreID.Int64).Times(1).Return(groceryStore, nil)
				store.EXPECT().GetStore(gomock.Any(), request2.StoreID.Int64).Times(0)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(1).Return(db.Community{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "FailToGetCommunity",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, shopper.ID, shopper.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().GetActiveErrand(gomock.Any(), shopper.ID).Times(1).Return(errand, nil)
				store.EXPECT().
					GetRequestsByErrandId(gomock.Any(), sql.NullInt64{Int64: errand.ID, Valid: true}).
					Times(1).
					Return([]db.Request{request1, request2}, nil)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request1.ID).Times(1).Return(request1Items, nil)
				store.EXPECT().GetItemsByRequest(gomock.Any(), request2.ID).Times(1).Return(request2Items, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer1.ID).Times(1).Return(buyer1, nil)
				store.EXPECT().GetUser(gomock.Any(), buyer2.ID).Times(1).Return(buyer2, nil)
				store.EXPECT().GetStore(gomock.Any(), request1.StoreID.Int64).Times(1).Return(groceryStore, nil)
				store.EXPECT().GetStore(gomock.Any(), request2.StoreID.Int64).Times(0)
				store.EXPECT().GetCommunity(gomock.Any(), community.ID).Times(1).Return(db.Community{}, sql.ErrConnDone)
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
			url := "/errand/active"
			request, err := http.NewRequest(http.MethodGet, url, nil)
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
