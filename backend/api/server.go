package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/git-adithyanair/cs130-group-project/util"
)

type Server struct {
	config util.Config
	// store      db.Store
	router *gin.Engine
	// tokenMaker token.Maker
}

// Initializes and returns a new Server instance.
func NewServer(config util.Config) (*Server, error) {

	server := &Server{
		config: config,
	}
	server.setupRouter()

	return server, nil

}

// Set up all the routes and attach to server object.
func (server *Server) setupRouter() {

	router := gin.Default()

	// Routes that require auth middleware.
	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})

	server.router = router
}

// Start the server.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
