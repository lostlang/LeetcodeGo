package leetcode

import (
	"reflect"
	"testing"
)

type isPalindromeTestPair struct {
	input *ListNode
	out   bool
}

var isPalindromeTestCases = []isPalindromeTestPair{
	{generateListNode([]int{1, 2, 2, 1}), true},
	{generateListNode([]int{1, 2}), false},
}

func TestEvalIsPalindrome(t *testing.T) {
	for _, pair := range isPalindromeTestCases {
		var newOut = isPalindrome(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
