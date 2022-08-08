package leetcode

func dailyTemperatures(temperatures []int) []int {
	var out = make([]int, len(temperatures))

	for i := len(temperatures) - 2; i >= 0; i-- {
		for j := i + 1; j < len(temperatures); {
			if temperatures[j] > temperatures[i] {
				out[i] = j - i
				break
			}

			if temperatures[j] <= temperatures[i] && out[j] == 0 {
				out[i] = 0
				break
			}

			j += out[j]

		}
	}

	return out
}
