package leetcode

func nextGreatestLetter(letters []byte, target byte) byte {
	var lettersList = make([]int, 26)

	for _, letter := range letters {
		lettersList[letter-'a'] = 1
	}

	for i, letter := range lettersList {
		if letter == 1 && byte(i)+'a' > target {
			return byte(i) + 'a'
		}
	}

	for i, letter := range lettersList {
		if letter == 1 {
			return byte(i) + 'a'
		}
	}

	return 0
}
