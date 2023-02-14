package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/git-adithyanair/cs130-group-project/util"
)

const (
	numUsers               = 10
	numCommunities         = 4
	numStores              = 2
	minNumRequestsPerStore = 1
	maxNumRequestsPerStore = 2
	minNumItemsPerRequest  = 1
	maxNumItemsPerRequest  = 10
)

func createRandomUserParam(id string, hashedPassword string) CreateUserParams {
	return CreateUserParams{
		Email:          "user" + id + "@test.com",
		HashedPassword: hashedPassword,
		FullName:       "User " + id,
		PhoneNumber:    util.RandomNumericString(10),
		PlaceID:        util.RandomString(5),
		Address:        id + " User St",
		XCoord:         util.RandomFloat(0, 1000),
		YCoord:         util.RandomFloat(0, 1000),
	}
}

func addRandomUsers(queries *Queries, count int) []User {
	hashedPassword, err := util.HashPassword("password")
	if err != nil {
		log.Println("[DB POPULATE] could not hash password: ", err)
		return []User{}
	}

	users := []User{}

	for i := 0; i < count; i++ {
		createUserParams := createRandomUserParam(fmt.Sprint(i+1), hashedPassword)
		user, err := queries.CreateUser(context.Background(), createUserParams)
		if err != nil {
			log.Println("[DB POPULATE] could not create user: ", err)
			return []User{}
		}
		users = append(users, user)
	}

	return users
}

func createRandomCommunityParam(id string, admin int64) CreateCommunityParams {
	return CreateCommunityParams{
		Name:         "Community " + id,
		Admin:        admin,
		CenterXCoord: util.RandomFloat(0, 1000),
		CenterYCoord: util.RandomFloat(0, 1000),
		PlaceID:      util.RandomString(5),
		Address:      id + " Community St",
		Range:        int32(util.RandomInt(0, 10)),
	}
}

func addRandomCommunities(queries *Queries, count int) []Community {

	communities := []Community{}

	for i := 0; i < count; i++ {
		createCommunityParams := createRandomCommunityParam(fmt.Sprint(i+1), 1)
		community, err := queries.CreateCommunity(context.Background(), createCommunityParams)
		if err != nil {
			log.Println("[DB POPULATE] could not create community: ", err)
			return []Community{}
		}
		communities = append(communities, community)
	}

	return communities
}

func createRandomStoreParam(id string) CreateStoreParams {
	return CreateStoreParams{
		Name:    "Store " + id,
		PlaceID: util.RandomString(5),
		Address: id + " Store St",
		XCoord:  util.RandomFloat(0, 1000),
		YCoord:  util.RandomFloat(0, 1000),
	}
}

func addRandomStores(queries *Queries, count int) []Store {

	stores := []Store{}

	for i := 0; i < count; i++ {
		createStoreParams := createRandomStoreParam(fmt.Sprint(i + 1))
		store, err := queries.CreateStore(context.Background(), createStoreParams)
		if err != nil {
			log.Println("[DB POPULATE] could not create store: ", err)
			return []Store{}
		}
		stores = append(stores, store)
	}

	return stores
}

func createCommunityStoreParam(communityID int64, storeID int64) CreateCommunityStoreParams {
	return CreateCommunityStoreParams{
		CommunityID: communityID,
		StoreID:     storeID,
	}
}

func addRandomCommunityStores(queries *Queries, communitiesCount, storesCount int) []CommunityStore {

	communityStores := []CommunityStore{}

	for i := 0; i < communitiesCount; i++ {
		for j := 0; j < storesCount; j++ {
			createCommunityStoreParams := createCommunityStoreParam(int64(i+1), int64(j+1))
			communityStore, err := queries.CreateCommunityStore(context.Background(), createCommunityStoreParams)
			if err != nil {
				log.Println("[DB POPULATE] could not create community store: ", err)
				return []CommunityStore{}
			}
			communityStores = append(communityStores, communityStore)
		}
	}

	return communityStores
}

func createMemberParam(communityID int64, userID int64) CreateMemberParams {
	return CreateMemberParams{
		CommunityID: communityID,
		UserID:      userID,
	}
}

func addRandomMembers(queries *Queries, usersCount, communitiesCount int) []Member {

	members := []Member{}

	for i := 0; i < usersCount; i++ {
		createMemberParams := createMemberParam(int64((i%communitiesCount)+1), int64(i+1))
		member, err := queries.CreateMember(context.Background(), createMemberParams)
		if err != nil {
			log.Println("[DB POPULATE] could not create member: ", err)
			return []Member{}
		}
		members = append(members, member)
	}

	return members
}

