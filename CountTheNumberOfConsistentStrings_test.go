package leetcode

import (
	"reflect"
	"testing"
)

type countConsistentStringsTestPair struct {
	inputAllowed string
	inputWords   []string
	out          int
}

var countConsistentStringsTestCases = []countConsistentStringsTestPair{
	{"ab", []string{"ad", "bd", "aaab", "baa", "badab"}, 2},
	{"abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}, 7},
}

func TestEvalCountConsistentStrings(t *testing.T) {
	for _, pair := range countConsistentStringsTestCases {
		newOut := countConsistentStrings(pair.inputAllowed, pair.inputWords)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputAllowed, pair.inputWords,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
