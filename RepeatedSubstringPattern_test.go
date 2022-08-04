package leetcode

import (
	"reflect"
	"testing"
)

type repeatedSubstringPatternTestPair struct {
	input string
	out   bool
}

var repeatedSubstringPatternTestCases = []repeatedSubstringPatternTestPair{
	{"abab", true},
	{"aba", false},
	{"abcabcabcabc", true},
}

func TestEvalRepeatedSubstringPattern(t *testing.T) {
	for _, pair := range repeatedSubstringPatternTestCases {
		var newOut = repeatedSubstringPattern(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
