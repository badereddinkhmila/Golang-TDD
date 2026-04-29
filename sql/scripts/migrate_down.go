package scripts

import (
	"fmt"
	"go_tdd/sql"

	"github.com/golang-migrate/migrate/v4"
)

func RunMigrateDown(dbPath string) error {
	m, err := sql.NewMigrator(dbPath)
	if err != nil {
		return err
	}
	defer m.Close()

	fmt.Println("Rolling back last migration...")
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration down failed: %w", err)
	}

	fmt.Println("✅ Rollback completed")
	return nil
}
