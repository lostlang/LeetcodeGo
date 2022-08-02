package leetcode

import (
	"reflect"
	"testing"
)

type containsNearbyDuplicateTestPair struct {
	inputArr []int
	inputK   int
	out      bool
}

var containsNearbyDuplicateTestCases = []containsNearbyDuplicateTestPair{
	{[]int{1, 2, 3, 1}, 3, true},
	{[]int{1, 0, 1, 1}, 1, true},
	{[]int{1, 2, 3, 1, 2, 3}, 2, false},
}

func TestEvalContainsNearbyDuplicate(t *testing.T) {
	for _, pair := range containsNearbyDuplicateTestCases {
		var newOut = containsNearbyDuplicate(pair.inputArr, pair.inputK)
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
