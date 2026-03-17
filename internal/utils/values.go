package utils

import (
	"errors"
	"slices"
	"strings"
)

type Values struct {
	Data   [][]any
	Rows   map[string][]int
	Cols   map[string][]int
	domain int
	err    error
}

func (v *Values) New(data [][]any) error {
	v.Data = data
	v.Cols = v.parseCols()
	v.Rows = v.parseRows()
	return v.err
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

	if len(r) == 0 || len(c) == 0 {
		return
	}
	for _, v1 := range r {
		for _, v2 := range c {
			if val != nil {
				v.Data[v1][v2] = val
			} else {
				v.Data[v1][v2] = ""
			}
		}
	}
}

func (v *Values) Get(domain string, col string) any {
	if len(v.Rows[domain]) == 0 || len(v.Cols[col]) == 0 {
		return nil
	}
	r := v.Rows[domain][0]
	c := v.Cols[col][0]

	if v.Data[r][c] == "" {
		return nil
	}
	return v.Data[r][c]
}

func (v *Values) SetMap(m map[string]any) {
	for k := range v.Cols {
		keyM := strings.Split(k, ".")
		var val any
		if len(keyM) > 1 {
			val = m[keyM[0]].(map[string]any)[keyM[1]]
		} else {
			val = m[keyM[0]]
		}

		v.Set(m["Domain"].(string), k, val)
	}
}

func (val *Values) GetMap(m map[string]any) map[string]any {
	for k := range m {
		if t, ok := m[k].(map[string]any); ok {
			for k2 := range t {
				t[k2] = formatValue(val.Get(m["Domain"].(string), strings.Join([]string{k, k2}, ".")))
			}
		} else {
			m[k] = formatValue(val.Get(m["Domain"].(string), k))
		}
	}
	return m
}

func (v *Values) parseRows() map[string][]int {
	var domains = make(map[string][]int)
	if len(v.Data) <= 2 {
		v.err = errors.New("Can`t process rows. Array length is 0.")
		return nil
	}
	for i, row := range v.Data[2:] {
		if m, ok := row[v.domain].(string); ok {
			domains[m] = append(domains[m], i+2)
		}
	}

	return domains
}

func (v *Values) parseCols() map[string][]int {
	var group string = ""
	var cols = make(map[string][]int)

	if len(v.Data) == 0 {
		v.err = errors.New("Can`t process cols. Array length is 0.")
		return nil
	}

	for i, col := range v.Data[1] {

		if m, ok := v.Data[0][i].(string); ok && m != "" {
			group = m
		}

		if m, ok := col.(string); ok {
			if slices.Contains([]string{"domain"}, strings.ToLower(m)) {
				v.domain = i
			}

			var key string = ""
			if group != "" {
				key = group + "." + m
			} else {
				key = m
			}
			cols[key] = append(cols[key], i)
		}
	}

	return cols
}
