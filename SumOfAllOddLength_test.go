package leetcode

import (
	"reflect"
	"testing"
)

type SubarraysTestPair struct {
	input []int
	out   int
}

var SubarraysTestCases = []SubarraysTestPair{
	{[]int{1, 4, 2, 5, 3}, 58},
	{[]int{1, 2}, 3},
	{[]int{10, 11, 12}, 66},
}

func TestEvalSubarrays(t *testing.T) {
	for _, pair := range SubarraysTestCases {
		var newOut = Subarrays(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
