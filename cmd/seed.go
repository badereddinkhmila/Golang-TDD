package cmd

import (
	"fmt"
	"go_tdd/sql/seed"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

var dbDSN string

var seedCmd = &cobra.Command{
	Use:   "seed [nbrUsers]",
	Short: "Seed the database with users",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		nbrUsers, err := strconv.ParseInt(args[0], 10, 8)
		if err != nil {
			fmt.Printf("error parsing nbr of users %v", err)
		}
		log.Printf("the number of users to seed %d", nbrUsers)

		return seed.SeedUsers(dbDSN, int(nbrUsers))
	},
}

func init() {
	projectRoot, _ := os.Getwd()
	dbDSN = filepath.Join(projectRoot, "sql/testing_database.db")
	seedCmd.PersistentFlags().StringVarP(&dbDSN, "db", "d", fmt.Sprintf("%s?cache=shared&_journal_mode=WAL", dbDSN), "SQLite DSN")
}
