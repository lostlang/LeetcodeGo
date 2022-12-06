package leetcode

import (
	"reflect"
	"testing"
)

type lengthOfLongestSubstringTestPair struct {
	input string
	out   int
}

var lengthOfLongestSubstringTestCases = []lengthOfLongestSubstringTestPair{
	{"abcabcbb", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
	{" ", 1},
	{"au", 2},
}

func TestEvalLengthOfLongestSubstring(t *testing.T) {
	for _, pair := range lengthOfLongestSubstringTestCases {
		var newOut = lengthOfLongestSubstring(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
