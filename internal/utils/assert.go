package utils

import "errors"

func Assert[T any](v any) (T, error) {
	if m, ok := v.(T); ok {
		return m, nil
	}
	var zero T
	return zero, errors.New("Type assertion failed")
}
