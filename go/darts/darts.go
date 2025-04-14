package darts

import "math"

const (
	radiusOuterCircle  = 10
	radiusMiddleCircle = 5
	radiusInnerCircle  = 1

	scoreOutOfTarget  = 0
	scoreOuterCircle  = 1
	scoreMiddleCircle = 5
	scoreInnerCircle  = 10
)

func Score(x, y float64) int {
	r := math.Sqrt(x*x + y*y)
	switch {
	case r <= radiusInnerCircle:
		return scoreInnerCircle
	case r <= radiusMiddleCircle:
		return scoreMiddleCircle
	case r <= radiusOuterCircle:
		return scoreOuterCircle
	default:
		return scoreOutOfTarget
	}
}
