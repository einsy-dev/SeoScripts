package utils

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func ToCamelCase(str string) string {
	if str == "" {
		return str
	}
	if len(str) == 2 {
		return strings.ToLower(str)
	}

	// Decode the first rune in the string
	r, size := utf8.DecodeRuneInString(str)

	// Check for a decoding error and ensure size is valid
	if r == utf8.RuneError && size <= 1 {
		return str
	}

	lc := unicode.ToLower(r)

	// If the rune is already lowercase, return the original string for efficiency
	if r == lc {
		return str
	}

	// Reconstruct the string with the lowercase first rune and the rest of the original string
	return string(lc) + str[size:]

}
