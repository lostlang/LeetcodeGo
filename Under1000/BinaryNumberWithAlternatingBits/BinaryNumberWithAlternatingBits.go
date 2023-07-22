package BinaryNumberWithAlternatingBits

func hasAlternatingBits(n int) bool {
	i := 1
	swap := false

	for i < n {
		if swap {
			i = 2*i + 1
		} else {
			i *= 2
		}

		swap = !swap
	}

	return i == n
}
