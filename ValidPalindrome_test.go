package leetcode

import (
	"reflect"
	"testing"
)

type isPalindromeStrTestPair struct {
	input string
	out   bool
}

var isPalindromeStrTestCases = []isPalindromeStrTestPair{
	{"A man, a plan, a canal: Panama", true},
	{"race a car", false},
	{" ", true},
}

func TestEvalIsPalindromeStr(t *testing.T) {
	for _, pair := range isPalindromeStrTestCases {
		var newOut = isPalindromeStr(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
