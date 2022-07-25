package advancedlenscorrection

import "github.com/ungerik/go3d/float64/vec2"

// p sould not be much bigger than 1
func CorrectRawToSexy(p *float64, point, center *vec2.T, width, height *float64) *vec2.T {
	relativeCoordinates := convertImageToRelativeCoordiantes(point, center, width, height)

	r := relativeCoordinates.Length()

	//http://paulbourke.net/miscellaneous/imagewarp/imagewarp.c
	denomInv := 1.0 / (1.0 - *p*r)

	xtmp := relativeCoordinates[0] * denomInv
	ytmp := relativeCoordinates[1] * denomInv

	if xtmp <= -1.0 || xtmp >= 1.0 || ytmp <= -1.0 || ytmp >= 1.0 {
		return nil
	}

	tmp := convertRelativeToImageCoordinates(relativeCoordinates.Scale(denomInv), center, width, height)
	return &tmp
}
