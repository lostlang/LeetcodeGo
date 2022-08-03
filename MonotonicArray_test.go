package leetcode

import (
	"reflect"
	"testing"
)

type isMonotonicTestPair struct {
	input []int
	out   bool
}

var isMonotonicTestCases = []isMonotonicTestPair{
	{[]int{1, 2, 2, 3}, true},
	{[]int{6, 5, 4, 4}, true},
	{[]int{1, 3, 2}, false},
}

func TestEvalIsMonotonic(t *testing.T) {
	for _, pair := range isMonotonicTestCases {
		var newOut = isMonotonic(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
