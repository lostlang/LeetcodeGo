package leetcode

func arithmeticTriplets(nums []int, diff int) int {
	var out int
	var hash = make(map[int]int)

	for i, num := range nums {
		hash[num] = i + 1
	}

	for num := range hash {
		if hash[num+diff] != 0 && hash[num+diff*2] != 0 {
			out++
		}
	}

	return out
}
