package ellipse

import (
	drawingprecision "github.com/simonmau/spacial-base-calculation/drawing-precision"
	rangegen "github.com/simonmau/spacial-base-calculation/range-gen"
	"github.com/ungerik/go3d/float64/vec2"
)

func (c *Data) GenerateEdgePoints(dp *drawingprecision.DrawingPrecision, result *[]vec2.T) {
	min := 0.0
	max := 1920.0
	steps := dp.CalcSteps(max - min)

	arr := rangegen.GenLinearRangeArray(&min, &max, &steps)

	for _, item := range arr {
		c.CalculateYtoArray(item, result)
	}
}
