package dto

import "github.com/ungerik/go3d/float64/vec3"

type JsonVector3 struct {
	X float64 `json:"x" validate:"required"`
	Y float64 `json:"y" validate:"required"`
	Z float64 `json:"z" validate:"required"`
}

func GenVec3(x, y, z float64) JsonVector3 {
	return JsonVector3{
		X: x,
		Y: y,
		Z: z,
	}
}

func GenVec3FromGo3d(item *vec3.T) JsonVector3 {
	return JsonVector3{
		X: item[0],
		Y: item[1],
		Z: item[2],
	}
}

func (item *JsonVector3) GenVec() vec3.T {
	return vec3.T{item.X, item.Y, item.Z}
}
