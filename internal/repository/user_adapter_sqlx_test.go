package repository_test

import (
	"context"
	"database/sql"
	"go_tdd/internal/model"
	"go_tdd/internal/repository"
	"testing"

	"github.com/jmoiron/sqlx"
)

type FakeDB struct {
	SelectContextFunc    func(ctx context.Context, dest any, query string, args ...any) error
	GetContextFunc       func(ctx context.Context, dest any, query string, args ...any) error
	ExecContextFunc      func(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContextFunc     func(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContextFunc  func(ctx context.Context, query string, args ...any) *sql.Row
	QueryRowxContextFunc func(ctx context.Context, query string, args ...any) *sqlx.Row
	QueryxContextFunc    func(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)
	NamedExecContextFunc func(ctx context.Context, query string, arg any) (sql.Result, error)
	NamedQueryFunc       func(query string, arg any) (*sqlx.Rows, error)
	PreparexContextFunc  func(ctx context.Context, query string) (*sqlx.Stmt, error)
}

func (f *FakeDB) SelectContext(ctx context.Context, dest any, query string, args ...any) error {
	if f.SelectContextFunc != nil {
		return f.SelectContextFunc(ctx, dest, query, args...)
	}
	panic("SelectContext not implemented")
}
func (f *FakeDB) GetContext(ctx context.Context, dest any, query string, args ...any) error {
	if f.GetContextFunc != nil {
		return f.GetContextFunc(ctx, dest, query, args...)
	}
	panic("GetContext not implemented")
}
func (f *FakeDB) QueryRowxContext(ctx context.Context, query string, args ...any) *sqlx.Row {
	if f.QueryRowxContextFunc != nil {
		return f.QueryRowxContextFunc(ctx, query, args...)
	}
	panic("QueryRowxContext not implemented")
}
func (f *FakeDB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	panic("not needed for this test")
}
func (f *FakeDB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	panic("not needed for this test")
}
func (f *FakeDB) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	panic("not needed for this test")
}
func (f *FakeDB) QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error) {
	panic("not needed for this test")
}
func (f *FakeDB) NamedExec(query string, arg any) (sql.Result, error) {
	panic("not needed for this test")
}
func (f *FakeDB) NamedExecContext(ctx context.Context, query string, arg any) (sql.Result, error) {
	panic("not needed for this test")
}
func (f *FakeDB) NamedQuery(query string, arg any) (*sqlx.Rows, error) {
	panic("not needed for this test")
}
func (f *FakeDB) PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error) {
	panic("not needed for this test")
}

func TestGetAllUsers(t *testing.T) {
	db := &FakeDB{
		SelectContextFunc: func(ctx context.Context, dest any, query string, args ...any) error {
			users := dest.(*[]*model.User)
			*users = []*model.User{
				{ID: 1, Name: "A"},
				{ID: 2, Name: "B"},
			}
			return nil
		},
	}

	repo := repository.UserRepositorySqlx(db)

	users, err := repo.GetAllUsers(context.Background(), 10, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(users) != 2 {
		t.Fatalf("expected 2 users, got %d", len(users))
	}
}
