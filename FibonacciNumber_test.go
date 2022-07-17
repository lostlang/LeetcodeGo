package leetcode

import (
	"reflect"
	"testing"
)

type fibTestPair struct {
	input int
	out   int
}

var fibTestCases = []fibTestPair{
	{2, 1},
	{3, 2},
	{4, 3},
}

func TestEvalFib(t *testing.T) {
	for _, pair := range fibTestCases {
		var newOut = fib(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
