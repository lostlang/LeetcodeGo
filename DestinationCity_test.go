package leetcode

import (
	"reflect"
	"testing"
)

type destCityTestPair struct {
	input [][]string
	out   string
}

var destCityTestCases = []destCityTestPair{
	{[][]string{[]string{"London", "New York"}}, "New York"},
}

func TestEvalDestCity(t *testing.T) {
	for _, pair := range destCityTestCases {
		var newOut = destCity(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
