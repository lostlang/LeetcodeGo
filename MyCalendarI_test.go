package leetcode

import (
	"reflect"
	"testing"
)

type bookingCalendarTestPair struct {
	inputStart int
	inputEnd   int
	out        bool
}

var bookingCalendarTestCases = []bookingCalendarTestPair{
	{10, 20, true},
	{15, 25, false},
	{20, 30, true},
}

func TestEvalCalendar(t *testing.T) {
	var testCalendar = ConstructorCalendar()

	for _, pair := range bookingCalendarTestCases {
		var newOut = testCalendar.Book(pair.inputStart, pair.inputEnd)
		if !reflect.DeepEqual(newOut, pair.out) {
			t.Error(
				"For", pair.inputStart,
				"and", pair.inputEnd,
				"expected", pair.out,
				"got", newOut,
			)
		}
	}
}
