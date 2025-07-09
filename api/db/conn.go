package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// QueryWrapper wraps the sqlc Querier interface and adds the `WithTx` method
type QueryWrapper interface {
	Querier
	WithTx(tx pgx.Tx) QueryWrapper
}

// QueriesWrapper implements QueryWrapper
type QueriesWrapper struct {
	*Queries
}

func (q *QueriesWrapper) WithTx(tx pgx.Tx) QueryWrapper {
	return &QueriesWrapper{
		Queries: q.Queries.WithTx(tx),
	}
}

func GetQueriesWithConnection(connectionString string) (QueryWrapper, *pgxpool.Pool, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		return nil, nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, nil, err
	}

	queries := New(pool)
	// Wrap the sqlc Queries struct in our wrapper so we can have an interface with the `WithTx` method
	return &QueriesWrapper{
		Queries: queries,
	}, pool, nil
}
