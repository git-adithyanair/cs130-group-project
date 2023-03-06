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

func createRandomRequestWithRandomStatus(t *testing.T, userID int64) db.Request {
	requestStatus := []db.RequestStatus{
		db.RequestStatusCompleted,
		db.RequestStatusInProgress,
		db.RequestStatusInProgress,
	}
	return db.Request{
		ID:          util.RandomID(),
		UserID:      userID,
		CommunityID: sql.NullInt64{Int64: util.RandomID(), Valid: true},
		StoreID:     sql.NullInt64{Int64: util.RandomID(), Valid: true},
		Status:      requestStatus[util.RandomInt(0, len(requestStatus)-1)],
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

func createRandomItemWithUser(t *testing.T, requestedBy int64) db.Item {
	preferredBrand := sql.NullString{String: util.RandomString(6), Valid: true}
	image := sql.NullString{String: "", Valid: false}
	extraNotes := sql.NullString{String: util.RandomString(20), Valid: true}
	quantityTypes := []db.ItemQuantityType{
		db.ItemQuantityTypeNumerical,
		db.ItemQuantityTypeOz,
		db.ItemQuantityTypeLbs,
		db.ItemQuantityTypeFlOz,
		db.ItemQuantityTypeGal,
		db.ItemQuantityTypeLitres,
	}

	return db.Item{
		ID:             util.RandomID(),
		Name:           util.RandomItemName(),
		RequestedBy:    requestedBy,
		RequestID:      util.RandomID(),
		QuantityType:   quantityTypes[util.RandomInt(0, len(quantityTypes)-1)],
		Quantity:       util.RandomFloat(0, 10),
		PreferredBrand: preferredBrand,
		Image:          image,
		Found:          sql.NullBool{Bool: false, Valid: false},
		ExtraNotes:     extraNotes,
	}
}

func createRandomMember(t *testing.T, userID int64, communityID int64) db.Member {
	return db.Member{
		UserID:      userID,
		CommunityID: communityID,
	}
}

func createRandomErrand(t *testing.T, userID int64, communityID int64) db.Errand {
	return db.Errand{
		ID:          util.RandomID(),
		UserID:      userID,
		CommunityID: communityID,
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

func requireBodyMatchItemWithStatus(t *testing.T, body *bytes.Buffer, item db.Item, found bool) {

	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var itemResponse db.Item
	err = json.Unmarshal(data, &itemResponse)
	require.NoError(t, err)

	require.Equal(t, item.ID, itemResponse.ID)
	require.Equal(t, item.Name, itemResponse.Name)
	require.Equal(t, item.RequestedBy, itemResponse.RequestedBy)
	require.Equal(t, item.RequestID, itemResponse.RequestID)
	require.Equal(t, item.QuantityType, itemResponse.QuantityType)
	require.Equal(t, item.Quantity, itemResponse.Quantity)
	require.Equal(t, item.PreferredBrand, itemResponse.PreferredBrand)
	require.Equal(t, item.Image, itemResponse.Image)
	require.Equal(t, item.ExtraNotes, itemResponse.ExtraNotes)
	require.Equal(t, sql.NullBool{Bool: found, Valid: true}, itemResponse.Found)

}

func requiredBodyMatchUserRequestsResponse(t *testing.T, body *bytes.Buffer, pending []db.Request, inProgress []db.Request, complete []db.Request) {

	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var userRequestsRespone userRequestsResponse
	err = json.Unmarshal(data, &userRequestsRespone)
	require.NoError(t, err)

	require.Equal(t, len(pending), len(userRequestsRespone.Pending))
	require.Equal(t, len(inProgress), len(userRequestsRespone.InProgress))
	require.Equal(t, len(complete), len(userRequestsRespone.Complete))
}
