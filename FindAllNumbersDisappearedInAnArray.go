package leetcode

func findDisappearedNumbers(nums []int) []int {
	var out []int
	hash := make(map[int]bool)

	for _, num := range nums {
		hash[num] = true
	}

	for i := 1; i <= len(nums); i++ {
		if !hash[i] {
			out = append(out, i)
		}
	}

	return out
}
