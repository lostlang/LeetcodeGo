package leetcode

import (
	"reflect"
	"testing"
)

type rotateTestPair struct {
	inputArr []int
	inputK   int
	out      []int
}

var rotateTestCases = []rotateTestPair{
	{[]int{1, 2, 3, 4, 5, 6, 7}, 3., []int{5, 6, 7, 1, 2, 3, 4}},
	{[]int{-1, -100, 3, 99}, 2., []int{3, 99, -1, -100}},
}

func TestEvalRotate(t *testing.T) {
	for _, pair := range rotateTestCases {
		rotate(pair.inputArr, pair.inputK)
		if !reflect.DeepEqual(pair.inputArr, pair.out) {
			t.Error(
				"For", pair.inputK,
				"expected", pair.out,
				"got", pair.inputArr,
			)
		}
	}
}
