package lenscorrection

import (
	"math"

	"github.com/ungerik/go3d/float64/vec2"
)

func convertImageToRelativeCoordiantes(point *vec2.T, width, height *float64) vec2.T {
	biggerSize := math.Max(*width, *height)

	invSize := 1.0 / (biggerSize / 2.0)

	x := point[0] - (*width / 2.0)
	y := point[1] - (*height / 2.0)

	return vec2.T{x * invSize, y * invSize}
}

func convertRelativeToImageCoordinates(relative *vec2.T, width, height *float64) vec2.T {
	biggerSize := math.Max(*width, *height)

	size := biggerSize / 2.0

	x := relative[0] * size
	y := relative[1] * size

	return vec2.T{x + (*width / 2.0), y + (*height / 2.0)}
}
