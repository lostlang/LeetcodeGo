package leetcode

import (
	"reflect"
	"testing"
)

type averageTestPair struct {
	input []int
	out   float64
}

var averageTestCases = []averageTestPair{
	{[]int{4000, 3000, 1000, 2000}, 2500.0},
	{[]int{1000, 2000, 3000}, 2000.0},
}

func TestEvalAverage(t *testing.T) {
	for _, pair := range averageTestCases {
		var newOut = average(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
