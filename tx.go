package dbwrapper

import (
	"context"

	"github.com/jackc/pgx/v5"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=Tx --case=underscore

// Tx runs queries in transaction.
type Tx interface {
	Querier
	// Rollback rolls back the transaction.
	Rollback(ctx context.Context) error
	// Commit commits the transaction.
	Commit(ctx context.Context) error
}

type tx struct {
	pgTx pgx.Tx
}

func newTx(pgTx pgx.Tx) Tx {
	return &tx{
		pgTx: pgTx,
	}
}

func (t *tx) Rollback(ctx context.Context) error {
	return t.pgTx.Rollback(ctx)
}

func (t *tx) Commit(ctx context.Context) error {
	return t.pgTx.Commit(ctx)
}

func (t *tx) Get(ctx context.Context, qb Sqlizer, dest any) error {
	return get(ctx, qb, dest, t.pgTx)
}

func (t *tx) GetList(ctx context.Context, qb Sqlizer, dest any) error {
	return getList(ctx, qb, dest, t.pgTx)
}

func (t *tx) Exec(ctx context.Context, qb Sqlizer) error {
	return exec(ctx, qb, t.pgTx)
}

func (t *tx) BeginTx(ctx context.Context) (Tx, error) {
	return beginTx(ctx, t.pgTx)
}
