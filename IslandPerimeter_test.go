package leetcode

import (
	"reflect"
	"testing"
)

type islandPerimeterTestPair struct {
	input [][]int
	out   int
}

var islandPerimeterTestCases = []islandPerimeterTestPair{
	{[][]int{[]int{0, 1, 0, 0}, []int{1, 1, 1, 0}, []int{0, 1, 0, 0}, []int{1, 1, 0, 0}}, 16},
}

func TestEvalIslandPerimeter(t *testing.T) {
	for _, pair := range islandPerimeterTestCases {
		var newOut = islandPerimeter(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
