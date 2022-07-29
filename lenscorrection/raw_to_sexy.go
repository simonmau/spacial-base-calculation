package lenscorrection

import (
	"github.com/ungerik/go3d/float64/vec2"
)

const (
	_a = -0.063026124053362
	_b = 1.05272839776571
	_c = -1.27267833812949
	_d = 0.386857312388595

	_errorTolerance = 0.00001
)

// p sould not be much bigger than 1 - LIMIT AT 0 IS HARD-> NO VALUES
func CorrectRawToSexy(p *float64, point *vec2.T, width, height *float64) *vec2.T {
	relativeCoordinates := convertImageToRelativeCoordiantes(point, width, height)

	r2 := relativeCoordinates.LengthSqr()
	r := relativeCoordinates.Length()

	// rN := r2*_a + r*_b + *p*r2*_c + *p*r*_d
	rN := r + r2**p //kinda approximation

	corr := relativeCoordinates.Scaled(rN / r)

	tmp := convertRelativeToImageCoordinates(&corr, width, height)
	return &tmp
}
