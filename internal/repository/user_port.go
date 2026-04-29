package repository

import (
	"context"

	"go_tdd/internal/model"

	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	GetAllUsers(ctx context.Context, limit, offser uint) ([]*model.User, error)
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)

	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	CreateBatchUsers(ctx context.Context, users []*model.User) ([]*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id int) error

	WithTx(tx *sqlx.Tx) IUserRepository
}
