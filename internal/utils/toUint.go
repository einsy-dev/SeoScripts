package utils

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var reInt = regexp.MustCompile(`^[0-9]+(?:[\.\,][0-9]+)?[kKлЛrRкКmMьЬмМvVbBиИбБ]?$`)
var reTh = regexp.MustCompile(`[kKлЛrRкК]`)
var reMil = regexp.MustCompile(`[mMьЬмМvV]`)
var reBil = regexp.MustCompile(`[bBиИбБ]`)
var rePoint = regexp.MustCompile(`\,`)
var reDig = regexp.MustCompile(`\d+(\.\d{0,2})?`)
var reTrue = regexp.MustCompile(`(TRUE)|(true)`)
var reFalse = regexp.MustCompile(`(FALSE)|(false)`)

func ToUint(val *any) *uint {
	var nVal = *val

	if m, ok := nVal.(uint); ok {
		return &m
	} else if m, ok := nVal.(int); ok {
		res := uint(m)
		return &res
	} else if m, ok := nVal.(float64); ok {
		res := uint(math.Round(m))
		return &res
	}

	v := []byte(strings.TrimSpace(fmt.Sprintf("%v", nVal)))

	if reInt.Match(v) {
		var mult = 1
		// replace letters with multiplier
		if reTh.Match(v) {
			mult *= 1e3
			v = reTh.ReplaceAll(v, []byte(""))
		} else if reMil.Match(v) {
			mult *= 1e6
			v = reMil.ReplaceAll(v, []byte(""))
		} else if reBil.Match(v) {
			mult *= 1e9
			v = reBil.ReplaceAll(v, []byte(""))
		}

		// replace comma with point
		v = rePoint.ReplaceAll(v, []byte("."))
		v = reDig.Find(v)

		numI, err := strconv.ParseFloat(string(v), 64)

		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}

		res := uint(math.Round(numI * float64(mult)))

		return &res
	}

	return nil
}
