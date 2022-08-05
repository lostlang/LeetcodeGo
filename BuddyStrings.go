package leetcode

import (
	"reflect"
)

func buddyStrings(s string, goal string) bool {
	var hashS = stringToHashmapByte(s)
	var hashG = stringToHashmapByte(goal)

	if !reflect.DeepEqual(hashS, hashG) {
		return false
	}

	var diff = 0

	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			diff++
		}
	}

	if diff == 2 {
		return true
	}

	if diff != 0 {
		return false
	}

	for char := range hashS {
		if hashS[char] > 1 {
			return true
		}
	}

	return false
}
