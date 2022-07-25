package dto

import "github.com/ungerik/go3d/float64/vec3"

type JsonVector3 struct {
	X float64 `json:"X" validate:"required"`
	Y float64 `json:"Y" validate:"required"`
	Z float64 `json:"Z" validate:"required"`
}

func GenVec3(x, y, z float64) JsonVector3 {
	return JsonVector3{
		X: x,
		Y: y,
		Z: z,
	}
}

func (item *JsonVector3) GenVec() vec3.T {
	return vec3.T{item.X, item.Y, item.Z}
}
