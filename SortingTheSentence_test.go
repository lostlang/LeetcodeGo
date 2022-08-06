package leetcode

import (
	"reflect"
	"testing"
)

type sortSentenceTestPair struct {
	input string
	out   string
}

var sortSentenceTestCases = []sortSentenceTestPair{
	{"is2 sentence4 This1 a3", "This is a sentence"},
	{"Myself2 Me1 I4 and3", "Me Myself and I"},
}

func TestEvalSortSentence(t *testing.T) {
	for _, pair := range sortSentenceTestCases {
		var newOut = sortSentence(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
