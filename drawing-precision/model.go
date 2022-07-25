package drawingprecision

type DrawingPrecision struct {
	pointsPerUnit         float64
	distanceBetweenPoints float64
}

func GenDrawPrecByPointsPerUnit(pointsPerUnit float64) *DrawingPrecision {
	return &DrawingPrecision{
		pointsPerUnit:         pointsPerUnit,
		distanceBetweenPoints: 1.0 / pointsPerUnit,
	}
}

func GenDrawPrecByDistance(distance float64) *DrawingPrecision {
	return &DrawingPrecision{
		pointsPerUnit:         1.0 / distance,
		distanceBetweenPoints: distance,
	}
}
