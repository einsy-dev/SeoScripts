package services

import (
	"errors"
)

type Options struct {
	Keys *[]string
}

type CsvItem struct {
	Value    [][]any
	Cols     map[string]int
	Rows     map[string]int
	Key      string
	KeyIndex int
}

func (i *CsvItem) Join(csv ...*CsvItem) error {
	if len(csv) < 1 {
		return errors.New("Not enough arguments in 'Join' func")
	}

	items := append([]*CsvItem{i}, csv...)
	return Join(items...)
}

func (i *CsvItem) Update(csv ...*CsvItem) error {
	if len(csv) < 1 {
		return errors.New("Not enough arguments in 'Join' func")
	}
	items := append([]*CsvItem{i}, csv...)
	return Update(items...)
}

func (i *CsvItem) ToMap() ([]map[string]any, error) {
	if len(i.Value) <= 1 {
		return nil, errors.New("Insufficient array length. Must  be more than one")
	}
	return ToMap(i)
}

func (i *CsvItem) FormatCols(fn func(col string) string) {
	for c, v := range i.Cols {
		delete(i.Cols, c)
		c = fn(c)
		i.Cols[c] = v
		i.Value[0][v] = c
	}
}

func (i *CsvItem) FormatRows(fn func(row string) string) {
	for r, v := range i.Rows {
		delete(i.Rows, r)
		r = fn(r)
		i.Rows[r] = v
		i.Value[v][i.KeyIndex] = r
	}
}

func (i *CsvItem) Sort(h []string) {}

func (i *CsvItem) Filter(h []string) {}
