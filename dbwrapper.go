package dbwrapper

import "context"

//go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=ConnectionPool --case=underscore
//go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=Tx --case=underscore

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
