package leetcode

import (
	"reflect"
	"testing"
)

type totalMoneyTestPair struct {
	input int
	out   int
}

var totalMoneyTestCases = []totalMoneyTestPair{
	{4, 10},
	{10, 37},
	{20, 96},
}

func TestEvalTotalMoney(t *testing.T) {
	for _, pair := range totalMoneyTestCases {
		var newOut = totalMoney(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
