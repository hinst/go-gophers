package gophers

import (
	"encoding/json"
)

func EncodeJson[T any](value T) []byte {
	return AssertResultError(json.Marshal(value))
}

func DecodeJson[T any](data []byte, value T) T {
	AssertError(json.Unmarshal(data, value))
	return value
}
