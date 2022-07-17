package leetcode

import (
	"reflect"
	"testing"
)

type tribonacciTestPair struct {
	input int
	out   int
}

var tribonacciTestCases = []tribonacciTestPair{
	{4, 4},
	{25, 1389537},
}

func TestEvalTribonacci(t *testing.T) {
	for _, pair := range tribonacciTestCases {
		var newOut = tribonacci(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
