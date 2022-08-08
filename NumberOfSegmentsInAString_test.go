package leetcode

import (
	"reflect"
	"testing"
)

type countSegmentsTestPair struct {
	input string
	out   int
}

var countSegmentsTestCases = []countSegmentsTestPair{
	{"Hello, my name is John", 5},
	{"a", 1},
}

func TestEvalCountSegments(t *testing.T) {
	for _, pair := range countSegmentsTestCases {
		var newOut = countSegments(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
