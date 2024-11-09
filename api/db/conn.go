package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func GetQueriesWithConnection(config *pgx.ConnConfig) (*Queries, *pgx.Conn, error) {
	ctx := context.Background()
	conn, err := pgx.ConnectConfig(ctx, config)
	if err != nil {
		return nil, nil, err
	}

	queries := New(conn)
	return queries, conn, nil
}
