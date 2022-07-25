package drawingprecision

import "github.com/ungerik/go3d/float64/vec3"

func (dp *DrawingPrecision) ScaleDirection(dir *vec3.T) {
	l := dir.Length()
	dir.Scale(dp.distanceBetweenPoints / l)
}