func createRandomRequestParam(userID, storeID, communityID int64) CreateRequestParams {
	return CreateRequestParams{
		UserID:      userID,
		StoreID:     sql.NullInt64{Int64: storeID, Valid: true},
		CommunityID: sql.NullInt64{Int64: communityID, Valid: true},
	}
}

func getUsersInCommunity(communityID int64, members []Member) []int64 {

	userIDs := []int64{}

	for _, member := range members {
		if member.CommunityID == communityID {
			userIDs = append(userIDs, member.UserID)
		}
	}

	return userIDs
}

func addRandomRequests(queries *Queries, communityStores []CommunityStore, members []Member) []Request {

	requests := []Request{}

	for _, communityStore := range communityStores {
		userIDs := getUsersInCommunity(communityStore.CommunityID, members)
		if len(userIDs) == 0 {
			continue
		}
		numRequestsForStore := util.RandomInt(minNumRequestsPerStore, maxNumRequestsPerStore)
		for i := 0; i < numRequestsForStore; i++ {
			createRequestParams := createRandomRequestParam(
				userIDs[util.RandomInt(0, len(userIDs)-1)],
				communityStore.StoreID,
				communityStore.CommunityID,
			)
			request, err := queries.CreateRequest(context.Background(), createRequestParams)
			if err != nil {
				log.Println("[DB POPULATE] could not create request: ", err)
				return []Request{}
			}
			requests = append(requests, request)
		}
	}

	return requests
}

func createRandomItemParam(requestID, requesterID int64) CreateItemParams {
	quantityTypes := []ItemQuantityType{
		ItemQuantityTypeNumerical,
		ItemQuantityTypeOz,
		ItemQuantityTypeLbs,
		ItemQuantityTypeFlOz,
		ItemQuantityTypeGal,
		ItemQuantityTypeLitres,
	}

	return CreateItemParams{
		Name:           util.RandomString(10),
		RequestID:      requestID,
		RequestedBy:    requesterID,
		Quantity:       util.RandomFloat(0, 100),
		QuantityType:   quantityTypes[util.RandomInt(0, len(quantityTypes)-1)],
		PreferredBrand: sql.NullString{String: util.RandomString(10), Valid: true},
		ExtraNotes:     sql.NullString{String: util.RandomString(10), Valid: true},
	}
}

func addRandomItems(queries *Queries, requests []Request) []Item {

	items := []Item{}

	for _, request := range requests {
		numItemsForRequest := util.RandomInt(minNumItemsPerRequest, maxNumItemsPerRequest)
		for i := 0; i < numItemsForRequest; i++ {
			createItemParams := createRandomItemParam(request.ID, request.UserID)
			item, err := queries.CreateItem(context.Background(), createItemParams)
			if err != nil {
				log.Println("[DB POPULATE] could not create item: ", err)
				return []Item{}
			}
			items = append(items, item)
		}
	}

	return items
}

func PopulateWithData(queries *Queries) {

	if numUsers <= 0 || numCommunities <= 0 || numStores <= 0 {
		log.Fatal("[DB POPULATE] invalid number of users, communities, or stores")
	}

	users := addRandomUsers(queries, numUsers)
	if len(users) == 0 {
		log.Println("[DB POPULATE] could not populate users")
		return
	}

	// User 1 will be the admin of all communities.
	communities := addRandomCommunities(queries, numCommunities)
	if len(communities) == 0 {
		log.Println("[DB POPULATE] could not populate communities")
		return
	}

	// Add required number of stores.
	stores := addRandomStores(queries, numStores)
	if len(stores) == 0 {
		log.Println("[DB POPULATE] could not populate stores")
		return
	}

	// Add all stores to all communities.
	communityStores := addRandomCommunityStores(queries, len(communities), len(stores))
	if len(communityStores) == 0 {
		log.Println("[DB POPULATE] could not populate community stores")
		return
	}

	// Users sequentially added to communities.
	members := addRandomMembers(queries, len(users), len(communities))
	if len(members) == 0 {
		log.Println("[DB POPULATE] could not populate members")
		return
	}

	// Add a random number of requests to all communities.
	requests := addRandomRequests(queries, communityStores, members)
	if len(requests) == 0 {
		log.Println("[DB POPULATE] could not populate requests")
		return
	}

	// Add random number of items to all requests.
	items := addRandomItems(queries, requests)
	if len(items) == 0 {
		log.Println("[DB POPULATE] could not populate items")
		return
	}

	log.Println("[DB POPULATE] successfully populated database with fake data")

}
