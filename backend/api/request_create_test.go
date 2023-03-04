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

func TestCreateRequest(t *testing.T) {

	user, _ := createRandomUser(t)
	community := createRandomCommunity(t, user.ID)
	groceryStore := createRandomStore(t)
	createRandomCommunityStore(t, community.ID, groceryStore.ID)
	request := createRandomRequest(t, user.ID, community.ID, groceryStore.ID)

	validItems := []db.Item{
		createRandomItem(t, user.ID, request.ID, "oz", 1, true, true, true),
		createRandomItem(t, user.ID, request.ID, "fl_oz", 1, true, true, true),
	}
	validItemsBody := []gin.H{
		{
			"name":            validItems[0].Name,
			"quantity_type":   validItems[0].QuantityType,
			"quantity":        validItems[0].Quantity,
			"preferred_brand": validItems[0].PreferredBrand.String,
			"image":           validItems[0].Image.String,
			"extra_notes":     validItems[0].ExtraNotes.String,
		},
		{
			"name":            validItems[1].Name,
			"quantity_type":   validItems[1].QuantityType,
			"quantity":        validItems[1].Quantity,
			"preferred_brand": validItems[1].PreferredBrand.String,
			"image":           validItems[1].Image.String,
			"extra_notes":     validItems[1].ExtraNotes.String,
		},
	}
	validItemsWithMissingFields := []db.Item{
		createRandomItem(t, user.ID, request.ID, "oz", 1, true, false, true),
		createRandomItem(t, user.ID, request.ID, "fl_oz", 1, false, true, false),
	}
	validItemsBodyWithMissingFields := []gin.H{
		{
			"name":            validItemsWithMissingFields[0].Name,
			"quantity_type":   validItemsWithMissingFields[0].QuantityType,
			"quantity":        validItemsWithMissingFields[0].Quantity,
			"preferred_brand": validItemsWithMissingFields[0].PreferredBrand.String,
			"extra_notes":     validItemsWithMissingFields[0].ExtraNotes.String,
		},
		{
			"name":          validItemsWithMissingFields[1].Name,
			"quantity_type": validItemsWithMissingFields[1].QuantityType,
			"quantity":      validItemsWithMissingFields[1].Quantity,
			"image":         validItemsWithMissingFields[1].Image.String,
		},
	}

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
				"community_id": community.ID,
				"store_id":     groceryStore.ID,
				"items":        validItemsBody,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					GetStore(gomock.Any(), groceryStore.ID).
					Times(1).
					Return(groceryStore, nil)
				store.EXPECT().
					CreateRequest(gomock.Any(), db.CreateRequestParams{
						UserID:      user.ID,
						CommunityID: sql.NullInt64{Int64: community.ID, Valid: true},
						StoreID:     sql.NullInt64{Int64: groceryStore.ID, Valid: true},
					}).Times(1).Return(request, nil)
				for _, item := range validItems {
					store.EXPECT().
						CreateItem(gomock.Any(), db.CreateItemParams{
							Name:           item.Name,
							RequestedBy:    user.ID,
							RequestID:      request.ID,
							QuantityType:   item.QuantityType,
							Quantity:       item.Quantity,
							PreferredBrand: item.PreferredBrand,
							Image:          item.Image,
							ExtraNotes:     item.ExtraNotes,
						}).Times(1).Return(item, nil)
				}
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			body: gin.H{
				"store_id": groceryStore.ID,
				"items":    validItemsBody,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(0)
				store.EXPECT().
					GetStore(gomock.Any(), groceryStore.ID).
					Times(0)
				store.EXPECT().
					CreateRequest(gomock.Any(), db.CreateRequestParams{
						UserID:      user.ID,
						CommunityID: sql.NullInt64{Int64: community.ID, Valid: true},
						StoreID:     sql.NullInt64{Int64: groceryStore.ID, Valid: true},
					}).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "CommunityNotFound",
			body: gin.H{
				"community_id": community.ID + 100,
				"store_id":     groceryStore.ID,
				"items":        validItemsBody,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID+100).
					Times(1).
					Return(db.Community{}, sql.ErrNoRows)
				store.EXPECT().
					GetStore(gomock.Any(), groceryStore.ID).
					Times(0)
				store.EXPECT().
					CreateRequest(gomock.Any(), db.CreateRequestParams{
						UserID:      user.ID,
						CommunityID: sql.NullInt64{Int64: community.ID, Valid: true},
						StoreID:     sql.NullInt64{Int64: groceryStore.ID, Valid: true},
					}).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "CommunityQueryFail",
			body: gin.H{
				"community_id": community.ID,
				"store_id":     groceryStore.ID,
				"items":        validItemsBody,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(1).
					Return(db.Community{}, sql.ErrConnDone)
				store.EXPECT().
					GetStore(gomock.Any(), groceryStore.ID).
					Times(0)
				store.EXPECT().
					CreateRequest(gomock.Any(), db.CreateRequestParams{
						UserID:      user.ID,
						CommunityID: sql.NullInt64{Int64: community.ID, Valid: true},
						StoreID:     sql.NullInt64{Int64: groceryStore.ID, Valid: true},
					}).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "StoreNotFound",
			body: gin.H{
				"community_id": community.ID,
				"store_id":     groceryStore.ID + 100,
				"items":        validItemsBody,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					GetStore(gomock.Any(), groceryStore.ID+100).
					Times(1).
					Return(db.Store{}, sql.ErrNoRows)
				store.EXPECT().
					CreateRequest(gomock.Any(), db.CreateRequestParams{
						UserID:      user.ID,
						CommunityID: sql.NullInt64{Int64: community.ID, Valid: true},
						StoreID:     sql.NullInt64{Int64: groceryStore.ID, Valid: true},
					}).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "StoreQueryFail",
			body: gin.H{
				"community_id": community.ID,
				"store_id":     groceryStore.ID,
				"items":        validItemsBody,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					GetStore(gomock.Any(), groceryStore.ID).
					Times(1).
					Return(db.Store{}, sql.ErrConnDone)
				store.EXPECT().
					CreateRequest(gomock.Any(), db.CreateRequestParams{
						UserID:      user.ID,
						CommunityID: sql.NullInt64{Int64: community.ID, Valid: true},
						StoreID:     sql.NullInt64{Int64: groceryStore.ID, Valid: true},
					}).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "RequestCreateFail",
			body: gin.H{
				"community_id": community.ID,
				"store_id":     groceryStore.ID,
				"items":        validItemsBody,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					GetStore(gomock.Any(), groceryStore.ID).
					Times(1).
					Return(groceryStore, nil)
				store.EXPECT().
					CreateRequest(gomock.Any(), db.CreateRequestParams{
						UserID:      user.ID,
						CommunityID: sql.NullInt64{Int64: community.ID, Valid: true},
						StoreID:     sql.NullInt64{Int64: groceryStore.ID, Valid: true},
					}).Times(1).Return(db.Request{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "NoItems",
			body: gin.H{
				"community_id": community.ID,
				"store_id":     groceryStore.ID,
				"items":        []gin.H{},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(1)
				store.EXPECT().
					GetStore(gomock.Any(), groceryStore.ID).
					Times(0)
				store.EXPECT().
					CreateRequest(gomock.Any(), db.CreateRequestParams{
						UserID:      user.ID,
						CommunityID: sql.NullInt64{Int64: community.ID, Valid: true},
						StoreID:     sql.NullInt64{Int64: groceryStore.ID, Valid: true},
					}).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "ValidItemsWithMissingFields",
			body: gin.H{
				"community_id": community.ID,
				"store_id":     groceryStore.ID,
				"items":        validItemsBodyWithMissingFields,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					GetStore(gomock.Any(), groceryStore.ID).
					Times(1).
					Return(groceryStore, nil)
				store.EXPECT().
					CreateRequest(gomock.Any(), db.CreateRequestParams{
						UserID:      user.ID,
						CommunityID: sql.NullInt64{Int64: community.ID, Valid: true},
						StoreID:     sql.NullInt64{Int64: groceryStore.ID, Valid: true},
					}).Times(1).
					Return(request, nil)
				for _, item := range validItemsWithMissingFields {
					store.EXPECT().
						CreateItem(gomock.Any(), db.CreateItemParams{
							Name:           item.Name,
							RequestedBy:    user.ID,
							RequestID:      request.ID,
							QuantityType:   item.QuantityType,
							Quantity:       item.Quantity,
							PreferredBrand: item.PreferredBrand,
							Image:          item.Image,
							ExtraNotes:     item.ExtraNotes,
						}).Times(1).Return(item, nil)
				}
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
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
			url := "/request"
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
