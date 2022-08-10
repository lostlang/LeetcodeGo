package leetcode

import (
	"reflect"
	"testing"
)

type findRotationTestPair struct {
	inputMatrix [][]int
	inputTarget [][]int
	out         bool
}

var findRotationTestCases = []findRotationTestPair{
	{[][]int{[]int{0, 1}, []int{1, 0}}, [][]int{[]int{1, 0}, []int{0, 1}}, true},
	{[][]int{[]int{1, 1}, []int{0, 1}}, [][]int{[]int{1, 1}, []int{1, 0}}, true},
	{[][]int{[]int{0, 1}, []int{1, 1}}, [][]int{[]int{1, 0}, []int{0, 1}}, false},
	{[][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{1, 1, 1}}, [][]int{[]int{1, 1, 1}, []int{0, 1, 0}, []int{0, 0, 0}}, true},
}

func TestEvalFindRotation(t *testing.T) {
	for _, pair := range findRotationTestCases {
		var newOut = findRotation(pair.inputMatrix, pair.inputTarget)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputMatrix,
				"and", pair.inputTarget,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
