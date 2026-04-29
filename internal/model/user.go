package model

import (
	"go_tdd/internal/domain"
	"time"
)

type User struct {
	ID        int        `db:"id"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

func (obj *User) TableName() string { return "users" }
func (obj *User) GetPK() string     { return "id" }

func (obj *User) ToDomain() *domain.User {
	return &domain.User{
		ID:        obj.ID,
		Name:      obj.Name,
		Email:     obj.Email,
		CreatedAt: obj.CreatedAt,
		UpdatedAt: obj.UpdatedAt,
		DeletedAt: obj.DeletedAt,
	}
}

func (obj *User) FromDomain(user *domain.User) *User {
	obj.ID = user.ID
	obj.Name = user.Name
	obj.Email = user.Email
	obj.CreatedAt = user.CreatedAt
	obj.UpdatedAt = user.UpdatedAt
	obj.DeletedAt = user.DeletedAt

	return obj
}
