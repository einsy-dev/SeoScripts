package csvParser

import (
	s "domains/pkg/csvParser/services"
)

type Options = s.Options
type CsvItem = s.CsvItem

func Parse(csv [][]any, opt Options) (*CsvItem, error) {
	return s.Parse(csv, opt)
}
