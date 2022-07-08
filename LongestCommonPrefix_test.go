package leetcode

import (
	"reflect"
	"testing"
)

type longestCommonPrefixTestPair struct {
	input []string
	out   string
}

var longestCommonPrefixTestCases = []longestCommonPrefixTestPair{
	{[]string{"flower", "flow", "flight"}, "fl"},
	{[]string{"dog", "racecar", "car"}, ""},
	{[]string{"", "b"}, ""},
}

func TestEvalLongestCommonPrefix(t *testing.T) {
	for _, pair := range longestCommonPrefixTestCases {
		var newOut = longestCommonPrefix(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
