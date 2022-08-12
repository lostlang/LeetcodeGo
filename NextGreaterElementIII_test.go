package leetcode

import (
	"reflect"
	"testing"
)

type nextGreaterElementIIITestPair struct {
	input int
	out   int
}

var nextGreaterElementIIITestCases = []nextGreaterElementIIITestPair{
	{12, 21},
	{21, -1},
	{230241, 230412},
	{2147483476, 2147483647},
	{12443322, 13222344},
}

func TestEvalNextGreaterElementIII(t *testing.T) {
	for _, pair := range nextGreaterElementIIITestCases {
		var newOut = nextGreaterElementIII(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
