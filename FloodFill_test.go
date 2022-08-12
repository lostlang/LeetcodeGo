package leetcode

import (
	"reflect"
	"testing"
)

type floodFillTestPair struct {
	inputArray [][]int
	inputSR    int
	inputSC    int
	inputColor int
	out        [][]int
}

var floodFillTestCases = []floodFillTestPair{
	{[][]int{[]int{1, 1, 1}, []int{1, 1, 0}, []int{1, 0, 1}}, 1, 1, 2, [][]int{[]int{2, 2, 2}, []int{2, 2, 0}, []int{2, 0, 1}}},
}

func TestEvalFloodFill(t *testing.T) {
	for _, pair := range floodFillTestCases {
		var newOut = floodFill(pair.inputArray, pair.inputSR, pair.inputSC, pair.inputColor)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputArray,
				",", pair.inputSR,
				",", pair.inputSC,
				"and", pair.inputColor,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
