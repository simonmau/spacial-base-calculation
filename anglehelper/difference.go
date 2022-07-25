package anglehelper

import (
	"math"
)

//returns difference between the two angles, range [0 to +π]
func Difference(old, new float64) float64 {
	diff := DifferenceWithDirection(old, new, true)

	if diff > math.Pi {
		return math.Abs(TwoPi - diff)
	}

	return math.Abs(diff)
}

//returns difference between the two angles, range ]-π,+π]
func DifferenceWithDirection(old, new float64, clockwise bool) float64 {
	mod := Modulo(new - old)

	if mod > math.Pi {
		mod = -(TwoPi - mod)
	}

	if clockwise {
		return -mod
	}

	return mod
}
