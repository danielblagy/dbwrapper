package dbwrapper

import (
	"context"
	"errors"
	"reflect"

	"database/sql"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type querier interface {
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

func get(ctx context.Context, qb Sqlizer, dest any, q querier) error {
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

func getList(ctx context.Context, qb Sqlizer, dest any, q querier) error {
	sql, args, err := qb.ToSql()
	if err != nil {
		return ErrToSQLFail
	}

	if reflect.PointerTo(reflect.TypeOf(dest)).Kind() != reflect.Pointer {
		return ErrNotSlice
	}

	return pgxscan.Select(ctx, q, dest, sql, args...)
}

func exec(ctx context.Context, qb Sqlizer, q querier) error {
	sql, args, err := qb.ToSql()
	if err != nil {
		return ErrToSQLFail
	}

	_, err = q.Query(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func beginTx(ctx context.Context, q querier) (Tx, error) {
	tx, err := q.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return newTx(tx), nil
}
