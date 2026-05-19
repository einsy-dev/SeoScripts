package format

import "strings"

func ToLower(str any) any {
	if str == nil {
		return nil
	}
	var res = strings.ToLower(str.(string))
	if res == "" {
		return nil
	}
	return res
}
