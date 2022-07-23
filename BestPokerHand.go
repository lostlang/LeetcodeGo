package leetcode

func bestHand(ranks []int, suits []byte) string {

	var hashRanks = make(map[int]int)
	var hachSuits = make(map[byte]int)

	var maxRank, maxSuit int

	for _, val := range ranks {
		hashRanks[val]++

		if hashRanks[val] > maxRank {
			maxRank = hashRanks[val]
		}
	}

	for _, char := range suits {
		hachSuits[char]++

		if hachSuits[char] > maxSuit {
			maxSuit = hachSuits[char]
		}
	}

	if maxSuit == 5 {
		return "Flush"
	} else if maxRank >= 3 {
		return "Three of a Kind"
	} else if maxRank == 2 {
		return "Pair"
	}

	return "High Card"
}
