package leetcode

import (
	"reflect"
	"testing"
)

type decompressRLElistTestPair struct {
	input []int
	out   []int
}

var decompressRLElistTestCases = []decompressRLElistTestPair{
	{[]int{1, 2, 3, 4}, []int{2, 4, 4, 4}},
	{[]int{1, 1, 2, 3}, []int{1, 3, 3}},
}

func TestEvalDecompressRLElist(t *testing.T) {
	for _, pair := range decompressRLElistTestCases {
		var newOut = decompressRLElist(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
