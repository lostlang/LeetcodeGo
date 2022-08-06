package leetcode

import (
	"reflect"
	"testing"
)

type mergeSimilarItemsTestPair struct {
	inputItem1 [][]int
	inputItem2 [][]int
	out        [][]int
}

var mergeSimilarItemsTestCases = []mergeSimilarItemsTestPair{
	{[][]int{[]int{1, 1}, []int{4, 5}, []int{3, 8}}, [][]int{[]int{3, 1}, []int{1, 5}}, [][]int{[]int{1, 6}, []int{3, 9}, []int{4, 5}}},
}

func TestEvalMergeSimilarItems(t *testing.T) {
	for _, pair := range mergeSimilarItemsTestCases {
		var newOut = mergeSimilarItems(pair.inputItem1, pair.inputItem2)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputItem1,
				"and", pair.inputItem2,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
