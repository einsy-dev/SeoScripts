package utils

import "slices"

func ParseValues(val [][]any) []map[string]any {
	var g, k = getKeys(&val)
	var res = []map[string]any{}
	for _, row := range val[2:] {
		var item = make(map[string]any)
		for j, key := range k {
			if g[j] != "" {
				if m, ok := item[g[j]].(map[string]any); ok {
					m[key] = row[j]
				} else {
					item[g[j]] = make(map[string]any)
					item[g[j]].(map[string]any)[key] = row[j]
				}
			} else {
				item[key] = row[j]
			}
		}
		res = append(res, item)
	}

	return res
}

var groups = []string{"Ahrefs", "Semrush", "Majestic", "Moz"}

func getKeys(val *[][]any) (map[int]string, map[int]string) {
	var m = make(map[int]string)
	var memo string
	for i, col := range (*val)[0] {
		if col == "" && memo == "" {
			continue
		}

		if slices.Contains(groups, col.(string)) {
			memo = col.(string)
		}
		m[i] = memo
	}

	var k = make(map[int]string)
	for i, col := range (*val)[1] {
		k[i] = col.(string)
	}

	return m, k
}
