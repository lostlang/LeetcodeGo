package leetcode

func kWeakestRows(mat [][]int, k int) []int {
	var sums = make([][2]int, len(mat))

	for index, str := range mat {
		for _, val := range str {
			sums[index][1] += val
		}
		sums[index][0] = index
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat); j++ {
			if sums[i][1] < sums[j][1] || (sums[i][1] == sums[j][1] && sums[i][0] < sums[j][0]) {
				sums[i], sums[j] = sums[j], sums[i]
			}
		}
	}

	var out = make([]int, k)

	for i := 0; i < k; i++ {
		out[i] = sums[i][0]
	}

	return out
}
