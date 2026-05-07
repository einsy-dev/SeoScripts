package services

import (
	"fmt"
	"strings"
)

func ToMap(csv *CsvItem) ([]map[string]any, error) {
	var res = []map[string]any{}

	for _, row := range csv.Value[1:] {
		var val = make(map[string]any)
		for j, col := range row {
			key := fmt.Sprint(csv.Value[0][j])
			deepMap(val, key, col)
		}
		res = append(res, val)
	}
	return res, nil
}

func deepMap(m map[string]any, key string, value any) {
	parts := strings.Split(key, ".")
	current := m

	for i := 0; i < len(parts)-1; i++ {
		p := parts[i]
		next, ok := current[p].(map[string]any)
		if !ok {
			next = make(map[string]any)
			current[p] = next
		}
		current = next
	}
	current[parts[len(parts)-1]] = value
}
