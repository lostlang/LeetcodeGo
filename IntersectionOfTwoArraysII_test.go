package leetcode

import (
	"reflect"
	"testing"
)

type intersectTestPair struct {
	inputArr1 []int
	inputArr2 []int
	out       []int
}

var intersectTestCases = []intersectTestPair{
	{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
	{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
}

func TestEvalIntersect(t *testing.T) {
	for _, pair := range intersectTestCases {
		var newOut = intersect(pair.inputArr1, pair.inputArr2)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputArr1,
				"and", pair.inputArr2,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
