package graph

import (
	"errors"
	"fmt"
	"forkd/util"
	"time"

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
	Id    pgtype.UUID
	Limit int
}

func (cursor *ListRecipesCursor) Decode(encoded string) error {
	return util.DecodeBase64StringToStruct(encoded, cursor)
}

func (cursor ListRecipesCursor) Encode() (string, error) {
	return util.EncodeStructToBase64String(cursor)
}

func (cursor ListRecipesCursor) Validate(input int) bool {
	return cursor.Limit == input
}

type ListCommentsCursor struct {
	PostDate time.Time
	Limit    int
}

func (cursor *ListCommentsCursor) Decode(encoded string) error {
	return util.DecodeBase64StringToStruct(encoded, cursor)
}

func (cursor ListCommentsCursor) Encode() (string, error) {
	return util.EncodeStructToBase64String(cursor)
}

func (cursor ListCommentsCursor) Validate(input int) bool {
	return cursor.Limit == input
}
