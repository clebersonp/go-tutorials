package stringutil

import "unicode"

// ToUpper upper cases all the runes(alias for int32) in its argument string.
func ToUpper(s string) string {
	r := []rune(s)
	for i, char := range r {
		r[i] = unicode.ToUpper(char)
	}
	return string(r)
}
