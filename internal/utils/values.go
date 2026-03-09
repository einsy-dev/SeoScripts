package utils

import (
	"fmt"
	"strings"
)

type Values struct {
	Data [][]any
	Rows map[string]int
	Cols map[string]int // group.column
}

func (v *Values) New(data [][]any) {
	v.Data = data
	v.Cols = v.parseCols()
	v.Rows = v.parseRows()
}

func (v *Values) GetKeyRows() []string {
	var res = []string{}
	for i := range v.Rows {
		res = append(res, i)
	}
	return res
}

func (v *Values) Set(domain string, col string, val any) {
	r := v.Rows[domain]
	c := v.Cols[col]
	if val != nil {
		fmt.Println(domain, r, col, c, val)
		v.Data[r][c] = val
		if val == "google.com" {
			fmt.Println(v.Data)
		}
	}
}

func (v *Values) ToMap() map[string]map[string]any {
	var res = map[string]map[string]any{}

	for key, i := range v.Rows {
		var item = map[string]any{}

		for key2, i2 := range v.Cols {

			addr := strings.Split(key2, ".")
			if len(addr) > 1 {
				m, ok := item[addr[0]].(map[string]any)
				if ok {
					m[addr[1]] = v.Data[i][i2]
				} else {
					m = make(map[string]any)
					m[addr[1]] = v.Data[i][i2]
				}
				item[addr[0]] = m
			} else {
				item[key2] = v.Data[i][i2]
			}

		}
		res[key] = item
	}
	return res
}

func (v *Values) parseRows() map[string]int {
	var domains = make(map[string]int)

	for i, row := range v.Data[2:] {
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
	for i, col := range v.Data[1] {

		if m, ok := v.Data[0][i].(string); ok && m != "" {
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
