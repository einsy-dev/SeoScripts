package services

import (
	"slices"
)

func Parse(data [][]any, opt Options) (*CsvItem, error) {
	res := CsvItem{Value: data, Cols: map[string]int{}, Rows: map[string]int{}}

	// get keys (columns)
	for i, v := range data[0] {
		if m, ok := v.(string); ok {
			res.Cols[m] = i
		}
	}

	res.KeyIndex = slices.IndexFunc(res.Value[0], func(col any) bool {
		if opt.Keys != nil {
			return slices.Contains(*opt.Keys, col.(string))
		}
		return false
	})

	if res.KeyIndex == -1 {
		res.KeyIndex = 0
	}

	// get rows
	for i, v := range data[1:] {
		if m, ok := v[res.KeyIndex].(string); ok {
			res.Rows[m] = i + 1
		}
	}

	return &res, nil
}
