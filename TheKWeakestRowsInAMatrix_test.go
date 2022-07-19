package leetcode

import (
	"reflect"
	"testing"
)

type kWeakestRowsTestPair struct {
	inputArr [][]int
	inputK   int
	out      []int
}

var kWeakestRowsTestCases = []kWeakestRowsTestPair{
	{[][]int{[]int{1, 1, 0, 0, 0}, []int{1, 1, 1, 1, 0}, []int{1, 0, 0, 0, 0}, []int{1, 1, 0, 0, 0}, []int{1, 1, 1, 1, 1}}, 3, []int{2, 0, 3}},
	{[][]int{[]int{1, 0, 0, 0}, []int{1, 1, 1, 1}, []int{1, 0, 0, 0}, []int{1, 0, 0, 0}}, 2, []int{0, 2}},
}

func TestEvalKWeakestRows(t *testing.T) {
	for _, pair := range kWeakestRowsTestCases {
		var newOut = kWeakestRows(pair.inputArr, pair.inputK)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputArr,
				"and", pair.inputK,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
