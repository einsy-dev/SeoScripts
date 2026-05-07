package utils

func ForEach(s []string, fn func(el string) string) {
	for i, v := range s {
		s[i] = fn(v)
	}
}
