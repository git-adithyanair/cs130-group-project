package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	"github.com/git-adithyanair/cs130-group-project/token"
	"github.com/git-adithyanair/cs130-group-project/util"
)

type Server struct {
	config     util.Config
	queries    db.DBStore
	router     *gin.Engine
	tokenMaker token.Maker
}

// Initializes and returns a new Server instance.
func NewServer(config util.Config, store db.DBStore) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
		queries:    store,
	}
	server.setupRouter()

	return server, nil

}

// Set up all the routes and attach to server object.
func (server *Server) setupRouter() {

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})

	// Routes that require auth middleware.
	protectedRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// User routes.
	router.POST("/user", server.RegisterUser)
	router.POST("/user/login", server.LoginUser)
	protectedRoutes.GET("/user/community", server.GetUserCommunities)
	protectedRoutes.GET("/user/requests", server.GetUserRequest)
	protectedRoutes.POST("/user/update-location", server.UpdateUserLocation)
	protectedRoutes.POST("/user/update-profile-pic", server.UpdateUserProfilePic)
	protectedRoutes.POST("/user/update-name", server.UpdateUserName)

	// Community routes.
	protectedRoutes.POST("/community", server.CreateCommunity)
	protectedRoutes.GET("/community", server.GetAllCommunities)
	protectedRoutes.POST("/community/join", server.JoinCommunity)
	protectedRoutes.GET("/community/:id", server.GetCommunity)
	protectedRoutes.GET("/community/requests", server.GetRequestsByCommunity)
	protectedRoutes.GET("/community/stores/:id", server.GetCommunityStores)

	// Errand routes.
	protectedRoutes.POST("/errand", server.CreateErrand)
	protectedRoutes.POST("/errand/update-status", server.UpdateErrandStatus)
	protectedRoutes.POST("/errand/requests/:id", server.GetRequestsByErrand)

	// Request routes.
	protectedRoutes.POST("/request", server.CreateRequest)
	protectedRoutes.GET("/request/items/:id", server.GetItemsByRequest)
	protectedRoutes.POST("/request/update-status", server.UpdateRequestStatus)

	// Item routes.
	protectedRoutes.POST("/item/update-status", server.UpdateItemStatus)

	server.router = router
}

// Start the server.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
