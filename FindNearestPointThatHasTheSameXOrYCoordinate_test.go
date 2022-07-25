package leetcode

import (
	"reflect"
	"testing"
)

type nearestValidPointTestPair struct {
	inputX     int
	inputY     int
	inputArray [][]int
	out        int
}

var nearestValidPointTestCases = []nearestValidPointTestPair{
	{3, 4, [][]int{[]int{1, 2}, []int{3, 1}, []int{2, 4}, []int{2, 3}, []int{4, 4}}, 2},
	{3, 4, [][]int{[]int{3, 4}}, 0},
	{3, 4, [][]int{[]int{2, 3}}, -1},
}

func TestEvalNearestValidPoint(t *testing.T) {
	for _, pair := range nearestValidPointTestCases {
		var newOut = nearestValidPoint(pair.inputX, pair.inputY, pair.inputArray)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputX,
				",", pair.inputY,
				"and", pair.inputArray,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
