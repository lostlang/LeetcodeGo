package leetcode

import (
	"reflect"
	"testing"
)

type sortByBitsTestPair struct {
	input []int
	out   []int
}

var sortByBitsTestCases = []sortByBitsTestPair{
	{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, []int{0, 1, 2, 4, 8, 3, 5, 6, 7}},
	{[]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}, []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}},
}

func TestEvalSortByBits(t *testing.T) {
	for _, pair := range sortByBitsTestCases {
		var newOut = sortByBits(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
