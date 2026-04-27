package enums

import (
	"database/sql/driver"
	"fmt"
)

type Rel string

const (
	Follow   Rel = "follow"
	Nofollow Rel = "nofollow"
)

func (s *Rel) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}

	*s = Rel(value.(string))

	return nil
}

func (s Rel) Value() (driver.Value, error) {
	switch s {
	case Follow, Nofollow:
		return string(s), nil
	}
	return nil, fmt.Errorf("invalid Rel value: '%s'. Must be active, inactive, or pending", s)
}
