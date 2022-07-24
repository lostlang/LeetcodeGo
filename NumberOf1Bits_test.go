package leetcode

import (
	"reflect"
	"testing"
)

type hammingWeightTestPair struct {
	input uint32
	out   int
}

var hammingWeightTestCases = []hammingWeightTestPair{
	{11, 3},
	{128, 1},
	{4294967293, 31},
}

func TestEvalHammingWeight(t *testing.T) {
	for _, pair := range hammingWeightTestCases {
		var newOut = hammingWeight(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
