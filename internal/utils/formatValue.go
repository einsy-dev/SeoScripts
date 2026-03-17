package utils

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var reTh = regexp.MustCompile(`[kKлЛ]`)
var reMil = regexp.MustCompile(`[bBиИ]`)
var reBil = regexp.MustCompile(`[mMьЬ]`)
var rePoint = regexp.MustCompile(`\,`)
var reInt = regexp.MustCompile(`[0-9\.]*`)
var reTrue = regexp.MustCompile(`(TRUE)|(true)`)
var reFalse = regexp.MustCompile(`(FALSE)|(false)`)

func formatValue(k any) any {
	if k == nil || k == "" {
		return k
	}

	reInt := regexp.MustCompile(`^[0-9\.\,]+[kKлЛbBиИmMьЬ]$`)
	val, err := json.Marshal(k)
	if err != nil {
		log.Fatal("Error marshal ", k)
	}

	if reInt.Match(val) {
		var mult = 1

		// replace letters with multiplier
		if reTh.Match(val) {
			mult *= 1e3
			reTh.ReplaceAll(val, []byte(""))
		} else if reMil.Match(val) {
			mult *= 1e6
			reMil.ReplaceAll(val, []byte(""))
		} else if reBil.Match(val) {
			mult *= 1e9
			reBil.ReplaceAll(val, []byte(""))
		}
		// replace comma with point
		val = rePoint.ReplaceAll(val, []byte("."))

		// select only numbers
		num := reInt.Find(val)
		numI, err := strconv.Atoi(string(num))
		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}
		return numI * mult
	} else {
		if m, ok := k.(string); ok {
			return strings.TrimSpace(m)
		}
		return k
	}
}
