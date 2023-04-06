package AddToArrayFormOfInteger

func addToArrayForm(num []int, k int) []int {
	output := []int{}

	for i := len(num) - 1; i >= 0; i-- {
		k += num[i]
		output = append(output, k%10)
		k /= 10
	}

	for k > 0 {
		output = append(output, k%10)
		k /= 10
	}

	for i := 0; i < len(output)/2; i++ {
		output[i], output[len(output)-i-1] = output[len(output)-i-1], output[i]
	}

	return output
}
