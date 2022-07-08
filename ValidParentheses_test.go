package leetcode

import (
	"reflect"
	"testing"
)

type isValidTestPair struct {
	input string
	out   bool
}

var isValidTestCases = []isValidTestPair{
	{"()", true},
	{"()[]{}", true},
	{"(]", false},
	{"[", false},
	{"]", false},
}

func TestEvalIsValid(t *testing.T) {
	for _, pair := range isValidTestCases {
		var newOut = isValid(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
