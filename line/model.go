package line

import "github.com/ungerik/go3d/float64/vec3"

type Line struct {
	Pt0 *vec3.T
	Pt1 *vec3.T
}

func GenerateWithLength(pt0, dir *vec3.T, length float64) *Line {
	pt1 := vec3.Add(pt0, dir.Scale(length))

	return &Line{
		Pt0: pt0,
		Pt1: &pt1,
	}
}

func GenerateWithTwoPoints(pt0, pt1 *vec3.T) *Line {
	return &Line{
		Pt0: pt0,
		Pt1: pt1,
	}
}
