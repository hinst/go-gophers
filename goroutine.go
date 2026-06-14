package gophers

import (
	"runtime"
	"strconv"
	"strings"
)

func GetGoRoutineId() int64 {
	var buffer = make([]byte, 100)
	runtime.Stack(buffer, false)
	var text = string(buffer)
	text = strings.TrimPrefix(text, "goroutine ")
	var before, _, isFound = strings.Cut(text, " ")
	AssertCondition(isFound, func() string {
		return "Cannot parse: " + text
	})
	return AssertResultError(strconv.ParseInt(before, 10, 64))
}
