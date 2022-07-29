package leetcode

func checkStraightLine(coordinates [][]int) bool {
	var difX, difY int
	var k float64
	var infinity bool

	difX = coordinates[0][0] - coordinates[1][0]
	difY = coordinates[0][1] - coordinates[1][1]

	if difX == 0 {
		infinity = true
	}

	k = float64(difY) / float64(difX)

	for i := 1; i < len(coordinates); i++ {
		difX = coordinates[i-1][0] - coordinates[i][0]

		if infinity && difX == 0 {
			continue
		}

		if difX == 0 {
			return false
		}

		difY = coordinates[i-1][1] - coordinates[i][1]

		if k != float64(difY)/float64(difX) {
			return false
		}
	}

	return true
}
