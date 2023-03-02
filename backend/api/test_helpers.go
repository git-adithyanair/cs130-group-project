package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	"github.com/git-adithyanair/cs130-group-project/token"
	"github.com/git-adithyanair/cs130-group-project/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) (db.User, string) {
	rawPassword := util.RandomString(10)
	hashedPassword, err := util.HashPassword(rawPassword)
	require.NoError(t, err)
	return db.User{
		Email:          util.RandomEmail(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomFullName(),
		PhoneNumber:    util.RandomPhoneNumber(),
		PlaceID:        util.RandomPlaceID(),
		XCoord:         util.RandomCoordinate(),
		YCoord:         util.RandomCoordinate(),
		Address:        util.RandomAddress(),
	}, rawPassword
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {

	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var registerResponse authUserResponse
	err = json.Unmarshal(data, &registerResponse)
	require.NoError(t, err)

	require.Equal(t, user.ID, registerResponse.User.ID)
	require.Equal(t, user.Email, registerResponse.User.Email)
	require.Equal(t, user.FullName, registerResponse.User.FullName)
	require.WithinDuration(t, user.CreatedAt, registerResponse.User.CreatedAt, time.Second)

}

func createRandomCommunity(t *testing.T, adminID int64) db.Community {
	return db.Community{
		Name:         util.RandomCommunityName(),
		Admin:        adminID,
		PlaceID:      util.RandomPlaceID(),
		CenterXCoord: util.RandomCoordinate(),
		CenterYCoord: util.RandomCoordinate(),
		Address:      util.RandomAddress(),
		Range:        util.RandomRange(),
	}
}

func addAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	userID int64,
	email string,
) {
	accessToken, err := tokenMaker.CreateToken(userID, email)
	require.NoError(t, err)
	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, accessToken)
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}

func requireBodyMatchCommunity(t *testing.T, body *bytes.Buffer, community db.Community) {

	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var communityResponse db.Community
	err = json.Unmarshal(data, &communityResponse)
	require.NoError(t, err)

	require.Equal(t, community.ID, communityResponse.ID)
	require.Equal(t, community.Name, communityResponse.Name)
	require.Equal(t, community.Admin, communityResponse.Admin)
	require.Equal(t, community.PlaceID, communityResponse.PlaceID)
	require.Equal(t, community.CenterXCoord, communityResponse.CenterXCoord)
	require.Equal(t, community.CenterYCoord, communityResponse.CenterYCoord)
	require.Equal(t, community.Address, communityResponse.Address)
	require.Equal(t, community.Range, communityResponse.Range)
	require.WithinDuration(t, community.CreatedAt, communityResponse.CreatedAt, time.Second)

}
