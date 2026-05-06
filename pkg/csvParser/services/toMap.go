package services

import (
	"fmt"
)

func ToMap(csv *CsvItem) ([]map[string]any, error) {
	var res = []map[string]any{}

	for _, row := range csv.Value[1:] {
		var val = make(map[string]any)
		for j, col := range row {
			key := fmt.Sprint(csv.Value[0][j])
			val[key] = col
		}
		res = append(res, val)
	}
	return res, nil
}
