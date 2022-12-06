package leetcode

import (
	"reflect"
	"testing"
)

type longestPalindromeSubstringTestPair struct {
	input string
	out   string
}

var longestPalindromeSubstringTestCases = []longestPalindromeSubstringTestPair{
	{"babad", "bab"},
	{"cbbd", "bb"},
}

func TestEvalLongestPalindromeSubstring(t *testing.T) {
	for _, pair := range longestPalindromeSubstringTestCases {
		var newOut = longestPalindromeSub(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
