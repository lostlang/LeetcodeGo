package IntersectionOfTwoArrays

import (
	"reflect"
	"sort"
	"testing"
)

type intersectionTestPair struct {
	inputArr1 []int
	inputArr2 []int
	out       []int
}

var intersectionTestCases = []intersectionTestPair{
	{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
	{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
}

func TestEvalIntersection(t *testing.T) {
	for _, pair := range intersectionTestCases {
		newOut := intersection(pair.inputArr1, pair.inputArr2)
		sort.Ints(newOut)
		sort.Ints(pair.out)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputArr1, pair.inputArr2,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
