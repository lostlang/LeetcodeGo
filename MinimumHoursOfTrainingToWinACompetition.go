package leetcode

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	var out int
	var sumEnergy int

	for _, e := range energy {
		sumEnergy += e
	}

	if sumEnergy >= initialEnergy {
		out = sumEnergy - initialEnergy + 1
	}

	for _, e := range experience {
		if e >= initialExperience {
			out += e - initialExperience + 1
			initialExperience = e + 1
		}
		initialExperience += e
	}

	return out
}
