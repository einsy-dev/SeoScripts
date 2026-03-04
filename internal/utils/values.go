package utils

import (
	"strings"
)

type Values struct {
	data [][]any
	rows map[string]int
	cols map[string]int // group.column
}

func (v *Values) New(data [][]any) {
	v.data = data
	v.cols = v.parseCols()
	v.rows = v.parseRows()
}

func (v *Values) ToMap() map[string]map[string]any {
	var res = map[string]map[string]any{}

	for key, i := range v.rows {
		var item = map[string]any{}

		for key2, i2 := range v.cols {

			addr := strings.Split(key2, ".")
			if len(addr) > 1 {
				m, ok := item[addr[0]].(map[string]any)
				if ok {
					m[addr[1]] = v.data[i][i2]
				} else {
					m = make(map[string]any)
					m[addr[1]] = v.data[i][i2]
				}
				item[addr[0]] = m
			} else {
				item[key2] = v.data[i][i2]
			}

		}
		res[key] = item
	}
	return res
}

func (v *Values) parseRows() map[string]int {
	var domains = make(map[string]int)

	for i, row := range v.data[2:] {
		if m, ok := row[0].(string); ok {
			domains[m] = i + 2
		} else {
			continue
		}
	}

	return domains
}

func (v *Values) parseCols() map[string]int {
	var group string = ""
	var cols = make(map[string]int)
	for i, col := range v.data[1] {

		if m, ok := v.data[0][i].(string); ok && m != "" {
			group = m
		}

		if m, ok := col.(string); ok {
			var key string = ""
			if group != "" {
				key = group + "." + m
			} else {
				key = m
			}
			cols[key] = i
		}
	}

	return cols
}
