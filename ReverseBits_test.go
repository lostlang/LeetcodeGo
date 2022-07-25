package leetcode

import (
	"reflect"
	"testing"
)

type reverseBitsTestPair struct {
	input uint32
	out   uint32
}

var reverseBitsTestCases = []reverseBitsTestPair{
	{1, 2147483648},
	{2, 1073741824},
}

func TestEvalReverseBits(t *testing.T) {
	for _, pair := range reverseBitsTestCases {
		var newOut = reverseBits(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
