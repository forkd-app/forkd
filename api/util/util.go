package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
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
