package leetcode

import (
	"reflect"
	"testing"
)

type minPartitionsTestPair struct {
	input string
	out   int
}

var minPartitionsTestCases = []minPartitionsTestPair{
	{"32", 3},
	{"82734", 8},
	{"27346209830709182346", 9},
}

func TestEvalMinPartitions(t *testing.T) {
	for _, pair := range minPartitionsTestCases {
		var newOut = minPartitions(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
