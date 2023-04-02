package leetcode

import (
	"reflect"
	"testing"
)

type convertTemperatureTestPair struct {
	input float64
	out   []float64
}

var convertTemperatureTestCases = []convertTemperatureTestPair{
	{36.5, []float64{309.65, 97.7}},
	{122.11, []float64{395.26, 251.798}},
}

func TestEvalConvertTemperature(t *testing.T) {
	for _, pair := range convertTemperatureTestCases {
		newOut := convertTemperature(pair.input)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.input,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
