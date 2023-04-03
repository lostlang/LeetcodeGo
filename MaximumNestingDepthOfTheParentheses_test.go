package leetcode

import (
	"reflect"
	"testing"
)

type maxDepthPairsTestPair struct {
	input string
	out   int
}

var maxDepthPairsTestCases = []maxDepthPairsTestPair{
	{"(1+(2*3)+((8)/4))+1", 3},
	{"(1)+((2))+(((3)))", 3},
}

func TestEvalMaxDepthPairs(t *testing.T) {
	for _, pair := range maxDepthPairsTestCases {
		newOut := maxDepthPairs(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
