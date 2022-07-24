package leetcode

import (
	"reflect"
	"testing"
)

type firstUniqCharTestPair struct {
	input string
	out   int
}

var firstUniqCharTestCases = []firstUniqCharTestPair{
	{"leetcode", 0},
	{"loveleetcode", 2},
	{"aabb", -1},
}

func TestEvalFirstUniqChar(t *testing.T) {
	for _, pair := range firstUniqCharTestCases {
		var newOut = firstUniqChar(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
