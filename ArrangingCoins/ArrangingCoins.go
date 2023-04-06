package ArrangingCoins

import (
	"math"
)

func arrangeCoins(n int) int {
	x := (math.Sqrt(float64(n*8+1)) - 1) / 2

	output := int(x)

	return output
}
