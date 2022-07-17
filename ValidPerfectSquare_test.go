package leetcode

import (
	"reflect"
	"testing"
)

type isPerfectSquareTestPair struct {
	input int
	out   bool
}

var isPerfectSquareTestCases = []isPerfectSquareTestPair{
	{16, true},
	{14, false},
}

func TestEvalIsPerfectSquare(t *testing.T) {
	for _, pair := range isPerfectSquareTestCases {
		var newOut = isPerfectSquare(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
