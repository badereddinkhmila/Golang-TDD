package db

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"sync"
	"time"

	goqu "github.com/doug-martin/goqu/v9"
	sqlx "github.com/jmoiron/sqlx"
	log "github.com/rs/zerolog/log"
	lo "github.com/samber/lo"
	fx "go.uber.org/fx"
)

var connOnce sync.Once
var conn *sqlx.DB

func initDB(lc *fx.Lifecycle) error {
	ctx, cancelConnFn := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelConnFn()

	projectRoot, _ := os.Getwd()
	dsn := filepath.Join(projectRoot, "sql/testing_database.db?cache=shared&_journal_mode=WAL")

	db, err := sqlx.Open("sqlite", dsn)
	if err != nil {
		log.Printf("failed to open database connection: %v", err)
		return err
	}

	_, err = db.ExecContext(ctx, `
        PRAGMA journal_mode=WAL;           -- Write-Ahead Logging (huge concurrency win)
        PRAGMA synchronous=NORMAL;         -- Good balance of speed/safety
        PRAGMA cache_size=-20000;          -- ~80 MB cache
        PRAGMA foreign_keys=ON;            -- Enforce FKs
        PRAGMA busy_timeout=5000;          -- 5s busy timeout
    `)
	if err != nil {
		log.Printf("failed to execute sqlite optimisations: %v", err)
		return err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(2 * time.Minute)

	if err := db.PingContext(ctx); err != nil {
		log.Printf("failed to ping sqlite database: %v", err)
		return err
	}

	log.Info().Str("dsn", dsn).Msg("SQLite connected with WAL mode")

	if lc != nil {
		lo.FromPtr(lc).Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				log.Ctx(ctx).Print("Pinging database...")
				return db.PingContext(ctx)
			},
			OnStop: func(ctx context.Context) error {
				log.Ctx(ctx).Println("Closing database connection pool...")
				return db.Close()
			},
		})
	}

	conn = db
	return nil
}

type Queryable interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryRowxContext(ctx context.Context, query string, args ...any) *sqlx.Row
	QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)
	GetContext(ctx context.Context, dest any, query string, args ...any) error
	SelectContext(ctx context.Context, dest any, query string, args ...any) error
	NamedExecContext(ctx context.Context, query string, arg any) (sql.Result, error)
	NamedQuery(query string, arg any) (*sqlx.Rows, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
}

func Client(lc *fx.Lifecycle) Queryable {
	connOnce.Do(func() { initDB(lc) })
	return conn
}

func Close() (err error) {
	if err := conn.Close(); err != nil {
		log.Logger.Debug().Err(err).Msg("Failed to close to database connection")
	}
	return err
}

func SQLBuilder() goqu.DialectWrapper {
	return goqu.Dialect("sqlite3")
}

func NamedGet[T any](ctx context.Context, q Queryable, dest *T, query string, arg any) error {
	stmt, _, err := sqlx.Named(query, arg)
	if err != nil {
		return err
	}
	return q.GetContext(ctx, dest, stmt)
}

func NamedSelect[T any](ctx context.Context, q Queryable, dest *[]T, query string, arg any) error {
	stmt, _, err := sqlx.Named(query, arg)
	if err != nil {
		return err
	}
	return q.SelectContext(ctx, dest, stmt)
}
