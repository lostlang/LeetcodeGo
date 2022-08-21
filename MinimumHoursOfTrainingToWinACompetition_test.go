package leetcode

import (
	"reflect"
	"testing"
)

type minNumberOfHoursTestPair struct {
	inputInitialEnergy     int
	inputInitialExperience int
	inputEnergy            []int
	inputExperience        []int
	out                    int
}

var minNumberOfHoursTestCases = []minNumberOfHoursTestPair{
	{5, 3, []int{1, 4, 3, 2}, []int{2, 6, 3, 1}, 8},
	{2, 4, []int{1}, []int{3}, 0},
	{5, 3, []int{1, 4}, []int{2, 5}, 2},
	{1, 1, []int{1, 1, 1, 1}, []int{1, 1, 1, 50}, 51},
	{11, 23,
		[]int{69, 22, 93, 68, 57, 76, 54, 72, 8, 78, 88, 15, 58, 61, 25, 70, 58, 91, 79, 22,
			91, 74, 90, 75, 31, 53, 31, 7, 51, 96, 76, 17, 64, 89, 83, 54, 28, 33, 32, 45, 20},
		[]int{51, 81, 46, 80, 56, 7, 46, 74, 64, 20, 84, 66, 13, 91, 49, 30, 75, 43, 74, 88,
			82, 51, 72, 4, 80, 30, 10, 19, 40, 27, 21, 71, 24, 94, 79, 13, 40, 28, 63, 85, 70},
		2323},
}

func TestEvalMinNumberOfHours(t *testing.T) {
	for _, pair := range minNumberOfHoursTestCases {
		var newOut = minNumberOfHours(pair.inputInitialEnergy, pair.inputInitialExperience, pair.inputEnergy, pair.inputExperience)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputInitialEnergy,
				",", pair.inputInitialExperience,
				",", pair.inputEnergy,
				"and", pair.inputExperience,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
