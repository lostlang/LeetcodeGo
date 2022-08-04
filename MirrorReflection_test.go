package leetcode

import (
	"reflect"
	"testing"
)

type mirrorReflectionTestPair struct {
	inputP int
	inputQ int
	out    int
}

var mirrorReflectionTestCases = []mirrorReflectionTestPair{
	{2, 1, 2},
	{3, 1, 1},
	{4, 3, 2},
	{3, 2, 0},
}

func TestEvalMirrorReflection(t *testing.T) {
	for _, pair := range mirrorReflectionTestCases {
		var newOut = mirrorReflection(pair.inputP, pair.inputQ)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputP,
				"and", pair.inputQ,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
