package server

import (
	"context"
	"go_tdd/server/routes"
	"net/http"

	"go.uber.org/fx"
)

func LaunchServer(lc fx.Lifecycle) {
	server := routes.ApplicationRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: server,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
}
