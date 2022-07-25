package leetcode

func nearestValidPoint(x int, y int, points [][]int) int {
	var index = -1
	var distance = -1
	var newDistX, newDistY, newDist int

	for i, val := range points {
		newDistX = val[0] - x
		newDistY = val[1] - y

		if newDistX < 0 {
			newDistX = -newDistX
		}

		if newDistY < 0 {
			newDistY = -newDistY
		}

		newDist = newDistX + newDistY

		if (val[0] == x || val[1] == y) && (newDist < distance || distance == -1) {
			distance = newDist
			index = i
		}
	}

	return index
}
