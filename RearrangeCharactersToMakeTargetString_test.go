package leetcode

import (
	"reflect"
	"testing"
)

type rearrangeCharactersTestPair struct {
	inputString string
	inputTarget string
	out         int
}

var rearrangeCharactersTestCases = []rearrangeCharactersTestPair{
	{"ilovecodingonleetcode", "code", 2},
	{"abcba", "abc", 1},
	{"abbaccaddaeea", "aaaaa", 1},
}

func TestEvalRearrangeCharacters(t *testing.T) {
	for _, pair := range rearrangeCharactersTestCases {
		var newOut = rearrangeCharacters(pair.inputString, pair.inputTarget)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputString,
				"and", pair.inputTarget,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
