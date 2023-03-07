package api

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_db "github.com/git-adithyanair/cs130-group-project/db/mock"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	"github.com/git-adithyanair/cs130-group-project/token"
	"github.com/git-adithyanair/cs130-group-project/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetUserRequest(t *testing.T) {

	user, _ := createRandomUser(t)
	in_progress, complete, pending := []db.Request{}, []db.Request{}, []db.Request{}
	groceryStore := createRandomStore(t)
	items := make(map[int64]([]db.Item))
	requests_no_items := []db.Request{createRandomRequest(t, user.ID, util.RandomID(), groceryStore.ID)}
	requests_invalid_store := []db.Request{createRandomRequest(t, user.ID, util.RandomID(), groceryStore.ID+100)}
	requests_invalid_store_items := []db.Item{createRandomItem(t, user.ID, requests_invalid_store[0].ID, db.ItemQuantityTypeGal, 0.0, false, false, false)}

	for i := 0; i < 10; i++ {
		request := createRandomRequestWithRandomStatus(t, user.ID, groceryStore.ID)
		if request.Status == db.RequestStatusInProgress {
			in_progress = append(in_progress, request)
		} else if request.Status == db.RequestStatusPending {
			pending = append(pending, request)
		} else if request.Status == db.RequestStatusCompleted {
			complete = append(complete, request)
		}
		items[request.ID] = []db.Item{
			createRandomItem(t, user.ID, request.ID, db.ItemQuantityTypeGal, 0.0, false, false, false),
			createRandomItem(t, user.ID, request.ID, db.ItemQuantityTypeGal, 0.0, false, false, false),
		}
	}

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
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusPending,
					}).
					Times(1).
					Return(pending, nil)
				for _, request := range pending {
					store.EXPECT().
						GetItemsByRequest(gomock.Any(), request.ID).
						Times(1).
						Return(items[request.ID], nil)
					store.EXPECT().
						GetStore(gomock.Any(), groceryStore.ID).
						Times(1).
						Return(groceryStore, nil)
				}
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusInProgress,
					}).
					Times(1).
					Return(in_progress, nil)
				for _, request := range in_progress {
					store.EXPECT().
						GetItemsByRequest(gomock.Any(), request.ID).
						Times(1).
						Return(items[request.ID], nil)
					store.EXPECT().
						GetStore(gomock.Any(), groceryStore.ID).
						Times(1).
						Return(groceryStore, nil)
				}
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusCompleted,
					}).
					Times(1).
					Return(complete, nil)
				for _, request := range complete {
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
				requiredBodyMatchUserRequestsResponse(t, recorder.Body, pending, in_progress, complete)
			},
		},
		{
			name: "OKEmpty",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusPending,
					}).
					Times(1).
					Return([]db.Request{}, nil)
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusInProgress,
					}).
					Times(1).
					Return([]db.Request{}, nil)
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusCompleted,
					}).
					Times(1).
					Return([]db.Request{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requiredBodyMatchUserRequestsResponse(t, recorder.Body, []db.Request{}, []db.Request{}, []db.Request{})
			},
		},
		{
			name: "InternalServerErrorPending",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusPending,
					}).
					Times(1).
					Return([]db.Request{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InternalServerErrorInProgress",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusPending,
					}).
					Times(1).
					Return(pending, nil)
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusInProgress,
					}).
					Times(1).
					Return([]db.Request{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InternalServerErrorComplete",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusPending,
					}).
					Times(1).
					Return(pending, nil)
				for _, request := range pending {
					store.EXPECT().
						GetItemsByRequest(gomock.Any(), request.ID).
						Times(1).
						Return(items[request.ID], nil)
					store.EXPECT().
						GetStore(gomock.Any(), groceryStore.ID).
						Times(1).
						Return(groceryStore, nil)
				}
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusInProgress,
					}).
					Times(1).
					Return(in_progress, nil)
				for _, request := range in_progress {
					store.EXPECT().
						GetItemsByRequest(gomock.Any(), request.ID).
						Times(1).
						Return(items[request.ID], nil)
					store.EXPECT().
						GetStore(gomock.Any(), groceryStore.ID).
						Times(1).
						Return(groceryStore, nil)
				}
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusCompleted,
					}).
					Times(1).
					Return([]db.Request{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "StatusNotFoundItem",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusPending,
					}).
					Times(1).
					Return(requests_no_items, nil)
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
			name: "InternalServerErrorItem",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusPending,
					}).
					Times(1).
					Return(requests_no_items, nil)
				store.EXPECT().
					GetItemsByRequest(gomock.Any(), requests_no_items[0].ID).
					Times(1).
					Return([]db.Item{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "StatusNotFoundStore",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusPending,
					}).
					Times(1).
					Return(requests_invalid_store, nil)
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
			name: "InternalServerErrorStore",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusPending,
					}).
					Times(1).
					Return(requests_invalid_store, nil)
				store.EXPECT().
					GetItemsByRequest(gomock.Any(), requests_invalid_store[0].ID).
					Times(1).
					Return(requests_invalid_store_items, nil)
				store.EXPECT().
					GetStore(gomock.Any(), groceryStore.ID+100).
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
			url := "/user/requests"
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
