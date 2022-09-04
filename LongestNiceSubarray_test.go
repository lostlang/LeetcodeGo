package leetcode

import (
	"reflect"
	"testing"
)

type longestNiceSubarrayTestPair struct {
	input []int
	out   int
}

var longestNiceSubarrayTestCases = []longestNiceSubarrayTestPair{
	{[]int{1, 3, 8, 48, 10}, 3},
	{[]int{3, 1, 5, 11, 13}, 1},
	{[]int{744437702, 379056602, 145555074, 392756761, 560864007, 934981918, 113312475, 1090, 16384, 33, 217313281, 117883195, 978927664}, 3},
	{[]int{84139415, 693324769, 614626365, 497710833, 615598711, 264, 65552, 50331652, 1, 1048576, 16384, 544, 270532608, 151813349, 221976871, 678178917, 845710321, 751376227, 331656525, 739558112, 267703680}, 8},
	{[]int{1}, 1},
}

func TestEvalLongestNiceSubarray(t *testing.T) {
	for _, pair := range longestNiceSubarrayTestCases {
		var newOut = longestNiceSubarray(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
