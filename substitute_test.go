package grename

import (
	"testing"
)

type RenamerTest struct {
	matchre string
	subst   string
	input   string
	result  string
}

var RenamerTests = []RenamerTest{
	RenamerTest{`\.(png|jpg|gif)$`, `.not_a.$1`, `test.gif`, `test.not_a.gif`},
	RenamerTest{`_unwanted_ext$`, ``, `test_unwanted_ext`, `test`},
}

func TestMakeRE2Renamer(t *testing.T) {
	for _, test := range RenamerTests {
		r, _ := MakeRE2Renamer(test.matchre, test.subst)
		result := r(test.input)
		if result != test.result {
			t.Errorf("Renamer('%s', '%s') for string '%s' return '%s', expected '%s'\n", test.matchre, test.subst, test.input, result, test.result)
		}
	}
}
