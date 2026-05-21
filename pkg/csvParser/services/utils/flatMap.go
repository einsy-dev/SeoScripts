package utils

import (
	"strings"
)

// FlatMap flattens a nested map[string]any into a single-level map
// where nested keys are joined by a dot (".") separator.
func FlatMap(src map[string]any) map[string]any {
	acc := make(map[string]any)
	_flatMap(src, "", acc)
	return acc
}

func _flatMap(src map[string]any, prefix string, acc map[string]any) {
	for k, v := range src {
		nk := k
		if prefix != "" {
			nk = strings.Join([]string{prefix, k}, ".")
		}

		if m, ok := v.(map[string]any); ok {
			_flatMap(m, nk, acc)
		} else {
			acc[nk] = v
		}
	}
}
