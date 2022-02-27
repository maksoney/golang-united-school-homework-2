package square

import "math"

type mySidesType int

const (
	SidesTriangle mySidesType = 3
	SidesSquare   mySidesType = 4
	SidesCircle   mySidesType = 0
)

func CalcSquare(sideLen float64, sidesNum mySidesType) (square float64) {

	switch sidesNum {
	case SidesCircle:
		// square = (sideLen * sideLen) / (4 * math.Pi)
		square = math.Pi * sideLen * sideLen
	case SidesTriangle:
		square = (sideLen * sideLen * math.Sqrt(3)) / 4
	case SidesSquare:
		square = sideLen * sideLen
	default:
		square = 0
	}

	return
}
