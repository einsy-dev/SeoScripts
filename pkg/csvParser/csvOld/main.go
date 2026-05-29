package csv

import (
	"strings"
)

func Arr2dToMap(data [][]any) []map[string]any {
	var cols = make([]string, len(data[0]))
	var res = []map[string]any{}

	var group string

	for i, col := range data[1] {
		if data[0][i] != "" {
			group = data[0][i].(string)
		}

		if group != "" {
			cols[i] = group + "." + col.(string)
		} else {
			cols[i] = col.(string)
		}
	}

	for _, row := range data[2:] {
		var val = make(map[string]any)
		for j, col := range row {
			var addr = strings.Split(cols[j], ".")

			if len(addr) == 1 {
				val[addr[0]] = col
			} else {
				if m, ok := val[addr[0]].(map[string]any); ok {
					m[addr[1]] = col
				} else {
					m = make(map[string]any)
					m[addr[1]] = col
					val[addr[0]] = m
				}
			}
		}
		res = append(res, val)
	}
	return res
}

func Get2dArrKeys(data [][]any) (map[string]int, map[string]int) {
	var domains = make(map[string]int)
	var columns = make(map[string]int)

	var group string

	for i, col := range data[1] {
		if data[0][i] != "" {
			group = data[0][i].(string)
		}

		if group != "" {
			columns[group+"."+col.(string)] = i
		} else {
			columns[col.(string)] = i
		}
	}

	for i, row := range data[2:] {
		var domainCol = columns["Domain"]
		domains[row[domainCol].(string)] = i + 2
	}

	return domains, columns
}
