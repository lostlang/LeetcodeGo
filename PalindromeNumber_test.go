package leetcode

import (
	"reflect"
	"testing"
)

type isPalindromeIntTestPair struct {
	input int
	out   bool
}

var isPalindromeIntTestCases = []isPalindromeIntTestPair{
	{121, true},
	{-121, false},
	{10, false},
}

func TestEvalIsPalindromeInt(t *testing.T) {
	for _, pair := range isPalindromeIntTestCases {
		var newOut = isPalindromeInt(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
