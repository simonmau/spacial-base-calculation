package triangle

import (
	drawingprecision "github.com/simonmau/spacial-base-calculation/drawing-precision"
	"github.com/simonmau/spacial-base-calculation/line"
	"github.com/ungerik/go3d/float64/vec3"
)

func (t *Triangle) GenerateEdgePoints(dp *drawingprecision.DrawingPrecision, result *[]vec3.T) {
	item := line.GenerateWithTwoPoints(t.Pt0, t.Pt1)
	item.GenerateEdgePoints(dp, result)

	item = line.GenerateWithTwoPoints(t.Pt1, t.Pt2)
	item.GenerateEdgePoints(dp, result)

	item = line.GenerateWithTwoPoints(t.Pt2, t.Pt0)
	item.GenerateEdgePoints(dp, result)
}
