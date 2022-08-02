package leetcode

import (
	"reflect"
	"testing"
)

type wordPatternTestPair struct {
	inputPattern string
	inputString  string
	out          bool
}

var wordPatternTestCases = []wordPatternTestPair{
	{"abba", "dog cat cat dog", true},
	{"abba", "dog cat cat fish", false},
	{"aaaa", "dog cat cat dog", false},
	{"e", "eukera", true},
}

func TestEvalWordPattern(t *testing.T) {
	for _, pair := range wordPatternTestCases {
		var newOut = wordPattern(pair.inputPattern, pair.inputString)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputPattern,
				"and", pair.inputString,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
