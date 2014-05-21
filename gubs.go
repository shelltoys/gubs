// Package gubs translates alphanumeric characters into bubbly versions of themselves.
package gubs

import "strings"

var offsetMap = map[string]rune{
	"abcdefghijklmnopqrstuvwxyz": (rune('ⓐ') - rune('a')),
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ": (rune('Ⓐ') - rune('A')),
	"123456789":                  (rune('①') - rune('1')),
	"0":                          (rune('⓪') - rune('0')),
}

// Tr translates a string into a bubbly string.
func Tr(input string) string {
	results := ""

	for _, c := range input {
		result := string(c)
		for key, offset := range offsetMap {
			if strings.Contains(key, string(c)) {
				result = string(c+offset) + " "
				break
			}
		}
		results += result
	}

	return strings.TrimSpace(results)
}
