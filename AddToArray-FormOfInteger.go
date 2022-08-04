package leetcode

func addToArrayForm(num []int, k int) []int {
	var out = []int{}

	for i := len(num) - 1; i >= 0; i-- {
		k += num[i]

		out = append(out, k%10)

		k /= 10
	}

	for k > 0 {
		out = append(out, k%10)
		k /= 10
	}

	for i := 0; i < len(out)/2; i++ {
		out[i], out[len(out)-i-1] = out[len(out)-i-1], out[i]
	}

	return out
}
