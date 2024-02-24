package dbwrapper

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=Query --case=underscore
//go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=Tx --case=underscore

// Sqlizer defines interface for converter from a query builder to sql query with args.
type Sqlizer interface {
	ToSql() (sqlQuery string, args []any, err error)
}

// Query defines interface for query runner.
type Query interface {
	// Get retrieves a single row.
	Get(ctx context.Context, qb Sqlizer, dest any) error
	// GetList retrieves multiple rows.
	GetList(ctx context.Context, qb Sqlizer, dest any) error
	// Exec executes a query.
	Exec(ctx context.Context, qb Sqlizer) error
	// BeginTx begins a transaction.
	BeginTx(ctx context.Context) (Tx, error)
}

// Tx defines interface for transaction runner.
type Tx interface {
	Query
	// Rollback rolls back the transaction.
	Rollback(ctx context.Context) error
	// Commit commits the transaction.
	Commit(ctx context.Context) error
}

// Connect establishes a connection to the database and initializes a connection pool.
func Connect(ctx context.Context, url string) (*connectionPool, error) {
	pgPool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	return &connectionPool{
		pgPool: pgPool,
	}, nil
}
