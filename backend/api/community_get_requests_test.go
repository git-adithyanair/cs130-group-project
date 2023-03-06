package api

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	mock_db "github.com/git-adithyanair/cs130-group-project/db/mock"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	"github.com/git-adithyanair/cs130-group-project/token"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetRequestsByCommunity(t *testing.T) {

	user, _ := createRandomUser(t)
	community := createRandomCommunity(t, user.ID)
	groceryStore := createRandomStore(t)
	member := createRandomMember(t, user.ID, community.ID)
	requests := []db.Request{}
	items := make(map[int64]([]db.Item))

	for i := 0; i < 10; i++ {
		request := createRandomRequest(t, user.ID, community.ID, groceryStore.ID)
		requests = append(requests, request)
		items[request.ID] = []db.Item{
			createRandomItem(t, user.ID, request.ID, db.ItemQuantityTypeGal, 0.0, false, false, false),
			createRandomItem(t, user.ID, request.ID, db.ItemQuantityTypeGal, 0.0, false, false, false),
		}
	}

	requests_invalid_user := []db.Request{createRandomRequest(t, user.ID+100, community.ID, groceryStore.ID)}
	requests_no_items := []db.Request{createRandomRequest(t, user.ID, community.ID, groceryStore.ID)}
	requests_invalid_store := []db.Request{createRandomRequest(t, user.ID, community.ID, groceryStore.ID+100)}
	requests_invalid_store_items := []db.Item{createRandomItem(t, user.ID, requests_invalid_store[0].ID, db.ItemQuantityTypeGal, 0.0, false, false, false)}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_db.NewMockDBStore(ctrl)

	testCases := []struct {
		name          string
		communityID   int64
		limit         int64
		offset        int64
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mock_db.MockDBStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:        "OK",
			communityID: community.ID,
			limit:       10,
			offset:      0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					GetMember(gomock.Any(), db.GetMemberParams{
						UserID:      user.ID,
						CommunityID: community.ID,
					}).
					Times(1).
					Return(member, nil)
				store.EXPECT().
					GetRequestsByCommunityId(gomock.Any(), db.GetRequestsByCommunityIdParams{
						CommunityID: sql.NullInt64{Int64: community.ID, Valid: true},
						Limit:       10,
						Offset:      0,
					}).
					Times(1).
					Return(requests, nil)
				for _, request := range requests {
					store.EXPECT().
						GetUser(gomock.Any(), request.UserID).
						Times(1).
						Return(user, nil)
					store.EXPECT().
						GetItemsByRequest(gomock.Any(), request.ID).
						Times(1).
						Return(items[request.ID], nil)
					store.EXPECT().
						GetStore(gomock.Any(), groceryStore.ID).
						Times(1).
						Return(groceryStore, nil)
				}
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:        "BadRequest",
			communityID: -1,
			limit:       10,
			offset:      0,
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
			name:        "StatusNotFoundCommunity",
			communityID: community.ID + 100,
			limit:       10,
			offset:      0,
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
			name:        "InternalServerErrorCommunity",
			communityID: community.ID,
			limit:       10,
			offset:      0,
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
			name:        "StatusNotFoundMember",
			communityID: community.ID,
			limit:       10,
			offset:      0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID+100, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					GetMember(gomock.Any(), db.GetMemberParams{
						UserID:      user.ID + 100,
						CommunityID: community.ID,
					}).
					Times(1).
					Return(db.Member{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:        "InternalServerErrorMember",
			communityID: community.ID,
			limit:       10,
			offset:      0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					GetMember(gomock.Any(), db.GetMemberParams{
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
		{
			name:        "InternalServerErrorGetRequestsByCommunityId",
			communityID: community.ID,
			limit:       10,
			offset:      0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					GetMember(gomock.Any(), db.GetMemberParams{
						UserID:      user.ID,
						CommunityID: community.ID,
					}).
					Times(1).
					Return(member, nil)
				store.EXPECT().
					GetRequestsByCommunityId(gomock.Any(), db.GetRequestsByCommunityIdParams{
						CommunityID: sql.NullInt64{Int64: community.ID, Valid: true},
						Limit:       10,
						Offset:      0,
					}).
					Times(1).
					Return([]db.Request{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:        "StatusNotFoundUser",
			communityID: community.ID,
			limit:       10,
			offset:      0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					GetMember(gomock.Any(), db.GetMemberParams{
						UserID:      user.ID,
						CommunityID: community.ID,
					}).
					Times(1).
					Return(member, nil)
				store.EXPECT().
					GetRequestsByCommunityId(gomock.Any(), db.GetRequestsByCommunityIdParams{
						CommunityID: sql.NullInt64{Int64: community.ID, Valid: true},
						Limit:       10,
						Offset:      0,
					}).
					Times(1).
					Return(requests_invalid_user, nil)
				store.EXPECT().
					GetUser(gomock.Any(), requests[0].UserID+100).
					Times(1).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:        "InternalServerErrorUser",
			communityID: community.ID,
			limit:       10,
			offset:      0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					GetMember(gomock.Any(), db.GetMemberParams{
						UserID:      user.ID,
						CommunityID: community.ID,
					}).
					Times(1).
					Return(member, nil)
				store.EXPECT().
					GetRequestsByCommunityId(gomock.Any(), db.GetRequestsByCommunityIdParams{
						CommunityID: sql.NullInt64{Int64: community.ID, Valid: true},
						Limit:       10,
						Offset:      0,
					}).
					Times(1).
					Return(requests, nil)
				store.EXPECT().
					GetUser(gomock.Any(), requests[0].UserID).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:        "StatusNotFoundGetItemByRequest",
			communityID: community.ID,
			limit:       10,
			offset:      0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					GetMember(gomock.Any(), db.GetMemberParams{
						UserID:      user.ID,
						CommunityID: community.ID,
					}).
					Times(1).
					Return(member, nil)
				store.EXPECT().
					GetRequestsByCommunityId(gomock.Any(), db.GetRequestsByCommunityIdParams{
						CommunityID: sql.NullInt64{Int64: community.ID, Valid: true},
						Limit:       10,
						Offset:      0,
					}).
					Times(1).
					Return(requests_no_items, nil)
				store.EXPECT().
					GetUser(gomock.Any(), requests_no_items[0].UserID).
					Times(1).
					Return(user, nil)
				store.EXPECT().
					GetItemsByRequest(gomock.Any(), requests_no_items[0].ID).
					Times(1).
					Return([]db.Item{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:        "InternalServerErrorGetItemByRequest",
			communityID: community.ID,
			limit:       10,
			offset:      0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					GetMember(gomock.Any(), db.GetMemberParams{
						UserID:      user.ID,
						CommunityID: community.ID,
					}).
					Times(1).
					Return(member, nil)
				store.EXPECT().
					GetRequestsByCommunityId(gomock.Any(), db.GetRequestsByCommunityIdParams{
						CommunityID: sql.NullInt64{Int64: community.ID, Valid: true},
						Limit:       10,
						Offset:      0,
					}).
					Times(1).
					Return(requests, nil)
				store.EXPECT().
					GetUser(gomock.Any(), requests[0].UserID).
					Times(1).
					Return(user, nil)
				store.EXPECT().
					GetItemsByRequest(gomock.Any(), requests[0].ID).
					Times(1).
					Return([]db.Item{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:        "StatusNotFoundStore",
			communityID: community.ID,
			limit:       10,
			offset:      0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					GetMember(gomock.Any(), db.GetMemberParams{
						UserID:      user.ID,
						CommunityID: community.ID,
					}).
					Times(1).
					Return(member, nil)
				store.EXPECT().
					GetRequestsByCommunityId(gomock.Any(), db.GetRequestsByCommunityIdParams{
						CommunityID: sql.NullInt64{Int64: community.ID, Valid: true},
						Limit:       10,
						Offset:      0,
					}).
					Times(1).
					Return(requests_invalid_store, nil)
				store.EXPECT().
					GetUser(gomock.Any(), requests_invalid_store[0].UserID).
					Times(1).
					Return(user, nil)
				store.EXPECT().
					GetItemsByRequest(gomock.Any(), requests_invalid_store[0].ID).
					Times(1).
					Return(requests_invalid_store_items, nil)
				store.EXPECT().
					GetStore(gomock.Any(), groceryStore.ID+100).
					Times(1).
					Return(db.Store{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:        "InternalServerErrorStore",
			communityID: community.ID,
			limit:       10,
			offset:      0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetCommunity(gomock.Any(), community.ID).
					Times(1).
					Return(community, nil)
				store.EXPECT().
					GetMember(gomock.Any(), db.GetMemberParams{
						UserID:      user.ID,
						CommunityID: community.ID,
					}).
					Times(1).
					Return(member, nil)
				store.EXPECT().
					GetRequestsByCommunityId(gomock.Any(), db.GetRequestsByCommunityIdParams{
						CommunityID: sql.NullInt64{Int64: community.ID, Valid: true},
						Limit:       10,
						Offset:      0,
					}).
					Times(1).
					Return(requests, nil)
				store.EXPECT().
					GetUser(gomock.Any(), requests[0].UserID).
					Times(1).
					Return(user, nil)
				store.EXPECT().
					GetItemsByRequest(gomock.Any(), requests[0].ID).
					Times(1).
					Return(items[requests[0].ID], nil)
				store.EXPECT().
					GetStore(gomock.Any(), groceryStore.ID).
					Times(1).
					Return(db.Store{}, sql.ErrConnDone)
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
			url := "/community/requests"
			request, err := http.NewRequest(http.MethodGet, url, nil)
			q := request.URL.Query()
			q.Add("id", strconv.FormatInt(testCase.communityID, 10))
			q.Add("limit", strconv.FormatInt(testCase.limit, 10))
			q.Add("offset", strconv.FormatInt(testCase.offset, 10))
			request.URL.RawQuery = q.Encode()

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
