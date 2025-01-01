package graph

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"

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

type Cursor[T any] interface {
	Encode() (string, error)
	Decode(string) (*Cursor[T], error)
	Validate(input T) bool
}

type ListRecipesCursor struct {
	Id    int
	Limit int
}

func (cursor *ListRecipesCursor) Decode(encoded string) (*ListRecipesCursor, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(decoded, cursor)
	if err != nil {
		return nil, err
	}

	return cursor, nil
}

func (cursor ListRecipesCursor) Encode() (string, error) {
	str, err := json.Marshal(cursor)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(str), nil
}

func (cursor ListRecipesCursor) Validate(input int) bool {
	return cursor.Limit == input
}

type ListCommentsCursor struct {
	PostDate time.Time
	Limit    int
}

func (cursor *ListCommentsCursor) Decode(encoded string) (*ListCommentsCursor, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(decoded, cursor)
	if err != nil {
		return nil, err
	}

	return cursor, nil
}

func (cursor ListCommentsCursor) Encode() (string, error) {
	str, err := json.Marshal(cursor)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(str), nil
}

func (cursor ListCommentsCursor) Validate(input int) bool {
	return cursor.Limit == input
}
