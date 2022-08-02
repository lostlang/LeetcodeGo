package leetcode

import (
	"reflect"
	"testing"
)

type kthSmallestTestPair struct {
	inputMatrix [][]int
	inputK      int
	out         int
}

var kthSmallestTestCases = []kthSmallestTestPair{
	{[][]int{[]int{1, 5, 9}, []int{10, 11, 13}, []int{12, 13, 15}}, 8, 13},
	{[][]int{[]int{-5}}, 1, -5},
}

func TestEvalKthSmallest(t *testing.T) {
	for _, pair := range kthSmallestTestCases {
		var newOut = kthSmallest(pair.inputMatrix, pair.inputK)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputMatrix,
				"and", pair.inputK,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
