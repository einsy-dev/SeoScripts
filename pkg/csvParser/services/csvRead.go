package services

import (
	c "encoding/csv"
	"strings"
)

func CsvRead(v string) ([][]string, error) {
	r := c.NewReader(strings.NewReader(v))
	record, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return record, nil
}
