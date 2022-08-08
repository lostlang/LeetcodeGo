package leetcode

import (
	"reflect"
	"testing"
)

type dailyTemperaturesTestPair struct {
	input []int
	out   []int
}

var dailyTemperaturesTestCases = []dailyTemperaturesTestPair{
	{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
	{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
}

func TestEvalDailyTemperatures(t *testing.T) {
	for _, pair := range dailyTemperaturesTestCases {
		var newOut = dailyTemperatures(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
