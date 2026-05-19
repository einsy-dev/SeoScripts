package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func ToString(v any) *string {
	if v == nil {
		return nil
	}

	var res string

	switch t := v.(type) {
	case string:
		res = t
	case json.Number:
		res = t.String()
	case float64:
		res = strconv.FormatFloat(t, 'f', -1, 64)
	case bool:
		res = strconv.FormatBool(t)
	case int, int64:
		res = fmt.Sprintf("%d", t)
	default:
		res = fmt.Sprintf("%v", v)
	}

	if res == "" || res == "<nil>" {
		return nil
	}

	res = strings.TrimSpace(res)

	return &res
}
