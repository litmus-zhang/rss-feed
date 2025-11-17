package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/litmus-zhang/rss-feed/internal/config"
	"github.com/litmus-zhang/rss-feed/internal/db"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func newTestServer(t *testing.T, cfg *config.Config, store db.Store, logger *zap.Logger) *Server {
	server, err := NewServer(cfg, store, logger)
	require.NoError(t, err)
	return server
}
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	code := m.Run()

	os.Exit(code)

}
