package leetcode

import (
	"reflect"
	"testing"
)

type arrangeCoinsTestPair struct {
	input int
	out   int
}

var arrangeCoinsTestCases = []arrangeCoinsTestPair{
	{1, 1},
	{2, 1},
	{5, 2},
	{8, 3},
}

func TestEvalArrangeCoins(t *testing.T) {
	for _, pair := range arrangeCoinsTestCases {
		var newOut = arrangeCoins(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
