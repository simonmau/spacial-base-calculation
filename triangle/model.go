package triangle

import "github.com/ungerik/go3d/float64/vec3"

type Triangle struct {
	Pt0 *vec3.T
	Pt1 *vec3.T
	Pt2 *vec3.T
}

func GenTriangle(pt0 *vec3.T, pt1 *vec3.T, pt2 *vec3.T) Triangle {
	//u := vec3.Sub(pt1, pt0)
	//v := vec3.Sub(pt2, pt0)
	//
	//n := vec3.Cross(&u, &v)
	//
	//uu := u[0]*u[0] + u[1]*u[1] + u[2]*u[2]
	//uv := u[0]*v[0] + u[1]*v[1] + u[2]*v[2]
	//vv := v[0]*v[0] + v[1]*v[1] + v[2]*v[2]

	return Triangle{
		Pt0: pt0,
		Pt1: pt1,
		Pt2: pt2,
	}

	//	intU:  u,
	//	intV:  v,
	//	intN:  n,
	//	intUU: uu,
	//	intUV: uv,
	//	intVV: vv,
	//	D:     1.0 / (uv*uv - uu*vv),
}
