package leetcode

import "sort"

func minMovesToSeat(seats []int, students []int) int {
	var out int
	sort.Ints(seats)
	sort.Ints(students)

	for i := 0; i < len(students); i++ {
		out += abs(students[i] - seats[i])
	}

	return out
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
