package leetcode

import (
	"reflect"
	"testing"
)

type majorityElementTestPair struct {
	input []int
	out   int
}

var majorityElementTestCases = []majorityElementTestPair{
	{[]int{3, 2, 3}, 3},
	{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
}

func TestEvalMajorityElement(t *testing.T) {
	for _, pair := range majorityElementTestCases {
		var newOut = majorityElement(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
