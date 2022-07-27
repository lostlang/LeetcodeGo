package leetcode

func Subarrays(arr []int) int {
	var sum int

	for i, val := range arr {
		sum += val * ((1 + (len(arr)-i)*(i+1)) / 2)
	}

	return sum
}
