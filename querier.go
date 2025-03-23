package dbwrapper

import "context"

//go:generate go run github.com/vektra/mockery/v2@v2.42.0 --name=Querier --case=underscore

// Querier runs db queries.
type Querier interface {
	// Get retrieves a single row.
	Get(ctx context.Context, qb Sqlizer, dest any) error
	// GetList retrieves multiple rows.
	GetList(ctx context.Context, qb Sqlizer, dest any) error
	// Exec executes a query.
	Exec(ctx context.Context, qb Sqlizer) error
	// BeginTx begins a transaction.
	BeginTx(ctx context.Context) (Tx, error)
}
