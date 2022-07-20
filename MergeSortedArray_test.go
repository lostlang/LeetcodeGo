package leetcode

import (
	"reflect"
	"testing"
)

type mergeTestPair struct {
	inputArray1 []int
	inputM      int
	inputArray2 []int
	inputN      int
	out         []int
}

var mergeTestCases = []mergeTestPair{
	{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
	{[]int{1}, 1, []int{}, 0, []int{1}},
	{[]int{0}, 0, []int{1}, 1, []int{1}},
}

func TestEvalMerge(t *testing.T) {
	for _, pair := range mergeTestCases {
		merge(pair.inputArray1, pair.inputM, pair.inputArray2, pair.inputN)
		if !reflect.DeepEqual(pair.inputArray1, pair.out) {
			t.Error(
				"Expected", pair.out,
				"got", pair.inputArray1,
			)
		}
	}
}
