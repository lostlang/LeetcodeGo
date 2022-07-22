package leetcode

import (
	"reflect"
	"testing"
)

type longestPalindromeTestPair struct {
	input string
	out   int
}

var longestPalindromeTestCases = []longestPalindromeTestPair{
	{"abccccdd", 7},
	{"a", 1},
}

func TestEvalLongestPalindrome(t *testing.T) {
	for _, pair := range longestPalindromeTestCases {
		var newOut = longestPalindrome(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
