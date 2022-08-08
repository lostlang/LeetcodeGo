package leetcode

func countSegments(s string) int {
	var out int
	var word bool

	for _, c := range s {
		if c != ' ' {
			word = true
		} else {
			if word {
				out++
			}

			word = false
		}
	}

	if word {
		out++
	}

	return out
}
