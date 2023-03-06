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

func TestCreateCommunity(t *testing.T) {

	user, _ := createRandomUser(t)
	community := createRandomCommunity(t, user.ID)
	member := createRandomMember(t, user.ID, community.ID)

	validStores := []db.Store{createRandomStore(t), createRandomStore(t)}
	validStoreBody := []gin.H{
		{
			"name":     validStores[0].Name,
			"x_coord":  validStores[0].XCoord,
			"y_coord":  validStores[0].YCoord,
			"place_id": validStores[0].PlaceID,
			"address":  validStores[0].Address,
		},
		{
			"name":     validStores[1].Name,
			"x_coord":  validStores[1].XCoord,
			"y_coord":  validStores[1].YCoord,
			"place_id": validStores[1].PlaceID,
			"address":  validStores[1].Address,
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
			name: "OKNewStores",
			body: gin.H{
				"name":           community.Name,
				"center_x_coord": community.CenterXCoord,
				"center_y_coord": community.CenterYCoord,
				"place_id":       community.PlaceID,
				"range":          community.Range,
				"address":        community.Address,
				"stores":         validStoreBody,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunityByPlaceID(gomock.Any(), community.PlaceID).
					Times(1).
					Return(db.Community{}, sql.ErrNoRows)
				store.EXPECT().
					CreateCommunity(gomock.Any(), db.CreateCommunityParams{
						Name:         community.Name,
						Admin:        user.ID,
						CenterXCoord: community.CenterXCoord,
						CenterYCoord: community.CenterYCoord,
						Range:        community.Range,
						PlaceID:      community.PlaceID,
						Address:      community.Address,
					}).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					CreateMember(gomock.Any(), db.CreateMemberParams{
						CommunityID: community.ID,
						UserID:      user.ID,
					}).
					Times(1).
					Return(member, nil)
				for _, groceryStore := range validStores {
					store.EXPECT().
						GetStoreByPlaceId(gomock.Any(), groceryStore.PlaceID).
						Times(1).
						Return(db.Store{}, sql.ErrNoRows)
					store.EXPECT().
						CreateStore(gomock.Any(), db.CreateStoreParams{
							Name:    groceryStore.Name,
							XCoord:  groceryStore.XCoord,
							YCoord:  groceryStore.YCoord,
							PlaceID: groceryStore.PlaceID,
							Address: groceryStore.Address,
						}).
						Times(1).
						Return(groceryStore, nil)
					store.EXPECT().
						CreateCommunityStore(gomock.Any(), db.CreateCommunityStoreParams{
							CommunityID: community.ID,
							StoreID:     groceryStore.ID,
						})
				}
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "OKOldStores",
			body: gin.H{
				"name":           community.Name,
				"center_x_coord": community.CenterXCoord,
				"center_y_coord": community.CenterYCoord,
				"place_id":       community.PlaceID,
				"range":          community.Range,
				"address":        community.Address,
				"stores":         validStoreBody,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunityByPlaceID(gomock.Any(), community.PlaceID).
					Times(1).
					Return(db.Community{}, sql.ErrNoRows)
				store.EXPECT().
					CreateCommunity(gomock.Any(), db.CreateCommunityParams{
						Name:         community.Name,
						Admin:        user.ID,
						CenterXCoord: community.CenterXCoord,
						CenterYCoord: community.CenterYCoord,
						Range:        community.Range,
						PlaceID:      community.PlaceID,
						Address:      community.Address,
					}).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					CreateMember(gomock.Any(), db.CreateMemberParams{
						CommunityID: community.ID,
						UserID:      user.ID,
					}).
					Times(1).
					Return(member, nil)
				for _, groceryStore := range validStores {
					store.EXPECT().
						GetStoreByPlaceId(gomock.Any(), groceryStore.PlaceID).
						Times(1).
						Return(groceryStore, nil)
					store.EXPECT().
						CreateCommunityStore(gomock.Any(), db.CreateCommunityStoreParams{
							CommunityID: community.ID,
							StoreID:     groceryStore.ID,
						})
				}
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
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
			name: "CreateCommunityInternalServiceError",
			body: gin.H{
				"name":           community.Name,
				"center_x_coord": community.CenterXCoord,
				"center_y_coord": community.CenterYCoord,
				"place_id":       community.PlaceID,
				"range":          community.Range,
				"address":        community.Address,
				"stores":         validStoreBody,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunityByPlaceID(gomock.Any(), community.PlaceID).
					Times(1).
					Return(db.Community{}, sql.ErrNoRows)
				store.EXPECT().
					CreateCommunity(gomock.Any(), db.CreateCommunityParams{
						Name:         community.Name,
						Admin:        user.ID,
						CenterXCoord: community.CenterXCoord,
						CenterYCoord: community.CenterYCoord,
						Range:        community.Range,
						PlaceID:      community.PlaceID,
						Address:      community.Address,
					}).
					Times(1).
					Return(db.Community{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "CreateMemberInternalServerError",
			body: gin.H{
				"name":           community.Name,
				"center_x_coord": community.CenterXCoord,
				"center_y_coord": community.CenterYCoord,
				"place_id":       community.PlaceID,
				"range":          community.Range,
				"address":        community.Address,
				"stores":         []db.Store{},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunityByPlaceID(gomock.Any(), community.PlaceID).
					Times(1).
					Return(db.Community{}, sql.ErrNoRows)
				store.EXPECT().
					CreateCommunity(gomock.Any(), db.CreateCommunityParams{
						Name:         community.Name,
						Admin:        user.ID,
						CenterXCoord: community.CenterXCoord,
						CenterYCoord: community.CenterYCoord,
						Range:        community.Range,
						PlaceID:      community.PlaceID,
						Address:      community.Address,
					}).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					CreateMember(gomock.Any(), db.CreateMemberParams{
						CommunityID: community.ID,
						UserID:      user.ID,
					}).
					Times(1).
					Return(db.Member{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "BadRequestNoStores",
			body: gin.H{
				"name":           community.Name,
				"center_x_coord": community.CenterXCoord,
				"center_y_coord": community.CenterYCoord,
				"place_id":       community.PlaceID,
				"range":          community.Range,
				"address":        community.Address,
				"stores":         []db.Store{},
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunityByPlaceID(gomock.Any(), community.PlaceID).
					Times(1).
					Return(db.Community{}, sql.ErrNoRows)
				store.EXPECT().
					CreateCommunity(gomock.Any(), db.CreateCommunityParams{
						Name:         community.Name,
						Admin:        user.ID,
						CenterXCoord: community.CenterXCoord,
						CenterYCoord: community.CenterYCoord,
						Range:        community.Range,
						PlaceID:      community.PlaceID,
						Address:      community.Address,
					}).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					CreateMember(gomock.Any(), db.CreateMemberParams{
						CommunityID: community.ID,
						UserID:      user.ID,
					}).
					Times(1).
					Return(member, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "GetCommunityByPlaceIDInternalServerError",
			body: gin.H{
				"name":           community.Name,
				"center_x_coord": community.CenterXCoord,
				"center_y_coord": community.CenterYCoord,
				"place_id":       community.PlaceID,
				"range":          community.Range,
				"address":        community.Address,
				"stores":         validStoreBody,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunityByPlaceID(gomock.Any(), community.PlaceID).
					Times(1).
					Return(db.Community{}, sql.ErrConnDone)
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
			url := "/community"
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
