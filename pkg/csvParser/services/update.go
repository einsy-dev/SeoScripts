package services

func Update(csv ...*CsvItem) error {
	for _, csvItem := range csv[1:] { // iterate each csv struct

		for rv, i := range csvItem.Rows { // iterate iterate csv struct rows
			var rowI = csv[0].Rows[rv]
			for cv, j := range csvItem.Cols { // iterate iterate csv struct cols
				if colI, ok := csv[0].Cols[cv]; ok {
					csv[0].Value[rowI][colI] = csvItem.Value[i][j]
				}
			}
		}

	}

	return nil
}
