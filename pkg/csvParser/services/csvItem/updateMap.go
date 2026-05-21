package csvItem

import (
	"domains/pkg/csvParser/services/utils"
	"fmt"
	"reflect" // Need this to inspect underlying pointers
)

// keyF - key Field which is row id for csv
func UpdateMap(csv *CsvItem, keyF string, m ...map[string]any) error {
	for _, v := range m {
		fmap := utils.FlatMap(v)
		row, ok := fmap[keyF]
		if !ok || row == nil {
			continue
		}

		rowK := fmt.Sprintf("%v", row)
		keyRow, ok := csv.Rows[rowK]
		if !ok {
			continue
		}
		for key, v2 := range fmap {
			if keyCol, ok := csv.Cols[key]; ok {
				if v2 == nil || isInterfaceNil(v2) {
					csv.Value[keyRow][keyCol] = ""
				} else {
					csv.Value[keyRow][keyCol] = v2
				}
			}
		}
	}
	return nil
}

// Helper function to check if an interface contains a typed nil pointer
func isInterfaceNil(i any) bool {
	val := reflect.ValueOf(i)
	kind := val.Kind()

	if kind == reflect.Ptr || kind == reflect.Slice || kind == reflect.Map || kind == reflect.Chan || kind == reflect.Func || kind == reflect.Interface {
		return val.IsNil()
	}

	return false
}
