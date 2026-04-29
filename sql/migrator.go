package sql

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

func NewMigrator(dbPath string) (*migrate.Migrate, error) {
	sqlDB, err := sqlx.Open("sqlite", dbPath)
	if err != nil {
		log.Printf("Error opening migrator db connection")
		return nil, err
	}

	driver, err := sqlite.WithInstance(sqlDB.DB, &sqlite.Config{})
	if err != nil {
		sqlDB.Close()
		log.Printf("failed to create sqlite driver: %v", err)
		return nil, fmt.Errorf("failed to create sqlite driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://sql/migrations",
		"sqlite",
		driver,
	)

	if err != nil {
		sqlDB.Close()
		log.Printf("failed to create migrator: %v", err)
		return nil, fmt.Errorf("failed to create migrator: %w", err)
	}

	return m, nil
}
