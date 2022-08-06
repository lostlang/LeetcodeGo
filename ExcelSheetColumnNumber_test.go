package leetcode

import (
	"reflect"
	"testing"
)

type titleToNumberTestPair struct {
	input string
	out   int
}

var titleToNumberTestCases = []titleToNumberTestPair{
	{"A", 1},
	{"AB", 28},
	{"ZY", 701},
}

func TestEvalTitleToNumber(t *testing.T) {
	for _, pair := range titleToNumberTestCases {
		var newOut = titleToNumber(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
