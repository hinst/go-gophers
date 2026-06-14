package gophers

import (
	"encoding/xml"
)

func DecodeXml[T any](data []byte, value T) T {
	AssertError(xml.Unmarshal(data, value))
	return value
}
