package drawingprecision

func (dp *DrawingPrecision) CalcSteps(total float64) float64 {
	return total * dp.pointsPerUnit
}
