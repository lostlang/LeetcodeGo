package leetcode

import (
	"reflect"
	"testing"
)

type arrayStringsAreEqualTestPair struct {
	inputWord1 []string
	inputWord2 []string
	out        bool
}

var arrayStringsAreEqualTestCases = []arrayStringsAreEqualTestPair{
	{[]string{"ab", "c"}, []string{"a", "bc"}, true},
	{[]string{"a", "cb"}, []string{"ab", "c"}, false},
	{[]string{"abc", "d", "defg"}, []string{"abcddefg"}, true},
}

func TestEvalArrayStringsAreEqual(t *testing.T) {
	for _, pair := range arrayStringsAreEqualTestCases {
		newOut := arrayStringsAreEqual(pair.inputWord1, pair.inputWord2)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputWord1, pair.inputWord2,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
