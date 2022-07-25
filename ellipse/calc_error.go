package ellipse

import (
	"math"

	"github.com/ungerik/go3d/float64/vec2"
)

func (d *Data) CalcAverageError(pointSlice *[]vec2.T) *float64 {
	if len(*pointSlice) == 0 {
		return nil
	}

	errorTotal := 0.0
	count := 0.0

	for _, item := range *pointSlice {
		errCurrent := d.CalcError(&item)

		if errCurrent != nil {
			errorTotal += *errCurrent
			count++
		}
	}

	average := errorTotal / count

	return &average
}

//error sqared is returned (so it is always positive and goes up)
func (d *Data) CalcError(point *vec2.T) *float64 {
	BxE := d.B*point[0] + d.E
	AxD := d.A*point[0] + d.D

	if d.C > -_ERROR_OFFSET && d.C < _ERROR_OFFSET && BxE > -_ERROR_OFFSET && BxE < _ERROR_OFFSET {
		y1 := -(point[0]*AxD + d.F) / BxE

		diff := y1 - point[1]
		diff *= diff

		return &diff
	} else if d.C > 0.0 {
		wurzel := math.Sqrt(BxE*BxE - 4.0*d.C*(point[0]*AxD+d.F))

		divisor := 1.0 / (2.0 * d.C)

		if wurzel > -_ERROR_OFFSET && wurzel < _ERROR_OFFSET {
			y3 := -BxE * divisor
			val := y3 - point[1]
			val *= val

			return &val
		}

		y1 := -(wurzel + BxE) * divisor
		y2 := -(-wurzel + BxE) * divisor

		diff1 := y1 - point[1]
		diff2 := y2 - point[1]

		diff1 *= diff1
		diff2 *= diff2

		if diff1 < diff2 {
			return &diff1
		}

		return &diff2
	}

	return nil
}
