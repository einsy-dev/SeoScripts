package linkParser

import (
	"regexp"
)

var dom = regexp.MustCompile(`^(?:https?:\/\/)?(?:www\.)?([^\/\s]+\.[^\/\s?]+)`)
var pathname = regexp.MustCompile(``)

func Domain(url string) string {
	matches := dom.FindStringSubmatch(url)
	if len(matches) > 1 {
		return matches[1]
	}
	return url
} // Domain

func RootDomain(url string) string { return "" } // SLD | TLD

func PathName(url string) string { return "" } // "/...?"

func Params(url string) string { return "" } // "?..."
