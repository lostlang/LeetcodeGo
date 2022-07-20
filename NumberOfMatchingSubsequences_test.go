package leetcode

import (
	"reflect"
	"testing"
)

type numMatchingSubseqTestPair struct {
	inputS     string
	inputWords []string
	out        int
}

var numMatchingSubseqTestCases = []numMatchingSubseqTestPair{
	{"abcde", []string{"a", "bb", "acd", "ace"}, 3},
	{"dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}, 2},
}

func TestEvalNumMatchingSubseq(t *testing.T) {
	for _, pair := range numMatchingSubseqTestCases {
		var newOut = numMatchingSubseq(pair.inputS, pair.inputWords)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputS,
				"and", pair.inputWords,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
