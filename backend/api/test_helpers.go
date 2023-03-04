package api

import (
	"bytes"
	"database/sql"
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
		ID:             util.RandomID(),
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
		ID:           util.RandomID(),
		Name:         util.RandomCommunityName(),
		Admin:        adminID,
		PlaceID:      util.RandomPlaceID(),
		CenterXCoord: util.RandomCoordinate(),
		CenterYCoord: util.RandomCoordinate(),
		Address:      util.RandomAddress(),
		Range:        util.RandomRange(),
	}
}

func createRandomStore(t *testing.T) db.Store {
	return db.Store{
		ID:      util.RandomID(),
		Name:    util.RandomStoreName(),
		PlaceID: util.RandomPlaceID(),
		XCoord:  util.RandomCoordinate(),
		YCoord:  util.RandomCoordinate(),
		Address: util.RandomAddress(),
	}
}

func createRandomCommunityStore(t *testing.T, communityID int64, storeID int64) db.CommunityStore {
	return db.CommunityStore{
		CommunityID: communityID,
		StoreID:     storeID,
	}
}

func createRandomRequest(t *testing.T, userID int64, communityID int64, storeID int64) db.Request {
	return db.Request{
		ID:          util.RandomID(),
		UserID:      userID,
		CommunityID: sql.NullInt64{Int64: communityID, Valid: true},
		StoreID:     sql.NullInt64{Int64: storeID, Valid: true},
	}
}

func createRandomItem(
	t *testing.T,
	requestedBy int64,
	requestID int64,
	quantityType db.ItemQuantityType,
	quantity float64,
	withPreferredBrand bool,
	withImage bool,
	withExtraNotes bool,
) db.Item {
	preferredBrand := sql.NullString{String: "", Valid: false}
	if withPreferredBrand {
		preferredBrand.Valid = true
		preferredBrand.String = util.RandomString(6)
	}
	image := sql.NullString{String: "", Valid: false}
	if withImage {
		image.Valid = true
		image.String = util.RandomString(300)
	}
	extraNotes := sql.NullString{String: "", Valid: false}
	if withExtraNotes {
		extraNotes.Valid = true
		extraNotes.String = util.RandomString(20)
	}
	return db.Item{
		ID:             util.RandomID(),
		Name:           util.RandomItemName(),
		RequestedBy:    requestedBy,
		RequestID:      requestID,
		QuantityType:   quantityType,
		Quantity:       quantity,
		PreferredBrand: preferredBrand,
		Image:          image,
		Found:          sql.NullBool{Bool: false, Valid: false},
		ExtraNotes:     extraNotes,
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
