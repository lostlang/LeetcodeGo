package leetcode

import (
	"reflect"
	"testing"
)

type canMakeArithmeticProgressionTestPair struct {
	input []int
	out   bool
}

var canMakeArithmeticProgressionTestCases = []canMakeArithmeticProgressionTestPair{
	{[]int{3, 5, 1}, true},
	{[]int{1, 2, 4}, false},
}

func TestEvalCanMakeArithmeticProgression(t *testing.T) {
	for _, pair := range canMakeArithmeticProgressionTestCases {
		var newOut = canMakeArithmeticProgression(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
