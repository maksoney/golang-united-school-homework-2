package square

import "math"

type sidesType int8

const SidesTriangle sidesType = 3
const SidesSquare sidesType = 4
const SIdesCircle sidesType = 0

func CalcSquare(sideLen float64, sidesNum sidesType) float64 {
	var square float64 = 0

	switch sidesNum {
	case 0:
		// square = (sideLen * sideLen) / (4 * math.Pi)
		square = (math.Pi * sideLen * sideLen)
	case 3:
		square = (sideLen * sideLen * math.Sqrt(3)) / 4
	case 4:
		square = sideLen * sideLen
	}

	return square
}
