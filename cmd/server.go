package cmd

import (
	"go_tdd/internal/db"
	"go_tdd/internal/repository"
	"go_tdd/internal/service"
	"go_tdd/server"

	cobra "github.com/spf13/cobra"
	fx "go.uber.org/fx"
)

var launchServerCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Go TDD Server with all dependencies",
	Run: func(cmd *cobra.Command, args []string) {
		launchServer()
	},
}

func launchServer() {
	fx.New(
		fx.Provide(db.Client, repository.UserRepositorySqlx,
			service.NewUserService,
		),
		fx.Invoke(server.LaunchServer),
	).Run()
}
