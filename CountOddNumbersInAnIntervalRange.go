package leetcode

func countOdds(low int, high int) int {
	var out = (high - low) / 2

	if low%2 == 1 || high%2 == 1 {
		out++
	}

	return out
}
