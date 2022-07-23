package leetcode

import (
	"reflect"
	"testing"
)

type bestHandTestPair struct {
	inputRank []int
	inputSuit []byte
	out       string
}

var bestHandTestCases = []bestHandTestPair{
	{[]int{13, 2, 3, 1, 9}, []byte{'a', 'a', 'a', 'a', 'a'}, "Flush"},
	{[]int{4, 4, 2, 4, 4}, []byte{'d', 'a', 'a', 'b', 'c'}, "Three of a Kind"},
	{[]int{10, 10, 2, 12, 9}, []byte{'a', 'b', 'c', 'a', 'd'}, "Pair"},
}

func TestEvalBestHand(t *testing.T) {
	for _, pair := range bestHandTestCases {
		var newOut = bestHand(pair.inputRank, pair.inputSuit)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputRank,
				"and", pair.inputSuit,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
