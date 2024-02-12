package dbwrapper

import "context"

type ConnectionPool interface {
	GetSingle(ctx context.Context, dest any, query string, args ...any) error
	Get(ctx context.Context, dest any, query string, args ...any) error
	Exec(ctx context.Context, query string, args ...any) error
	BeginTx(ctx context.Context) (Tx, error)
}

type Tx interface {
	ConnectionPool
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}
