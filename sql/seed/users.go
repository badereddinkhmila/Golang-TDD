package seed

import (
	"context"
	"fmt"
	"go_tdd/internal/model"
	"go_tdd/internal/repository"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/jmoiron/sqlx"
)

func SeedUsers(dbPath string, nbrUsers int) error {
	sqlDB, err := sqlx.Open("sqlite", dbPath)
	if err != nil {
		log.Printf("Error opening migrator db connection")
		return err
	}
	defer sqlDB.Close()

	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	repo := repository.UserRepositorySqlx(sqlDB)
	var users []*model.User

	for idx := range nbrUsers {
		users = append(users, &model.User{
			ID:        idx,
			Name:      fmt.Sprintf("User %d", idx),
			Email:     fmt.Sprintf("user%d.email@test.com", idx),
			CreatedAt: time.Now(),
		})
	}

	if _, err := repo.CreateBatchUsers(context.Background(), users); err != nil {
		log.Printf("failed to seed users %v", err)
		return err
	}

	return nil
}
