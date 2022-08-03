package leetcode

import (
	"reflect"
	"testing"
)

type readBinaryWatchTestPair struct {
	input int
	out   []string
}

var readBinaryWatchTestCases = []readBinaryWatchTestPair{
	{1, []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"}},
	{0, []string{"0:00"}},
	{9, []string{}},
}

func TestEvalReadBinaryWatch(t *testing.T) {
	for _, pair := range readBinaryWatchTestCases {
		var newOut = readBinaryWatch(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
