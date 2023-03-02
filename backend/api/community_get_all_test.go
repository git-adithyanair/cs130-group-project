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

func TestGetAllCommmunities(t *testing.T) {

	user, _ := createRandomUser(t)

	communities := []db.Community{}
	for i := 0; i < 10; i++ {
		communities = append(communities, createRandomCommunity(t, user.ID))
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
					ListCommunities(gomock.Any(), db.ListCommunitiesParams{
						Limit:  100,
						Offset: 0,
					}).
					Times(1).
					Return(communities, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				data, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)
				var getCommunitiesResponse []db.Community
				err = json.Unmarshal(data, &getCommunitiesResponse)
				require.NoError(t, err)
				require.Equal(t, len(communities), len(getCommunitiesResponse))
			},
		},
		{
			name: "InternalServerError",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.ID, user.Email)
			},
			buildStubs: func(store *mock_db.MockDBStore) {
				store.EXPECT().
					ListCommunities(gomock.Any(), db.ListCommunitiesParams{
						Limit:  100,
						Offset: 0,
					}).
					Times(1).
					Return([]db.Community{}, sql.ErrConnDone)
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
