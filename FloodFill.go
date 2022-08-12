package leetcode

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	}

	var currentPosition = [][]int{[]int{sr, sc}}
	var nextPosition [][]int
	var currentColor = image[sr][sc]

	for len(currentPosition) > 0 {
		nextPosition = [][]int{}

		for _, position := range currentPosition {
			var row = position[0]
			var col = position[1]

			if image[row][col] == currentColor {
				image[row][col] = color

				if row > 0 {
					nextPosition = append(nextPosition, []int{row - 1, col})
				}

				if row < len(image)-1 {
					nextPosition = append(nextPosition, []int{row + 1, col})
				}

				if col > 0 {
					nextPosition = append(nextPosition, []int{row, col - 1})
				}

				if col < len(image[0])-1 {
					nextPosition = append(nextPosition, []int{row, col + 1})
				}
			}
		}

		currentPosition = nextPosition
	}

	return image
}
