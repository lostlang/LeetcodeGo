package leetcode

import (
	"reflect"
)

func findRotation(mat [][]int, target [][]int) bool {

	for i := 0; i < 4; i++ {
		if reflect.DeepEqual(mat, target) {
			return true
		}

		mat = rotate90(mat)
	}

	return false
}

func rotate90(arr [][]int) [][]int {
	var out = make([][]int, len(arr))

	for i := 0; i < len(arr[0]); i++ {
		out[i] = make([]int, len(arr))

		for j := 0; j < len(arr); j++ {
			out[i][j] = arr[len(arr)-1-j][i]
		}
	}

	return out
}
