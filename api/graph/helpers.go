package graph

import (
	"errors"
	"fmt"

	pgx "github.com/jackc/pgx/v5"
)

func handleNoRowsOnNullableType[T any, U any](result T, err error, mapper func(T) *U) (*U, error) {
	if mapper == nil {
		err = errors.New(("Mapping function cannot be nil"))
	}

	if err != nil {
		/**
		 * Check if the error is that no rows were returned.
		 * If so, just return nil, since we don't want to return an error for that
		 */
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf(("Error mapping row on nullable type: %w"), err)
	}

	return mapper(result), nil
}
