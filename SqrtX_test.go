package leetcode

import (
	"reflect"
	"testing"
)

type mySqrtTestPair struct {
	input int
	out   int
}

var mySqrtTestCases = []mySqrtTestPair{
	{4, 2},
	{8, 2},
}

func TestEvalMySqrt(t *testing.T) {
	for _, pair := range mySqrtTestCases {
		var newOut = mySqrt(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
