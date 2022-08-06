package leetcode

import (
	"reflect"
	"testing"
)

type poorPigsTestPair struct {
	inputBuckets       int
	inputMinutesToDie  int
	inputMinutesToTest int
	out                int
}

var poorPigsTestCases = []poorPigsTestPair{
	{1000, 15, 60, 5},
	{4, 15, 15, 2},
	{4, 15, 30, 2},
}

func TestEvalPoorPigs(t *testing.T) {
	for _, pair := range poorPigsTestCases {
		var newOut = poorPigs(pair.inputBuckets, pair.inputMinutesToDie, pair.inputMinutesToTest)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputBuckets,
				",", pair.inputMinutesToDie,
				"and", pair.inputMinutesToTest,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
