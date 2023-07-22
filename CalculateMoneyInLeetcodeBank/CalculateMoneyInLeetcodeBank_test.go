package CalculateMoneyInLeetcodeBank

import (
	"reflect"
	"testing"
)

type totalMoneyTestPair struct {
	input  int
	output int
}

var totalMoneyTestCases = []totalMoneyTestPair{
	{4, 10},
	{10, 37},
	{20, 96},
}

func TestEvalTotalMoney(t *testing.T) {
	for _, pair := range totalMoneyTestCases {
		newOut := totalMoney(pair.input)
		if !reflect.DeepEqual(newOut, pair.output) {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", newOut,
			)
		}
	}
}
