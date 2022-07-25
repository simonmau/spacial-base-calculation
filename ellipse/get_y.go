package ellipse

import (
	"math"

	"github.com/ungerik/go3d/float64/vec2"
)

func (d *Data) CalculateY(x float64) []vec2.T {
	result := make([]vec2.T, 0, 2)

	d.CalculateYtoArray(x, &result)

	return result
}

func (d *Data) CalculateYtoArray(x float64, result *[]vec2.T) {
	tmpInside := d.B*x + d.E

	if d.C > -_ERROR_OFFSET && d.C < _ERROR_OFFSET && tmpInside > -_ERROR_OFFSET && tmpInside < _ERROR_OFFSET {
		y1 := -(x*(d.A*x+d.D) + d.F) / tmpInside
		*result = append(*result, vec2.T{x, y1})
	} else if d.C > 0.0 {
		insideWurzel := tmpInside*tmpInside - 4.0*d.C*(x*(d.A*x+d.D)+d.F)

		if insideWurzel > _ERROR_OFFSET {
			wurzel := math.Sqrt(insideWurzel)

			divisor := 1.0 / (2.0 * d.C)

			*result = append(*result, vec2.T{x, -(wurzel + tmpInside) * divisor})
			*result = append(*result, vec2.T{x, -(-wurzel + tmpInside) * divisor})
		}
	}
}
