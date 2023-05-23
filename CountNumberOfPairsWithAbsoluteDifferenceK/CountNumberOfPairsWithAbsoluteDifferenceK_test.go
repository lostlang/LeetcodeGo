package CountNumberOfPairsWithAbsoluteDifferenceK

import (
	"reflect"
	"testing"
)

type countKDifferenceTestPair struct {
	inputNums []int
	inputK    int
	out       int
}

var countKDifferenceTestCases = []countKDifferenceTestPair{
	{[]int{1, 2, 2, 1}, 1, 4},
	{[]int{1, 3}, 3, 0},
	{[]int{3, 2, 1, 5, 4}, 2, 3},
}

func TestEvalCountKDifference(t *testing.T) {
	for _, pair := range countKDifferenceTestCases {
		newOut := countKDifference(pair.inputNums, pair.inputK)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputNums,
				"and", pair.inputK,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
