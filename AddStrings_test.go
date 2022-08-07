package leetcode

import (
	"reflect"
	"testing"
)

type addStringsTestPair struct {
	inputS1 string
	inputS2 string
	out     string
}

var addStringsTestCases = []addStringsTestPair{
	{"11", "123", "134"},
	{"0", "0", "0"},
	{"456", "77", "533"},
}

func TestEvalAddStrings(t *testing.T) {
	for _, pair := range addStringsTestCases {
		var newOut = addStrings(pair.inputS1, pair.inputS2)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputS1,
				"and", pair.inputS2,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
