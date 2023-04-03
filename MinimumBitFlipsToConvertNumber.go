package leetcode

func minBitFlips(start int, goal int) int {
	output := 0
	xor := start ^ goal
	for xor > 0 {
		if xor&1 == 1 {
			output++
		}
		xor >>= 1
	}

	return output
}
