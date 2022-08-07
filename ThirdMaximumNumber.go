package leetcode

import (
	"sort"
)

func thirdMax(nums []int) int {
	var out []int

	for _, num := range nums {
		if len(out) < 3 {
			switch len(out) {
			case 0:
				out = append(out, num)
			case 1:
				if num != out[0] {
					out = append(out, num)
				}
			case 2:
				if num != out[0] && num != out[1] {
					out = append(out, num)
				}
			}

			sort.Ints(out)
			continue
		}

		if num > out[2] {
			out[0] = out[1]
			out[1] = out[2]
			out[2] = num

			continue
		}

		if num > out[1] && num != out[2] {
			out[0] = out[1]
			out[1] = num
			continue
		}

		if num > out[0] && num != out[1] && num != out[2] {
			out[0] = num
			continue
		}
	}

	if len(out) == 3 {
		return out[0]
	}

	return out[len(out)-1]
}
