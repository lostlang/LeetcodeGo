package leetcode

import (
	"reflect"
	"testing"
)

type generateTestPair struct {
	input int
	out   [][]int
}

var generateTestCases = []generateTestPair{
	{5, [][]int{[]int{1}, []int{1, 1}, []int{1, 2, 1}, []int{1, 3, 3, 1}, []int{1, 4, 6, 4, 1}}},
	{1, [][]int{[]int{1}}},
}

func TestEvalGenerate(t *testing.T) {
	for _, pair := range generateTestCases {
		var newOut = generate(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
