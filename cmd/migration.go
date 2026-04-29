package cmd

import (
	"fmt"
	"go_tdd/sql/scripts"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	dbPath  string
	verbose bool
)

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Apply all pending migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		return scripts.RunMigrateUp(dbPath)
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Rollback the last migration",
	RunE: func(cmd *cobra.Command, args []string) error {
		return scripts.RunMigrateDown(dbPath)
	},
}

var migrateCreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new migration file pair",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return scripts.RunCreateMigration(args[0])
	},
}

func init() {
	projectRoot, _ := os.Getwd()
	dbPath := filepath.Join(projectRoot, "go_tdd/sql/testing_database.db")
	migrationCmd.PersistentFlags().StringVarP(&dbPath, "db", "d", fmt.Sprintf("%s?cache=shared&_journal_mode=WAL", dbPath), "SQLite DSN")
	migrationCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")

	migrationCmd.AddCommand(migrateUpCmd, migrateDownCmd, migrateCreateCmd)
}
