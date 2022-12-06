package leetcode

import (
	"reflect"
	"testing"
)

type groupAnagramsTestPair struct {
	input []string
	out   [][]string
}

var groupAnagramsTestCases = []groupAnagramsTestPair{
	//{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
	{[]string{""}, [][]string{{""}}},
	{[]string{"a"}, [][]string{{"a"}}},
}

func TestEvalGroupAnagrams(t *testing.T) {
	for _, pair := range groupAnagramsTestCases {
		var newOut = groupAnagrams(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
