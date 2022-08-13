package leetcode

import (
	"reflect"
	"testing"
)

type findSubstringTestPair struct {
	inputS     string
	inputWords []string
	out        []int
}

var findSubstringTestCases = []findSubstringTestPair{
	{"barfoothefoobarman", []string{"foo", "bar"}, []int{0, 9}},
	{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}, []int{}},
	{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}, []int{6, 9, 12}},
	{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}, []int{8}},
}

func TestEvalFindSubstring(t *testing.T) {
	for _, pair := range findSubstringTestCases {
		var newOut = findSubstring(pair.inputS, pair.inputWords)
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
