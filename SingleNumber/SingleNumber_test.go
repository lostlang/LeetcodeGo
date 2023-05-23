package SingleNumber

import (
	"reflect"
	"testing"
)

type singleNumberTestPair struct {
	input []int
	out   int
}

var singleNumberTestCases = []singleNumberTestPair{
	{[]int{2, 2, 1}, 1},
	{[]int{4, 1, 2, 1, 2}, 4},
	{[]int{1}, 1},
}

func TestEvalSingleNumber(t *testing.T) {
	for _, pair := range singleNumberTestCases {
		newOut := singleNumber(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
