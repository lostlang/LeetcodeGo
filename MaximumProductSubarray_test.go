package leetcode

import (
	"reflect"
	"testing"
)

type maxProductTestPair struct {
	input []int
	out   int
}

var maxProductTestCases = []maxProductTestPair{
	{[]int{2, 3, -2, 4}, 6},
	{[]int{-2, 0, -1}, 0},
	{[]int{-2, 3, -4}, 24},
	{[]int{0, 0, 0}, 0},
	{[]int{0}, 0},
	{[]int{-3, -1, -1}, 3},
}

func TestEvalMaxProduct(t *testing.T) {
	for _, pair := range maxProductTestCases {
		var newOut = maxProduct(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
