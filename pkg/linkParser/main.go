package linkParser

import (
	"regexp"
	"slices"
	"strings"
)

var dom = regexp.MustCompile(`^(?:https?:\/\/)?(?:www\.)?([^\/\s]+\.[^\/\s?]+)`)
var pathname = regexp.MustCompile(``)

func Domain(url string) string {
	matches := dom.FindStringSubmatch(url)
	if len(matches) > 1 {
		var g = matches[1]

		dIndex := slices.IndexFunc(replace, func(el string) bool {
			return strings.Contains(g, el)
		})

		if dIndex == -1 {
			return g
		}

		return replace[dIndex]
	}
	return ""
} // Domain

func RootDomain(url string) string { return "" } // SLD | TLD

func PathName(url string) string { return "" } // "/...?"

func Params(url string) string { return "" } // "?..."
