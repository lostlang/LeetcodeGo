package leetcode

import (
	"reflect"
	"testing"
)

type fizzBuzzTestPair struct {
	input int
	out   []string
}

var fizzBuzzTestCases = []fizzBuzzTestPair{
	{3, []string{"1", "2", "Fizz"}},
	{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
	{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
}

func TestEvalFizzBuzz(t *testing.T) {
	for _, pair := range fizzBuzzTestCases {
		var newOut = fizzBuzz(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
