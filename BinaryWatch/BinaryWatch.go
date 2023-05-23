package BinaryWatch

import (
	"math/bits"
	"strconv"
)

func readBinaryWatch(turnedOn int) []string {
	if turnedOn == 0 {
		return []string{"0:00"}
	}

	if turnedOn > 8 {
		return []string{}
	}

	out := []string{}

	for hour := 0; hour < 12; hour++ {
		for minute := 0; minute < 10; minute++ {
			if bits.OnesCount(uint(hour))+bits.OnesCount(uint(minute)) == turnedOn {
				out = append(out, strconv.Itoa(hour)+":0"+strconv.Itoa(minute))
			}
		}

		for minute := 10; minute < 60; minute++ {
			if bits.OnesCount(uint(hour))+bits.OnesCount(uint(minute)) == turnedOn {
				out = append(out, strconv.Itoa(hour)+":"+strconv.Itoa(minute))
			}
		}
	}

	return out
}
