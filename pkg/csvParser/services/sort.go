package services

import (
	"sort"
)

func Sort(csvItem *CsvItem, sortArr []string) error {

	for _, v := range csvItem.Value {
		sort.Slice(v, func(i, j int) bool {
			return true
			// return slices.Index(sortArr, v[i]) > slices.IndexFunc(sortArr, "") // Sort by string length
		})
	}

	return nil
}

// function filterCSV(csv: string[][], columns: string[]): string[][] {
// 		if (!columns.length || !csv.length) return csv;
// 		let res: string[][] = [];
// 		for (let i = 0; i < csv.length; i++) {
// 			for (let j = 0; j < csv[0].length; j++) {
// 				if (columns.includes(csv[0][j])) {
// 					if (!Array.isArray(res[i])) res[i] = [];
// 					res[i].push(csv[i][j]);
// 				}
// 			}
// 		}

// 		const currentHeader = res[0];

// 		// 2. Create a map of the target indices
// 		// This tells us: "Column 0 should actually be at index X"
// 		const sortedIndices = currentHeader
// 			.map((colName, index) => index)
// 			.sort((a, b) => {
// 				return defValue.indexOf(currentHeader[a]) - defValue.indexOf(currentHeader[b]);
// 			});

// 		// 3. Map every row to the new index order
// 		const sortedRes = res.map((row) => sortedIndices.map((i) => row[i]));

// 		return sortedRes;
// 	}
