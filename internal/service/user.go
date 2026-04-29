package service

import (
	"context"
	"log"

	"go_tdd/internal/db"
	"go_tdd/internal/domain"
	"go_tdd/internal/model"
	"go_tdd/internal/repository"
)

type userService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *userService {
	return &userService{repo: repo}
}

func NewUserSqlxService() *userService {
	repo := repository.UserRepositorySqlx(db.Client(nil))
	return NewUserService(repo)
}

func (s *userService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	createdUser := &model.User{}
	createdUser, err := s.repo.CreateUser(ctx, createdUser.FromDomain(user))
	if err != nil {
		log.Print("[UserSqlxService - CreateUser]: failed to create user: %w", err)
		return nil, err
	}
	return createdUser.ToDomain(), nil
}

func (s *userService) DeleteUser(ctx context.Context, userID int) error {
	if err := s.repo.DeleteUser(ctx, userID); err != nil {
		log.Print("[UserSqlxService - DeleteUser]: failed to delete user: %w", err)
		return err
	}
	return nil
}

func (s *userService) GetAllUser(ctx context.Context, limit, offset uint) ([]*domain.User, error) {
	var users []*domain.User
	dbUsers, err := s.repo.GetAllUsers(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	for _, dbUser := range dbUsers {
		users = append(users, dbUser.ToDomain())
	}

	return users, nil
}
