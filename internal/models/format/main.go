package format

import "domains/internal/utils"

func Format[T any](field T, value any) {
	switch f := any(field).(type) {
	case *string:
		res := utils.ToString(value)
		if res != nil {
			*f = *res // Dereference both to perform the assignment
		}
	case *uint:
		res := utils.ToUint(value)
		if res != nil {
			*f = *res // Dereference both to perform the assignment
		}
	}
}
