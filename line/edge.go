package line

import (
	drawingprecision "github.com/simonmau/spacial-base-calculation/drawing-precision"
	"github.com/ungerik/go3d/float64/vec3"
)

func (c *Line) GenerateEdgePoints(dp *drawingprecision.DrawingPrecision, result *[]vec3.T) {
	dir := vec3.Sub(c.Pt1, c.Pt0)

	l := dir.Length()

	steps := dp.CalcSteps(l) - 1.0
	dp.ScaleDirection(&dir)

	pos := vec3.Add(c.Pt0, &dir)

	*result = append(*result, *c.Pt0)

	for step := 1.0; step < steps; step++ {
		*result = append(*result, pos)
		pos = vec3.Add(&pos, &dir)
	}

	*result = append(*result, pos, *c.Pt1)
}
