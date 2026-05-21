package utils

import (
	"encoding/csv"
	"strings"
)

func Read(v string) ([][]string, error) {
	r := csv.NewReader(strings.NewReader(v))
	record, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return record, nil
}
