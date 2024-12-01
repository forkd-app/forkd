package graph

import (
	"errors"

	pgx "github.com/jackc/pgx/v5"
)

func handleNoRowsOnNullableType[T any, U any](result T, err error, mapper func(T) *U) (*U, error) {
	if err != nil {
		/**
		 * Check if the error is that no rows were returned.
		 * If so, just return nil, since we don't want to return an error for that
		 */
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return mapper(result), nil
}
