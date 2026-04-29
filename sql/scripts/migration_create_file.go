package scripts

import (
	"fmt"
	"os"
	"time"
)

func RunCreateMigration(name string) error {
	timestamp := time.Now().Unix()

	upFile := fmt.Sprintf("sql/migrations/%d_%s.up.sql", timestamp, name)
	downFile := fmt.Sprintf("sql/migrations/%d_%s.down.sql", timestamp, name)

	if err := os.WriteFile(upFile, []byte("-- Write your UP migration here\n"), 0644); err != nil {
		return err
	}
	if err := os.WriteFile(downFile, []byte("-- Write your DOWN migration here\n"), 0644); err != nil {
		return err
	}

	fmt.Printf("✅ Created migration files:\n   %s\n   %s\n", upFile, downFile)
	return nil
}
