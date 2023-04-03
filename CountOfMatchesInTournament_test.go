package leetcode

import (
	"reflect"
	"testing"
)

type numberOfMatchesTestPair struct {
	input int
	out   int
}

var numberOfMatchesTestCases = []numberOfMatchesTestPair{
	{7, 6},
	{14, 13},
}

func TestEvalNumberOfMatches(t *testing.T) {
	for _, pair := range numberOfMatchesTestCases {
		newOut := numberOfMatches(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
