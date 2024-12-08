package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetQueriesWithConnection(connectionString string) (*Queries, *pgxpool.Pool, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		return nil, nil, err
	}

	queries := New(pool)
	return queries, pool, nil
}
