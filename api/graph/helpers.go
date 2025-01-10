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
	Decode(string) error
	Validate(input T) bool
}

type ListRecipesCursor struct {
	Id    int
	Limit int
}

func (cursor *ListRecipesCursor) Decode(encoded string) error {
	return decodeBase64StringToStruct(encoded, cursor)
}

func (cursor ListRecipesCursor) Encode() (string, error) {
	return encodeStructToBase64String(cursor)
}

func (cursor ListRecipesCursor) Validate(input int) bool {
	return cursor.Limit == input
}

type ListCommentsCursor struct {
	PostDate time.Time
	Limit    int
}

func (cursor *ListCommentsCursor) Decode(encoded string) error {
	return decodeBase64StringToStruct(encoded, cursor)
}

func (cursor ListCommentsCursor) Encode() (string, error) {
	return encodeStructToBase64String(cursor)
}

func (cursor ListCommentsCursor) Validate(input int) bool {
	return cursor.Limit == input
}

func encodeStructToBase64String[T any](val T) (string, error) {
	str, err := json.Marshal(val)
	if err != nil {
		return "", fmt.Errorf("error marshaling json data to string: %w", err)
	}
	return base64.StdEncoding.EncodeToString(str), nil
}

func decodeBase64StringToStruct[T any](str string, val *T) error {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return fmt.Errorf("error decoding base64 string: %w", err)
	}

	err = json.Unmarshal(decoded, val)
	if err != nil {
		return fmt.Errorf("error unmarshaling string to struct: %w", err)
	}

	return nil
}
