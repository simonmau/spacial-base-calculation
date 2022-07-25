package spherecoordinates

import (
	"math"

	"github.com/ungerik/go3d/float64/vec3"
)

func ConvertToSphereCoordinates(pt1OrCamera, pt2 *vec3.T) Coor {
	dX := pt2[0] - pt1OrCamera[0]
	dY := pt2[1] - pt1OrCamera[1]
	dZ := pt2[2] - pt1OrCamera[2]

	rInv := 1.0 / math.Sqrt(dX*dX+dY*dY+dZ*dZ) //could be faster

	theta := math.Acos(dZ * rInv)

	var phi float64

	if dX > -roundingError && dX < roundingError {
		phi = math.Pi * 0.5

		if dY < 0.0 {
			phi = -phi
		}
	} else if dX > 0.0 {
		phi = math.Atan(dY / dX)
	} else if dX < 0.0 && dY >= 0.0 {
		phi = math.Atan(dY/dX) + math.Pi
	} else {
		phi = math.Atan(dY/dX) - math.Pi
	}

	return Coor{
		Phi:   phi,
		Theta: theta,
	}
}
