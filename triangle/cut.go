package triangle

import (
	"github.com/ungerik/go3d/float64/vec3"
)

const (
	roundingError = 0.0001
)

//https://en.wikipedia.org/wiki/M%C3%B6ller%E2%80%93Trumbore_intersection_algorithm
func (tri *Triangle) WasCutf(pt1OrCamera, pt2 *vec3.T) bool {
	ray := vec3.Sub(pt2, pt1OrCamera)

	edge1 := vec3.Sub(tri.Pt1, tri.Pt0)
	edge2 := vec3.Sub(tri.Pt2, tri.Pt0)

	h := vec3.Cross(&ray, &edge2)
	a := vec3.Dot(&edge1, &h)

	if a > -roundingError && a < roundingError {
		return false // This ray is parallel to this triangle.
	}

	f := 1.0 / a

	s := vec3.Sub(pt1OrCamera, tri.Pt0)
	u := f * vec3.Dot(&s, &h)

	if u < 0.0 || u > 1.0 {
		return false
	}

	q := vec3.Cross(&s, &edge1)
	v := f * vec3.Dot(&ray, &q)
	if v < 0.0 || u+v > 1.0 {
		return false
	}

	t := f * vec3.Dot(&edge2, &q)

	return t > roundingError && t < 1.0+roundingError
}

//https://en.wikipedia.org/wiki/M%C3%B6ller%E2%80%93Trumbore_intersection_algorithm
func (tri *Triangle) WasCut(pt1OrCamera, pt2 *vec3.T) bool {
	rX := pt2[0] - pt1OrCamera[0]
	rY := pt2[1] - pt1OrCamera[1]
	rZ := pt2[2] - pt1OrCamera[2]

	e1X := tri.Pt1[0] - tri.Pt0[0]
	e1Y := tri.Pt1[1] - tri.Pt0[1]
	e1Z := tri.Pt1[2] - tri.Pt0[2]

	e2X := tri.Pt2[0] - tri.Pt0[0]
	e2Y := tri.Pt2[1] - tri.Pt0[1]
	e2Z := tri.Pt2[2] - tri.Pt0[2]

	hX := rY*e2Z - rZ*e2Y
	hY := rZ*e2X - rX*e2Z
	hZ := rX*e2Y - rY*e2X

	a := e1X*hX + e1Y*hY + e1Z*hZ

	if a > -roundingError && a < roundingError {
		return false // This ray is parallel to this triangle.
	}

	f := 1.0 / a

	sX := pt1OrCamera[0] - tri.Pt0[0]
	sY := pt1OrCamera[1] - tri.Pt0[1]
	sZ := pt1OrCamera[2] - tri.Pt0[2]

	u := f * (sX*hX + sY*hY + sZ*hZ)

	if u < 0.0 || u > 1.0 {
		return false
	}

	qX := sY*e1Z - sZ*e1Y
	qY := sZ*e1X - sX*e1Z
	qZ := sX*e1Y - sY*e1X

	v := f * (rX*qX + rY*qY + rZ*qZ)

	if v < 0.0 || u+v > 1.0 {
		return false
	}

	t := f * (e2X*qX + e2Y*qY + e2Z*qZ)

	return t > roundingError && t < 1.0+roundingError
}

func (tri *Triangle) WasCutl(pt1OrCamera, pt2 *vec3.T) bool {
	intersection := tri.CaluclateIntersection(pt1OrCamera, pt2)

	if intersection == nil {
		return false
	}

	sightDirX := pt2[0] - pt1OrCamera[0]
	sightDirY := pt2[1] - pt1OrCamera[1]
	sightDirZ := pt2[2] - pt1OrCamera[2]

	cutDirX := intersection[0] - pt1OrCamera[0]
	cutDirY := intersection[1] - pt1OrCamera[1]
	cutDirZ := intersection[2] - pt1OrCamera[2]

	//watch out, cuttint < 1 may be wrong
	maxLenSq := sightDirX*sightDirX + sightDirY*sightDirY + sightDirZ*sightDirZ
	cutLenSq := cutDirX*cutDirX + cutDirY*cutDirY + cutDirZ*cutDirZ

	if cutLenSq > maxLenSq {
		return false
	}

	scalar := sightDirX*cutDirX + sightDirY*cutDirY + sightDirZ*cutDirZ

	return scalar > roundingError
}

//Practical Geometry Algorithms: with C++ Code from Amazon
// func (tri *Triangle) WasCutd(pt1OrCamera, pt2 *vec3.T) bool {
// 	dX := pt2[0] - pt1OrCamera[0]
// 	dY := pt2[1] - pt1OrCamera[1]
// 	dZ := pt2[2] - pt1OrCamera[2]
// 	l := dX*dX + dY*dY + dZ*dZ
// 	w0X := pt1OrCamera[0] - tri.Pt0[0]
// 	w0Y := pt1OrCamera[1] - tri.Pt0[1]
// 	w0Z := pt1OrCamera[2] - tri.Pt0[2]
// 	a := -(tri.intN[0]*w0X + tri.intN[1]*w0Y + tri.intN[2]*w0Z)
// 	b := tri.intN[0]*dX + tri.intN[1]*dY + tri.intN[2]*dZ
// 	if b > -roundingError && b < roundingError {
// 		return false //ray is parallel/inside triangle or disjoint
// 	}
// 	//get intersect point of ray with triangle point
// 	r := a / b
// 	if r < 0.0 {
// 		return false //no intersect
// 	}
// 	dX *= r
// 	dY *= r
// 	dZ *= r
// 	if dX*dX+dY*dY+dZ*dZ > l {
// 		return false //before
// 	}
// 	iX := pt1OrCamera[0] + dX
// 	iY := pt1OrCamera[1] + dY
// 	iZ := pt1OrCamera[2] + dZ
// 	wX := iX - tri.Pt0[0]
// 	wY := iY - tri.Pt0[1]
// 	wZ := iZ - tri.Pt0[2]
// 	wu := wX*tri.intU[0] + wY*tri.intU[1] + wZ*tri.intU[2]
// 	wv := wX*tri.intV[0] + wY*tri.intV[1] + wZ*tri.intV[2]
// 	//get and test paraetric coords
// 	s := (tri.intUV*wv - tri.intVV*wu) * tri.D
// 	if s < 0.0 || s > 1.0 { //I is outside T
// 		return false
// 	}
// 	t := (tri.intUV*wu - tri.intUU*wv) * tri.D
// 	if t < 0.0 || s+t > 1.0 { //I is outside T
// 		return false
// 	}
// 	return true
// }
