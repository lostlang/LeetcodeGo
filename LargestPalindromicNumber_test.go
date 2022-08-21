package leetcode

import (
	"reflect"
	"testing"
)

type largestPalindromicTestPair struct {
	input string
	out   string
}

var largestPalindromicTestCases = []largestPalindromicTestPair{
	{"444947137", "7449447"},
	{"00009", "9"},
	{"00001105", "1005001"},
}

func TestEvalLargestPalindromic(t *testing.T) {
	for _, pair := range largestPalindromicTestCases {
		var newOut = largestPalindromic(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
