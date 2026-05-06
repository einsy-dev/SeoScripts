package utils

import "strings"

func FlatCsv(h [][]any) []any {
	var newH = make([]any, len(h[0]))
	var acc = make([]any, len(h))

	for j := range h[0] { // col
		for i := range h { // row
			if h[i][j] != "" {
				acc[i] = h[i][j]
			}
		}

		// filter tempAcc
		var tAcc = []string{}
		for _, v := range acc {
			if v != "" && v != nil {
				tAcc = append(tAcc, v.(string))
			}
		}
		newH[j] = strings.Join(tAcc, ".")
	}
	return newH
}
