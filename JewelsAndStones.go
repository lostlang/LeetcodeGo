package leetcode

func numJewelsInStones(jewels string, stones string) int {
	var out int
	var stonesHash = make(map[rune]int)

	for _, stone := range stones {
		stonesHash[stone]++
	}

	for _, jem := range jewels {
		out += stonesHash[jem]
	}

	return out
}
