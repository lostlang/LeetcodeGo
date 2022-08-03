package leetcode

import (
	"reflect"
	"testing"
)

type decodeTestPair struct {
	inputArr   []int
	inputStart int
	out        []int
}

var decodeTestCases = []decodeTestPair{
	{[]int{1, 2, 3}, 1, []int{1, 0, 2, 1}},
	{[]int{6, 2, 7, 3}, 4, []int{4, 2, 0, 7, 4}},
}

func TestEvalDecode(t *testing.T) {
	for _, pair := range decodeTestCases {
		var newOut = decode(pair.inputArr, pair.inputStart)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputStart,
				"and", pair.inputArr,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
