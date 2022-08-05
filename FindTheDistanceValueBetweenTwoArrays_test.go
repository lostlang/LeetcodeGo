package leetcode

import (
	"reflect"
	"testing"
)

type findTheDistanceValueTestPair struct {
	inputArr1     []int
	inputArr2     []int
	inputDistanse int
	out           int
}

var findTheDistanceValueTestCases = []findTheDistanceValueTestPair{
	{[]int{4, 5, 8}, []int{10, 9, 1, 8}, 2, 2},
	{[]int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3, 2},
	{[]int{2, 1, 100, 3}, []int{-5, -2, 10, -3, 7}, 6, 1},
}

func TestEvalFindTheDistanceValue(t *testing.T) {
	for _, pair := range findTheDistanceValueTestCases {
		var newOut = findTheDistanceValue(pair.inputArr1, pair.inputArr2, pair.inputDistanse)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputArr1,
				",", pair.inputArr2,
				"and", pair.inputDistanse,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
