package leetcode

import (
	"reflect"
	"testing"
)

type findPoisonedDurationTestPair struct {
	inputTime     []int
	inpusDuration int
	out           int
}

var findPoisonedDurationTestCases = []findPoisonedDurationTestPair{
	{[]int{1, 4}, 2, 4},
	{[]int{1, 2}, 2, 3},
}

func TestEvalFindPoisonedDuration(t *testing.T) {
	for _, pair := range findPoisonedDurationTestCases {
		var newOut = findPoisonedDuration(pair.inputTime, pair.inpusDuration)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputTime,
				"and", pair.inpusDuration,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
