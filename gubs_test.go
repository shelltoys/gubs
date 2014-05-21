package gubs

import "testing"

var testTable = []struct {
	before string
	after  string
}{
	{"hello bubs", "ⓗ ⓔ ⓛ ⓛ ⓞ ⓑ ⓤ ⓑ ⓢ "},
	{"Zach Holman", "Ⓩ ⓐ ⓒ ⓗ Ⓗ ⓞ ⓛ ⓜ ⓐ ⓝ "},
}

func TestBlubsTranslations(t *testing.T) {
	for _, test := range testTable {
		actual := Tr(test.before)
		if actual != test.after {
			t.Errorf("[%s] result %s does not match expectation %s", test.before, actual, test.after)
		}
	}
}
