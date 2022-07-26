package leetcode

import (
	"reflect"
	"testing"
)

type largestOddNumberTestPair struct {
	input string
	out   string
}

var largestOddNumberTestCases = []largestOddNumberTestPair{
	{"52", "5"},
	{"4206", ""},
	{"35427", "35427"},
}

func TestEvalLargestOddNumber(t *testing.T) {
	for _, pair := range largestOddNumberTestCases {
		var newOut = largestOddNumber(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
