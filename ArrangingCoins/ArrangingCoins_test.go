package arrangingcoins

import (
	"reflect"
	"testing"
)

type arrangeCoinsTestPair struct {
	input  int
	output int
}

var arrangeCoinsTestCases = []arrangeCoinsTestPair{
	{1, 1},
	{2, 1},
	{5, 2},
	{8, 3},
}

func TestEvalArrangeCoins(t *testing.T) {
	for _, pair := range arrangeCoinsTestCases {
		newOutput := arrangeCoins(pair.input)
		if !reflect.DeepEqual(newOutput, pair.output) {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", newOutput,
			)
		}
	}
}
