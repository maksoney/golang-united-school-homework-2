package square

import "math"

type mySidesType int

// const SidesTriangle sidesType = 3
// const SidesSquare sidesType = 4
// const SidesCircle sidesType = 0

func CalcSquare(sideLen float64, sidesNum mySidesType) float64 {
	var square float64

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
