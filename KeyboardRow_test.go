package leetcode

import (
	"reflect"
	"testing"
)

type findWordsTestPair struct {
	input []string
	out   []string
}

var findWordsTestCases = []findWordsTestPair{
	{[]string{"Hello", "Alaska", "Dad", "Peace"}, []string{"Alaska", "Dad"}},
	{[]string{"omk"}, []string{}},
	{[]string{"adsdf", "sfd"}, []string{"adsdf", "sfd"}},
}

func TestEvalFindWords(t *testing.T) {
	for _, pair := range findWordsTestCases {
		var newOut = findWords(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
