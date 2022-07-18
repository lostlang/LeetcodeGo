package leetcode

func shuffle(nums []int, n int) []int {
	var out = make([]int, n*2)

	for i, val := range nums {
		if i < n {
			out[i*2] = val
		} else {
			out[(i-n)*2+1] = val
		}
	}
	return out
}
