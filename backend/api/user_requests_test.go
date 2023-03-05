package api

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_db "github.com/git-adithyanair/cs130-group-project/db/mock"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	"github.com/git-adithyanair/cs130-group-project/token"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetUserRequest(t *testing.T) {

	user, _ := createRandomUser(t)
	in_progress, complete, pending := []db.Request{}, []db.Request{}, []db.Request{}

	for i := 0; i < 10; i++ {
		request := createRandomRequestWithRandomStatus(t, user.ID)
		if request.Status == db.RequestStatusInProgress {
			in_progress = append(in_progress, request)
		} else if request.Status == db.RequestStatusPending {
			pending = append(pending, request)
		} else if request.Status == db.RequestStatusCompleted {
			complete = append(complete, request)
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
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusInProgress,
					}).
					Times(1).
					Return(in_progress, nil)
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusCompleted,
					}).
					Times(1).
					Return(complete, nil)
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
				store.EXPECT().
					GetRequestsForUserByStatus(gomock.Any(), db.GetRequestsForUserByStatusParams{
						UserID: user.ID,
						Status: db.RequestStatusInProgress,
					}).
					Times(1).
					Return(in_progress, nil)
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
