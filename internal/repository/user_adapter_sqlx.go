package repository

import (
	"context"
	"log"
	"time"

	"go_tdd/internal/db"
	"go_tdd/internal/model"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

type userRepositorySqlx struct {
	db db.Queryable
}

func UserRepositorySqlx(db db.Queryable) userRepositorySqlx {
	return userRepositorySqlx{db}
}

func (r userRepositorySqlx) WithTx(tx *sqlx.Tx) IUserRepository {
	return &userRepositorySqlx{db: tx}
}

/*************************************************************************/
/********************************    Read    *****************************/
/*************************************************************************/
func (r userRepositorySqlx) GetAllUsers(ctx context.Context, limit, offset uint) (users []*model.User, err error) {
	sql, args, err := db.SQLBuilder().From((&model.User{}).TableName()).
		Select("*").Limit(limit).Offset(offset * limit).ToSQL()

	if err != nil {
		panic(err)
	}

	if err = r.db.SelectContext(ctx, &users, sql, args...); err != nil {
		return nil, err
	}

	return users, nil
}

func (r userRepositorySqlx) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	return nil, nil
}

func (r userRepositorySqlx) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return nil, nil
}

/*************************************************************************/
/********************************    Write   *****************************/
/*************************************************************************/

func (r userRepositorySqlx) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	sql, args, err := db.SQLBuilder().Insert(user.TableName()).
		Rows(goqu.Record{
			"name":       user.Name,
			"email":      user.Email,
			"created_at": time.Now(),
		}).Returning("*").ToSQL()
	if err != nil {
		panic(err)
	}

	if err = r.db.QueryRowxContext(ctx, sql, args...).StructScan(user); err != nil {
		log.Printf("[UserSqlxRepository - CreateUser]: failed to create user: %s", err)
		return nil, err
	}

	return user, nil
}

func (r userRepositorySqlx) CreateBatchUsers(ctx context.Context, users []*model.User) ([]*model.User, error) {
	sql, args, err := db.SQLBuilder().Insert((&model.User{}).TableName()).
		Rows(users).Returning("*").ToSQL()
	if err != nil {
		panic(err)
	}

	if err = r.db.SelectContext(ctx, &users, sql, args...); err != nil {
		return nil, err
	}

	return users, nil
}

func (r userRepositorySqlx) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	sql, args, err := db.SQLBuilder().Update(user.TableName()).
		Set(user).Where(goqu.C("id").Eq(user.ID)).Returning("*").ToSQL()

	if err = r.db.QueryRowxContext(ctx, sql, args...).StructScan(&user); err != nil {
		log.Printf("[UserSqlxRepository - UpdateUser]: failed to create user: %s", err)
		return nil, err
	}

	return user, nil
}

func (r userRepositorySqlx) DeleteUser(ctx context.Context, id int) error {
	sql, args, err := db.SQLBuilder().Delete((&model.User{}).TableName()).
		Where(goqu.C("id").Eq(id)).Returning("*").ToSQL()

	if err = r.db.QueryRowxContext(ctx, sql, args...).Err(); err != nil {
		log.Printf("[UserSqlxRepository - DeleteUser]: failed to delete user: %s", err)
		return err
	}

	return nil
}
