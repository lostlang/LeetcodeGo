package leetcode

import (
	"reflect"
	"testing"
)

type countVowelPermutationTestPair struct {
	input int
	out   int
}

var countVowelPermutationTestCases = []countVowelPermutationTestPair{
	{1, 5},
	{2, 10},
	{5, 68},
	{144, 18208803},
}

func TestEvalCountVowelPermutation(t *testing.T) {
	for _, pair := range countVowelPermutationTestCases {
		var newOut = countVowelPermutation(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
