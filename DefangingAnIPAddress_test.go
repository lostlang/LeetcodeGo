package leetcode

import (
	"reflect"
	"testing"
)

type defangIPaddrTestPair struct {
	input string
	out   string
}

var defangIPaddrTestCases = []defangIPaddrTestPair{
	{"1.1.1.1", "1[.]1[.]1[.]1"},
	{"255.100.50.0", "255[.]100[.]50[.]0"},
}

func TestEvalDefangIPaddr(t *testing.T) {
	for _, pair := range defangIPaddrTestCases {
		var newOut = defangIPaddr(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
