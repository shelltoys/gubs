package gubs

import (
	"regexp"
	"strings"
)

var matcher = regexp.MustCompile(`A-Za-z1-90`)

func Tr(input string) string {
	results := ""

	lowers := "abcdefghijklmnopqrstuvwxyz"
	uppers := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	positives := "123456789"
	zero := "0"

	upperOffset := rune('Ⓐ') - rune('A')
	lowerOffset := rune('ⓐ') - rune('a')
	posOffset := rune('①') - rune('1')
	zeroOffset := rune('⓪') - rune('0')

	offsetMap := map[string]rune{lowers: lowerOffset, uppers: upperOffset, positives: posOffset, zero: zeroOffset}

	for _, c := range input {
		for key, offset := range offsetMap {
			if strings.Contains(key, string(c)) {
				results += string(c + offset)
				results += " "
			}
		}
	}

	return results
}
