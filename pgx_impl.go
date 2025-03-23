package dbwrapper

import (
	"context"
	"errors"
	"reflect"

	"database/sql"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type pgxQuerier interface {
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

func get(ctx context.Context, qb Sqlizer, dest any, q pgxQuerier) error {
	query, args, err := qb.ToSql()
	if err != nil {
		return ErrToSQLFail
	}

	err = pgxscan.Get(ctx, q, dest, query, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return sql.ErrNoRows
		}

		return err
	}

	return nil
}

func getList(ctx context.Context, qb Sqlizer, dest any, q pgxQuerier) error {
	sql, args, err := qb.ToSql()
	if err != nil {
		return ErrToSQLFail
	}

	if reflect.PointerTo(reflect.TypeOf(dest)).Kind() != reflect.Pointer {
		return ErrNotSlice
	}

	return pgxscan.Select(ctx, q, dest, sql, args...)
}

func exec(ctx context.Context, qb Sqlizer, q pgxQuerier) error {
	sql, args, err := qb.ToSql()
	if err != nil {
		return ErrToSQLFail
	}

	rows, err := q.Query(ctx, sql, args...)
	if err != nil {
		return err
	}

	// TODO for some reason using Query and not reading returned pgx.Rows object
	// makes connection busy when using pgx.Tx, so force-close is needed.
	// perhaps there's a better way to go about it
	rows.Close()

	return nil
}

func beginTx(ctx context.Context, q pgxQuerier) (Tx, error) {
	tx, err := q.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return newTx(tx), nil
}
