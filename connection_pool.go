package dbwrapper

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=ConnectionPool --case=underscore

type ConnectionPool interface {
	Querier
	// Close closes all DB connections in the connection pool.
	Close()
}

// connectionPool implements Query interface.
type connectionPool struct {
	pgPool *pgxpool.Pool
}

// Close closes db connection pool.
func (p *connectionPool) Close() {
	p.pgPool.Close()
}

func (p *connectionPool) Get(ctx context.Context, qb Sqlizer, dest any) error {
	return get(ctx, qb, dest, p.pgPool)
}

func (p *connectionPool) GetList(ctx context.Context, qb Sqlizer, dest any) error {
	return getList(ctx, qb, dest, p.pgPool)
}

func (p *connectionPool) Exec(ctx context.Context, qb Sqlizer) error {
	return exec(ctx, qb, p.pgPool)
}

func (p *connectionPool) BeginTx(ctx context.Context) (Tx, error) {
	return beginTx(ctx, p.pgPool)
}
