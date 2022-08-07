package leetcode

import (
	"reflect"
	"testing"
)

type hexadecimalTestPair struct {
	input int
	out   string
}

var hexadecimalTestCases = []hexadecimalTestPair{
	{26, "1a"},
	{-1, "ffffffff"},
}

func TestEvalHexadecimal(t *testing.T) {
	for _, pair := range hexadecimalTestCases {
		var newOut = hexadecimal(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
