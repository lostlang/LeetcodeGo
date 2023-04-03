package leetcode

import (
	"reflect"
	"testing"
)

type uniqueMorseRepresentationsTestPair struct {
	input []string
	out   int
}

var uniqueMorseRepresentationsTestCases = []uniqueMorseRepresentationsTestPair{
	{[]string{"gin", "zen", "gig", "msg"}, 2},
	{[]string{"a"}, 1},
}

func TestEvalUniqueMorseRepresentations(t *testing.T) {
	for _, pair := range uniqueMorseRepresentationsTestCases {
		newOut := uniqueMorseRepresentations(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
