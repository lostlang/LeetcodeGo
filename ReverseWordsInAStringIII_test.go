package leetcode

import (
	"reflect"
	"testing"
)

type reverseWordsTestPair struct {
	input string
	out   string
}

var reverseWordsTestCases = []reverseWordsTestPair{
	{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
	{"God Ding", "doG gniD"},
}

func TestEvalReverseWords(t *testing.T) {
	for _, pair := range reverseWordsTestCases {
		var newOut = reverseWords(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
