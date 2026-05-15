package format

import (
	"domains/internal/utils"
)

func Format[T any](field T, value any) {
	switch f := any(field).(type) {
	case *string:
		if res := utils.ToString(value); res != nil {
			*f = *res
		}
	case **string:
		if res := utils.ToString(value); res != nil {
			if *f == nil {
				*f = new(string)
			}
			**f = *res
		}
	case *uint:
		if res := utils.ToUint(value); res != nil {
			*f = *res
		}
	case **uint:
		if res := utils.ToUint(value); res != nil {
			if *f == nil {
				*f = new(uint)
			}
			**f = *res
		}

	case *bool:
		// Safely check if value is actually a bool
		if b, ok := value.(bool); ok {
			*f = b
		} else if s, ok := value.(string); ok {
			// Optional: handle string "true"/"false" if coming from CSV
			if s == "true" {
				*f = true
			}
			if s == "false" {
				*f = false
			}
		}

	case **bool:
		if b, ok := value.(bool); ok {
			if *f == nil {
				*f = new(bool)
			}
			**f = b
		}
	}
}
