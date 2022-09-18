package ellipse

import (
	"github.com/simonmau/spacial-base-calculation/dto"
	"github.com/ungerik/go3d/float64/vec2"
)

type EllipseDto struct {
	EllipseData *Data              `json:"ellipseData" validate:"required"`
	Points      *[]dto.JsonVector2 `json:"points" validate:"required"`
}

func (data *EllipseDto) GetPoints() []vec2.T {
	res := make([]vec2.T, 0, len(*data.Points))

	for _, item := range *data.Points {
		res = append(res, item.GenVec())
	}

	return res
}
