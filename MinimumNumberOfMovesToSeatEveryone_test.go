package leetcode

import (
	"reflect"
	"testing"
)

type minMovesToSeatTestPair struct {
	inputSeats    []int
	inputStudents []int
	out           int
}

var minMovesToSeatTestCases = []minMovesToSeatTestPair{
	{[]int{3, 1, 5}, []int{2, 7, 4}, 4},
	{[]int{4, 1, 5, 9}, []int{1, 3, 2, 6}, 7},
	{[]int{2, 2, 6, 6}, []int{1, 3, 2, 6}, 4},
}

func TestEvalMinMovesToSeat(t *testing.T) {
	for _, pair := range minMovesToSeatTestCases {
		var newOut = minMovesToSeat(pair.inputSeats, pair.inputStudents)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputSeats,
				"and", pair.inputStudents,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
