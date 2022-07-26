package leetcode

import (
	"reflect"
	"testing"
)

type maximum69NumberTestPair struct {
	input int
	out   int
}

var maximum69NumberTestCases = []maximum69NumberTestPair{
	{9669, 9969},
	{9996, 9999},
	{9999, 9999},
}

func TestEvalMaximum69Number(t *testing.T) {
	for _, pair := range maximum69NumberTestCases {
		var newOut = maximum69Number(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
