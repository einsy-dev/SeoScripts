package services

import (
	"strings"
)

// keyF - key Field which is row id for csv
func UpdateMap(csv *CsvItem, keyF string, m ...map[string]any) error {
	for _, v := range m {
		fv := flatMap("", v, nil)
		row := fv[keyF].(string)

		if row == "" {
			continue
		}

		if kr, ok := csv.Rows[fv[keyF].(string)]; ok {
			for k, v2 := range fv {
				if kc, ok := csv.Cols[k]; ok {
					csv.Value[kr][kc] = v2
				}
			}

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
