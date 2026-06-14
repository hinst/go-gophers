package gophers

import (
	"io"
)

func IfElse[T any](condition bool, ifTrue T, ifFalse T) T {
	if condition {
		return ifTrue
	} else {
		return ifFalse
	}
}

func IoClose(v io.Closer) {
	AssertError(v.Close())
}

func IoCloseSilently(v io.Closer) {
	IgnoreError(v.Close())
}
