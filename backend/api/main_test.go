package api

import (
	"io"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	db "github.com/git-adithyanair/cs130-group-project/db/sqlc"
	"github.com/git-adithyanair/cs130-group-project/util"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.DBStore) *Server {
	config := util.Config{
		TokenSymmetricKey: util.RandomString(32),
	}
	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)
	gin.DefaultWriter = io.Discard
	os.Exit(m.Run())

}
