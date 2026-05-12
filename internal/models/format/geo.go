package format

import "regexp"

var rGeo = regexp.MustCompile(`[a-zA-Z]+`)

func FormatGeo(geo any) any {
	var res = rGeo.FindString(geo.(string))
	if res == "" {
		return nil
	}
	return res
}
