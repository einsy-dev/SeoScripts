package utils

import "fmt"

func ToString(v any) *string {
	if m, ok := v.(string); ok {
		return &m
	}

	var res = fmt.Sprintf("%v", v)
	return &res
}
