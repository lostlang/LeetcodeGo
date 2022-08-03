package leetcode

import (
	"reflect"
	"testing"
)

type reverseVowelsTestPair struct {
	input string
	out   string
}

var reverseVowelsTestCases = []reverseVowelsTestPair{
	{"hello", "holle"},
	{"leetcode", "leotcede"},
	{"aA", "Aa"},
}

func TestEvalReverseVowels(t *testing.T) {
	for _, pair := range reverseVowelsTestCases {
		var newOut = reverseVowels(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
