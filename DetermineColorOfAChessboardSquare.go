package leetcode

func squareIsWhite(coordinates string) bool {
	var x = coordinates[0] - 'a'
	var y = coordinates[1] - '1'

	if y%2 == 1 {
		return x%2 == 0
	} else {
		return x%2 != 0
	}
}
