package CanMakeArithmeticProgressionFromSequence

import (
	"reflect"
	"testing"
)

type canMakeArithmeticProgressionTestPair struct {
	input  []int
	output bool
}

var canMakeArithmeticProgressionTestCases = []canMakeArithmeticProgressionTestPair{
	{[]int{3, 5, 1}, true},
	{[]int{1, 2, 4}, false},
}

func TestEvalCanMakeArithmeticProgression(t *testing.T) {
	for _, pair := range canMakeArithmeticProgressionTestCases {
		newOut := canMakeArithmeticProgression(pair.input)
		if !reflect.DeepEqual(newOut, pair.output) {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", newOut,
			)
		}
	}
}
