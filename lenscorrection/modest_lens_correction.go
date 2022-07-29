package lenscorrection

import "github.com/ungerik/go3d/float64/vec2"

//returns the modest lens correction, input and output are relative coordinates [-1,1]
//http://paulbourke.net/miscellaneous/imagewarp/imagewarp.c
//
// TODO: currently cannot be converted back - if a way can be found this would be faster
func modestLensCorrection(p float64, relative *vec2.T) vec2.T {
	denomInv := 1.0 / (1.0 - p*relative.Length())

	return relative.Scaled(denomInv)

	//not quiet the same as the good solution after here, but appropriate
	radius2 := relative.LengthSqr()

	denom := 1 - p*radius2
	xtmp := relative[0] / denom
	ytmp := relative[1] / denom

	denom = 1 - p*(xtmp*xtmp+ytmp*ytmp)
	x := relative[0] / denom
	y := relative[1] / denom

	return vec2.T{x, y}
}
