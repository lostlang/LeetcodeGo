package leetcode

import (
	"reflect"
	"testing"
)

type longestContinuousSubstringTestPair struct {
	input string
	out   int
}

var longestContinuousSubstringTestCases = []longestContinuousSubstringTestPair{
	{"abacaba", 2},
	{"abcde", 5},
}

func TestEvalLongestContinuousSubstring(t *testing.T) {
	for _, pair := range longestContinuousSubstringTestCases {
		var newOut = longestContinuousSubstring(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
