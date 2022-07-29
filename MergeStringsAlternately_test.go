package leetcode

import (
	"reflect"
	"testing"
)

type mergeAlternatelyTestPair struct {
	inputW1 string
	inputW2 string
	out     string
}

var mergeAlternatelyTestCases = []mergeAlternatelyTestPair{
	{"abc", "pqr", "apbqcr"},
	{"ab", "pqrs", "apbqrs"},
	{"abcd", "pq", "apbqcd"},
}

func TestEvalMergeAlternately(t *testing.T) {
	for _, pair := range mergeAlternatelyTestCases {
		var newOut = mergeAlternately(pair.inputW1, pair.inputW2)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputW1,
				"and", pair.inputW2,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
