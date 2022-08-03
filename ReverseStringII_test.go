package leetcode

import (
	"reflect"
	"testing"
)

type reverseStrTestPair struct {
	inputS string
	inputK int
	out    string
}

var reverseStrTestCases = []reverseStrTestPair{
	{"abcdefg", 2, "bacdfeg"},
	{"abcd", 2, "bacd"},
}

func TestEvalReverseStr(t *testing.T) {
	for _, pair := range reverseStrTestCases {
		var newOut = reverseStr(pair.inputS, pair.inputK)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputS,
				"and", pair.inputK,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
