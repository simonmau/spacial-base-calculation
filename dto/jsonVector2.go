package dto

import "github.com/ungerik/go3d/float64/vec2"

type JsonVector2 struct {
	X float64 `json:"X" validate:"required"`
	Y float64 `json:"Y" validate:"required"`
}

func GenVec2(x, y float64) JsonVector2 {
	return JsonVector2{
		X: x,
		Y: y,
	}
}

func GenVec2FromGo3d(item *vec2.T) JsonVector2 {
	return JsonVector2{
		X: item[0],
		Y: item[1],
	}
}

func (item *JsonVector2) GenVec() vec2.T {
	return vec2.T{item.X, item.Y}
}
