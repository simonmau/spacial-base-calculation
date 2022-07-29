package lenscorrection

import (
	"math"

	"github.com/ungerik/go3d/float64/vec2"
)

// p sould not be much bigger than 1 - LIMIT AT 0 IS HARD-> WRONG VALUES
func CorrectSexyToRaw(p *float64, point *vec2.T, width, height *float64) *vec2.T {
	relativeCoordinates := convertImageToRelativeCoordiantes(point, width, height)

	n := relativeCoordinates.Length()

	r := n

	//kinda approximation

	if math.Abs(*p) > _errorTolerance {
		wurzel := math.Sqrt(4.0*n**p + 1)

		r1 := -(wurzel + 1) / (2 * *p)
		r2 := (wurzel - 1) / (2 * *p)

		if r1 < 0 {
			r = r2
		} else if r2 < 0 {
			r = r1
		} else {
			r = math.Min(r1, r2)
		}
	}

	// if math.Abs(_a+_c**p) < _errorTolerance {
	// 	if math.Abs(_b+_d**p) < _errorTolerance {
	// 		panic("div by 0")
	// 	}

	// 	r = n / (_b + _d**p)
	// } else {
	// 	wurzel2 := 4.0*_a*n + _b*_b + 2.0*_b*_d**p + 4.0*_c*n**p + _d*_d**p**p
	// 	wurzel := math.Sqrt(wurzel2)

	// 	div := 1.0 / (2.0 * (_a + _c**p))

	// 	v1 := -(wurzel + _b + _d**p) * div
	// 	v2 := -(-wurzel + _b + _d**p) * div

	// 	if v2 < _errorTolerance {
	// 		r = v1
	// 	} else if v1 < _errorTolerance {
	// 		r = v2
	// 	} else {
	// 		r = math.Min(v1, v2)
	// 	}
	// }

	corr := relativeCoordinates.Scaled(r / n)

	tmp := convertRelativeToImageCoordinates(&corr, width, height)
	return &tmp
}
