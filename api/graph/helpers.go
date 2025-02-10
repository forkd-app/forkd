package graph

import (
	"errors"
	"fmt"
	"forkd/graph/model"
	"forkd/util"

	pgx "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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

type Cursor[T any] interface {
	Encode() (string, error)
	Decode(string) error
	Validate(input T) bool
}

type ListRecipesCursor struct {
	model.ListRecipeInput
	PublishCursor pgtype.Timestamp
	SlugCursor    pgtype.Text
}

func (cursor *ListRecipesCursor) Decode(encoded string) error {
	return util.DecodeBase64StringToStruct(encoded, cursor)
}

func (cursor ListRecipesCursor) Encode() (string, error) {
	return util.EncodeStructToBase64String(cursor)
}

func (cursor ListRecipesCursor) Validate(input ListRecipesCursor) bool {
	return comparePointerValues(cursor.Limit, input.Limit) &&
		comparePointerValues(cursor.SortCol, input.SortCol) &&
		comparePointerValues(cursor.SortDir, input.SortDir) &&
		comparePointerValues(cursor.AuthorID, input.AuthorID) &&
		comparePointerValues(cursor.PublishStart, input.PublishStart) &&
		comparePointerValues(cursor.PublishEnd, input.PublishEnd)
}

func comparePointerValues[T comparable](a *T, b *T) bool {
	bothNil := a == nil && b == nil
	bothNotNil := a != nil && b != nil
	return bothNil || (bothNotNil && *a == *b)
}
