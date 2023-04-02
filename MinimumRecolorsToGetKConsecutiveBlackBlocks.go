package leetcode

func minimumRecolors(blocks string, k int) int {
	out := 0
	max := 0

	for i := range blocks {
		if blocks[i] == 'B' {
			max++
		}
		if i >= k {
			if blocks[i-k] == 'B' {
				max--
			}
		}

		if max > out {
			out = max
		}
	}

	return k - out
}
