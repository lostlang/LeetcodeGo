package leetcode

import (
	"reflect"
	"testing"
)

type checkDistancesTestPair struct {
	inputS        string
	inputDistance []int
	out           bool
}

var checkDistancesTestCases = []checkDistancesTestPair{
	{"abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, true},
	{"aa", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, false},
}

func TestEvalCheckDistances(t *testing.T) {
	for _, pair := range checkDistancesTestCases {
		var newOut = checkDistances(pair.inputS, pair.inputDistance)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputS,
				"and", pair.inputDistance,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
