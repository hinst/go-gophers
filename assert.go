package gophers

// Panic if error is present
func AssertError(e error) {
	if e != nil {
		panic(e)
	}
}

func AssertResultError[T any](result T, e error) T {
	AssertError(e)
	return result
}

func AssertCondition[T any](condition bool, exception func() T) {
	if !condition {
		panic(exception())
	}
}

func IgnoreError(err error) {
}
