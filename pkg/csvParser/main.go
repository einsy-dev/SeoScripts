package csvParser

import (
	s "domains/pkg/csvParser/services"
	"domains/pkg/csvParser/services/csvItem"
	"domains/pkg/csvParser/services/utils"
)

type Options = csvItem.Options
type CsvItem = csvItem.CsvItem

func Parse(csv [][]any, opt Options) (*CsvItem, error) {
	return s.Parse(csv, opt)
}

func Read(v string) ([][]string, error) {
	return utils.Read(v)
}
