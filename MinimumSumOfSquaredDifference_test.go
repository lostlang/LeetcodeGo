package leetcode

import (
	"reflect"
	"testing"
)

type minSumSquareDiffTestPair struct {
	inputArr1 []int
	inputArr2 []int
	inputK1   int
	inputK2   int
	out       int
}

var minSumSquareDiffTestCases = []minSumSquareDiffTestPair{
	// {[]int{1, 2, 3, 4}, []int{2, 10, 20, 19}, 0, 0, 579},
	// {[]int{1, 4, 10, 12}, []int{5, 8, 6, 9}, 1, 1, 43},
	// {[]int{1, 4, 10, 12}, []int{5, 8, 6, 9}, 10, 5, 0},
}

func TestEvalMinSumSquareDiff(t *testing.T) {
	for _, pair := range minSumSquareDiffTestCases {
		var newOut int64 = minSumSquareDiff(pair.inputArr1, pair.inputArr2, pair.inputK1, pair.inputK2)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For",
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
