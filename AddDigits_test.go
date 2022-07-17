package leetcode

import (
	"reflect"
	"testing"
)

type addDigitsTestPair struct {
	input int
	out   int
}

var addDigitsTestCases = []addDigitsTestPair{
	{38, 2},
	{0, 0},
}

func TestEvalAddDigits(t *testing.T) {
	for _, pair := range addDigitsTestCases {
		var newOut = addDigits(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
