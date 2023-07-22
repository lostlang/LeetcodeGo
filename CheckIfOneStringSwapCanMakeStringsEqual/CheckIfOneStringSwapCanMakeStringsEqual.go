package CheckIfOneStringSwapCanMakeStringsEqual

func areAlmostEqual(s1 string, s2 string) bool {
	r1 := []rune(s1)
	r2 := []rune(s2)
	var notEq []int

	if len(r1) != len(r2) {
		return false
	}

	for i := range r1 {
		if r1[i] != r2[i] {
			notEq = append(notEq, i)
		}
	}

	if len(notEq) == 0 {
		return true
	}

	if len(notEq) == 2 {
		if r1[notEq[0]] == r2[notEq[1]] && r1[notEq[1]] == r2[notEq[0]] {
			return true
		}
	}

	return false
}
