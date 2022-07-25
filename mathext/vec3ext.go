package mathext

import (
	"math"

	"github.com/ungerik/go3d/float64/vec3"
)

func GenVec(x, y, z float64) *vec3.T {
	var v vec3.T

	v[0] = x
	v[1] = y
	v[2] = z

	return &v
}

func RotateVec(v *vec3.T, angle *float64) vec3.T {
	sin, cos := math.Sincos(*angle)

	x := cos*v[0] - sin*v[1]
	y := sin*v[0] + cos*v[1]

	return vec3.T{x, y, v[2]}
}

func RotateVecAxis(v, axis *vec3.T, angle *float64) *vec3.T {
	sin, cos := math.Sincos(*angle)

	klammer := 1.0 - cos
	n1 := axis[0]
	n2 := axis[1]
	n3 := axis[2]
	n1Sq := n1 * n1
	n2Sq := n2 * n2
	n3Sq := n3 * n3

	x := v[0]*(n1Sq*klammer+cos) + v[1]*(n1*n2*klammer-n3*sin) + v[2]*(n1*n3*klammer+n2*sin)
	y := v[0]*(n2*n1*klammer+n3*sin) + v[1]*(n2Sq*klammer+cos) + v[2]*(n2*n3*klammer-n1*sin)
	z := v[0]*(n3*n1*klammer-n2*sin) + v[1]*(n3*n2*klammer+n1*sin) + v[2]*(n3Sq*klammer+cos)

	return GenVec(x, y, z)
}
