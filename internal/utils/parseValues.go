package utils

func ParseValues(val [][]any) []map[string]any {
	var g, k = getKeys(&val)
	var res = []map[string]any{}
	for _, row := range val[2:] {
		var item = make(map[string]any)
		for j, key := range k {
			if m, ok := item[g[j]].(map[string]any); ok {
				m[key] = row[j]
			} else {
				item[g[j]] = make(map[string]any)
				item[g[j]].(map[string]any)[key] = row[j]
			}
		}
		res = append(res, item)
	}

	return res
}

func getKeys(val *[][]any) (map[int]string, map[int]string) {
	var m = make(map[int]string)
	var memo string
	for i, col := range (*val)[0] {
		if col == "" && memo == "" {
			continue
		}
		if col != "" {
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
