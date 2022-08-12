package leetcode

import (
	"reflect"
	"testing"
)

type letterCasePermutationTestPair struct {
	input string
	out   []string
}

var letterCasePermutationTestCases = []letterCasePermutationTestPair{
	{"3z4", []string{"3Z4", "3z4"}},
}

func TestEvalLetterCasePermutation(t *testing.T) {
	for _, pair := range letterCasePermutationTestCases {
		var newOut = letterCasePermutation(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
