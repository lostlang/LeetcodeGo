package leetcode

import (
	"reflect"
	"testing"
)

type smallestEvenMultipleTestPair struct {
	input int
	out   int
}

var smallestEvenMultipleTestCases = []smallestEvenMultipleTestPair{
	{5, 10},
	{6, 6},
}

func TestEvalSmallestEvenMultiple(t *testing.T) {
	for _, pair := range smallestEvenMultipleTestCases {
		var newOut = smallestEvenMultiple(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
