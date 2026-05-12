package services

import (
	"maps"
	"strings"
)

// keyF - key Field which is row id for csv
// !!mistake!! each map is uniq and shouldnt be merged
func UpdateMap(csv *CsvItem, keyF string, m ...map[string]any) error {
	acc := flatMap("", m[0], nil)

	for _, v := range m {
		maps.Copy(acc, flatMap("", v, nil))
	}

	// set values of flat map to csv
	for k, v := range acc {
		if i, ok := csv.Cols[k]; ok {
			csv.Value[i][0] = v
		}
	}

	return nil
}

func flatMap(key string, m map[string]any, acc map[string]any) map[string]any {
	if acc == nil {
		acc = make(map[string]any)
	}

	for k, v := range m {
		nk := k
		if key != "" {
			nk = strings.Join([]string{key, k}, ".")
		}

		if m, ok := v.(map[string]any); ok {
			flatMap(nk, m, acc)
		} else {
			acc[nk] = v
		}
	}
	return acc
}
