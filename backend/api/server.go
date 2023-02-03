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
	queries    *db.Queries
	router     *gin.Engine
	tokenMaker token.Maker
}

// Initializes and returns a new Server instance.
func NewServer(config util.Config, queries *db.Queries) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
		queries:    queries,
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

	router.POST("/user", server.RegisterUser)

	// Routes that require auth middleware.
	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	server.router = router
}

// Start the server.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// Function to convert errors into a form that can be sent as a response.
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
