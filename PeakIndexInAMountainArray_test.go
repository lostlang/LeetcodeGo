package leetcode

import (
	"reflect"
	"testing"
)

type peakIndexInMountainArrayTestPair struct {
	input []int
	out   int
}

var peakIndexInMountainArrayTestCases = []peakIndexInMountainArrayTestPair{
	{[]int{0, 1, 0}, 1},
	{[]int{0, 2, 1, 0}, 1},
	{[]int{0, 10, 5, 2}, 1},
}

func TestEvalPeakIndexInMountainArray(t *testing.T) {
	for _, pair := range peakIndexInMountainArrayTestCases {
		var newOut = peakIndexInMountainArray(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
