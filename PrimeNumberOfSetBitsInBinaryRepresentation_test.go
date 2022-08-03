package leetcode

import (
	"reflect"
	"testing"
)

type countPrimeSetBitsTestPair struct {
	inputStart int
	inputEnd   int
	out        int
}

var countPrimeSetBitsTestCases = []countPrimeSetBitsTestPair{
	{6, 10, 4},
	{10, 15, 5},
}

func TestEvalCountPrimeSetBits(t *testing.T) {
	for _, pair := range countPrimeSetBitsTestCases {
		var newOut = countPrimeSetBits(pair.inputStart, pair.inputEnd)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputStart,
				"and", pair.inputEnd,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
