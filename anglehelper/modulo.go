package anglehelper

import "math"

func Modulo(value float64) float64 {
	return value - (TwoPi * math.Floor(value*ReverseTwoPi))
}
