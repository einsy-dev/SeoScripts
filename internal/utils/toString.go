package utils

import "fmt"

func ToString(v *any) *string {
	if v == nil || *v == nil {
		return nil
	}

	var val = *v
	if m, ok := val.(string); ok {
		if m == "" {
			return nil
		}
		return &m
	} else {
		res := fmt.Sprintf("%v", val)
		if res == "<nil>" || res == "" {
			return nil
		}
		return &res
	}
}
