package enums

import (
	"database/sql/driver"
	"fmt"
)

type Status string

const (
	Active   Status = "active"
	Inactive Status = "inactive"
	Pending  Status = "pending"
)

func (s *Status) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}

	*s = Status(value.(string))

	return nil
}

func (s Status) Value() (driver.Value, error) {
	switch s {
	case Active, Inactive, Pending:
		return string(s), nil
	}
	return nil, fmt.Errorf("invalid Status value: '%s'. Must be active, inactive, or pending", s)
}
