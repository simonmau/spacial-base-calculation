package anglehelper

//calculates the nearest center for the two angles given
func Center(v1, v2 float64) float64 {
	diff := DifferenceWithDirection(v1, v2, true)
	return Modulo(v1 + (diff * 0.5))
}
