package util

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func EncodeStructToBase64String[T any](val T) (string, error) {
	str, err := json.Marshal(val)
	if err != nil {
		return "", fmt.Errorf("error marshaling json data to string: %w", err)
	}
	return base64.StdEncoding.EncodeToString(str), nil
}

func DecodeBase64StringToStruct[T any](str string, val *T) error {
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

func ComparePointerValues[T comparable](a *T, b *T) bool {
	bothNil := a == nil && b == nil
	bothNotNil := a != nil && b != nil
	return bothNil || (bothNotNil && *a == *b)
}

func HandleNoRowsOnNullableType[T any, U any](result T, err error, mapper func(T) *U) (*U, error) {
	if mapper == nil {
		err = errors.New(("mapping function cannot be nil"))
	}

	if err != nil {
		/**
		 * Check if the error is that no rows were returned.
		 * If so, just return nil, since we don't want to return an error for that
		 */
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf(("error mapping row on nullable type: %w"), err)
	}

	return mapper(result), nil
}
