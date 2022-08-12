package leetcode

import (
	"reflect"
	"testing"
)

type numIslandsTestPair struct {
	input [][]byte
	out   int
}

var numIslandsTestCases = []numIslandsTestPair{
	{[][]byte{[]byte{1, 1, 0, 0, 0}, []byte{1, 1, 0, 0, 0}, []byte{0, 0, 1, 0, 0}, []byte{0, 0, 0, 1, 1}}, 3},
	{[][]byte{[]byte{1, 1, 1, 1, 0}, []byte{1, 1, 0, 1, 0}, []byte{1, 1, 0, 0, 0}, []byte{0, 0, 0, 0, 0}}, 1},
}

func TestEvalNumIslands(t *testing.T) {
	for _, pair := range numIslandsTestCases {
		var newOut = numIslands(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
