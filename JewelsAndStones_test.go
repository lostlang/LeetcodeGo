package leetcode

import (
	"reflect"
	"testing"
)

type numJewelsInStonesTestPair struct {
	inputJewels string
	inputStones string
	out         int
}

var numJewelsInStonesTestCases = []numJewelsInStonesTestPair{
	{"aA", "aAAbbbb", 3},
	{"z", "ZZ", 0},
}

func TestEvalNumJewelsInStones(t *testing.T) {
	for _, pair := range numJewelsInStonesTestCases {
		var newOut = numJewelsInStones(pair.inputJewels, pair.inputStones)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputJewels,
				"and", pair.inputStones,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
