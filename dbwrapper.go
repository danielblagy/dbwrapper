package dbwrapper

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Sqlizer defines interface for converter from a query builder to sql query with args.
type Sqlizer interface {
	ToSql() (sqlQuery string, args []any, err error)
}

// Connect establishes a connection to the database and initializes a connection pool.
func Connect(ctx context.Context, url string) (ConnectionPool, error) {
	pgPool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	return &connectionPool{
		pgPool: pgPool,
	}, nil
}
