package pgximpl

import (
	"context"

	"github.com/danielblagy/dbwrapper"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type tx struct {
	pgTx pgx.Tx
}

func newTx(pgTx pgx.Tx) dbwrapper.Tx {
	return &tx{
		pgTx: pgTx,
	}
}

func (t *tx) BeginTx(ctx context.Context) (dbwrapper.Tx, error) {
	nestedTx, err := t.pgTx.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return newTx(nestedTx), nil
}

func (t *tx) Rollback(ctx context.Context) error {
	return t.pgTx.Rollback(ctx)
}

func (t *tx) Commit(ctx context.Context) error {
	return t.pgTx.Commit(ctx)
}

func (t *tx) GetSingle(ctx context.Context, dest any, query string, args ...any) error {
	return pgxscan.Get(ctx, t.pgTx, dest, query, args...)
}

func (t *tx) Get(ctx context.Context, dest any, query string, args ...any) error {
	return pgxscan.Select(ctx, t.pgTx, dest, query, args...)
}

func (t *tx) Exec(ctx context.Context, query string, args ...any) error {
	_, err := t.pgTx.Query(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
