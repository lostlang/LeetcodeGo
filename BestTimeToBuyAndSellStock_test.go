package leetcode

import (
	"reflect"
	"testing"
)

type maxProfitTestPair struct {
	input []int
	out   int
}

var maxProfitTestCases = []maxProfitTestPair{
	{[]int{7, 1, 5, 3, 6, 4}, 5},
	{[]int{7, 6, 4, 3, 1}, 0},
}

func TestEvalMaxProfit(t *testing.T) {
	for _, pair := range maxProfitTestCases {
		var newOut = maxProfit(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
