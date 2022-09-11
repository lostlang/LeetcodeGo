package leetcode

import (
	"reflect"
	"testing"
)

type partitionStringTestPair struct {
	input string
	out   int
}

var partitionStringTestCases = []partitionStringTestPair{
	{"ssssss", 6},
	{"abacaba", 4},
}

func TestEvalPartitionString(t *testing.T) {
	for _, pair := range partitionStringTestCases {
		var newOut = partitionString(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
