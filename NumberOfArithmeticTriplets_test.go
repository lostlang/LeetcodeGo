package leetcode

import (
	"reflect"
	"testing"
)

type arithmeticTripletsTestPair struct {
	inputArr   []int
	intputDiff int
	out        int
}

var arithmeticTripletsTestCases = []arithmeticTripletsTestPair{
	{[]int{0, 1, 4, 6, 7, 10}, 3, 2},
	{[]int{4, 5, 6, 7, 8, 9}, 2, 2},
}

func TestEvalArithmeticTriplets(t *testing.T) {
	for _, pair := range arithmeticTripletsTestCases {
		var newOut = arithmeticTriplets(pair.inputArr, pair.intputDiff)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputArr,
				"and", pair.intputDiff,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
