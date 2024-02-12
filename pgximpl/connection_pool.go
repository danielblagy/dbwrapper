package pgximpl

import (
	"context"

	"github.com/danielblagy/dbwrapper"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, url string) (dbwrapper.ConnectionPool, error) {
	pgPool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	return newConnectionPool(pgPool), nil
}

type connectionPool struct {
	pgPool *pgxpool.Pool
}

func newConnectionPool(pgPool *pgxpool.Pool) dbwrapper.ConnectionPool {
	return &connectionPool{
		pgPool: pgPool,
	}
}

func (p *connectionPool) GetSingle(ctx context.Context, dest any, query string, args ...any) error {
	return pgxscan.Get(ctx, p.pgPool, dest, query, args...)
}

func (p *connectionPool) Get(ctx context.Context, dest any, query string, args ...any) error {
	return pgxscan.Select(ctx, p.pgPool, dest, query, args...)
}

func (p *connectionPool) Exec(ctx context.Context, query string, args ...any) error {
	_, err := p.pgPool.Query(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (p *connectionPool) BeginTx(ctx context.Context) (dbwrapper.Tx, error) {
	tx, err := p.pgPool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return newTx(tx), nil
}
