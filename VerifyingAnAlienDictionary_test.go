package leetcode

import (
	"reflect"
	"testing"
)

type isAlienSortedTestPair struct {
	inputWords []string
	inputOrder string
	out        bool
}

var isAlienSortedTestCases = []isAlienSortedTestPair{
	// {[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true},
	// {[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz", false},
	// {[]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz", false},
}

func TestEvalIsAlienSorted(t *testing.T) {
	for _, pair := range isAlienSortedTestCases {
		var newOut = isAlienSorted(pair.inputWords, pair.inputOrder)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputWords,
				"and", pair.inputOrder,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
