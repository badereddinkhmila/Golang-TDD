package scripts

import (
	"fmt"
	"go_tdd/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
)

func RunMigrateUp(dbPath string) error {
	log.Printf("the default db path is %s", dbPath)
	m, err := sql.NewMigrator(dbPath)
	if err != nil {
		return err
	}
	defer m.Close()

	fmt.Println("Applying migrations...")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration up failed: %w", err)
	}

	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		log.Fatal(err)
	}

	if err == migrate.ErrNilVersion {
		fmt.Println("No migrations have been applied yet")
	} else {
		fmt.Printf("Current migration version: %d (dirty: %v)\n", version, dirty)
	}

	fmt.Println("✅ Migrations applied successfully (or already up to date)")
	return nil
}
