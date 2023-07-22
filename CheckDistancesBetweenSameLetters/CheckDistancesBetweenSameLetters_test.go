package CheckDistancesBetweenSameLetters

import (
	"reflect"
	"testing"
)

type checkDistancesTestPair struct {
	inputS        string
	inputDistance []int
	output        bool
}

var checkDistancesTestCases = []checkDistancesTestPair{
	{"abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, true},
	{"aa", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, false},
}

func TestEvalCheckDistances(t *testing.T) {
	for _, pair := range checkDistancesTestCases {
		newOut := checkDistances(pair.inputS, pair.inputDistance)
		if !reflect.DeepEqual(newOut, pair.output) {
			t.Error(
				"For", pair.inputS, pair.inputDistance,
				"expected", pair.output,
				"got", newOut,
			)
		}
	}
}
