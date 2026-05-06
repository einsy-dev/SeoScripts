package services

func Join(csv ...*CsvItem) error {
	for _, csvItem := range csv[1:] { // iterate each csv struct
		mergeMeta(csv[0], csvItem)
		for rv, i := range csvItem.Rows { // iterate iterate csv struct rows
			var rowI = csv[0].Rows[rv]
			for cv, j := range csvItem.Cols { // iterate iterate csv struct cols
				var colI = csv[0].Cols[cv]
				csv[0].Value[rowI][colI] = csvItem.Value[i][j]
			}
		}
	}

	return nil
}

func mergeMeta(csv1 *CsvItem, csv2 *CsvItem) {
	for v := range csv2.Rows {
		if _, ok := csv1.Rows[v]; !ok {
			csv1.Rows[v] = len(csv1.Rows)
		}
	}

	for v := range csv2.Cols {
		if _, ok := csv1.Cols[v]; !ok {
			csv1.Cols[v] = len(csv1.Cols)
		}
	}

	targetRows := len(csv1.Rows)
	if len(csv1.Value) < targetRows {
		newRows := make([][]any, targetRows-len(csv1.Value))
		csv1.Value = append(csv1.Value, newRows...)
	}

	targetCols := len(csv1.Cols)
	for i := range csv1.Value {
		if len(csv1.Value[i]) < targetCols {
			newCols := make([]any, targetCols-len(csv1.Value[i]))
			csv1.Value[i] = append(csv1.Value[i], newCols...)
		}
	}

	// It`s a fix for columns is null after resize
	for col, i := range csv1.Cols {
		csv1.Value[0][i] = col
	}
}
