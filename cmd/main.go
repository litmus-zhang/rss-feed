package main

import (
	"github.com/litmus-zhang/rss-feed/internal/api"
	"github.com/litmus-zhang/rss-feed/internal/config"
	"github.com/litmus-zhang/rss-feed/internal/db"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	app := fx.New(
		config.Module,
		db.Module,
		api.Module,
		fx.Provide(zap.NewProduction),
		fx.Invoke(func(lc fx.Lifecycle, cfg *config.Config, server *api.Server) error {
			return server.Start()
		}),
	)
	app.Run()
}
