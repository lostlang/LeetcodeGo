package leetcode

import (
	"reflect"
	"testing"
)

type subtractProductAndSumTestPair struct {
	input int
	out   int
}

var subtractProductAndSumTestCases = []subtractProductAndSumTestPair{
	{234, 15},
	{4421, 21},
}

func TestEvalSubtractProductAndSum(t *testing.T) {
	for _, pair := range subtractProductAndSumTestCases {
		var newOut = subtractProductAndSum(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
