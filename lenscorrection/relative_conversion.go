package lenscorrection

import "github.com/ungerik/go3d/float64/vec2"

func convertImageToRelativeCoordiantes(point *vec2.T, width, height *float64) vec2.T {
	return vec2.T{2.0*point[0] / *width - 1.0, 2.0*point[1] / *height - 1.0}
}

func convertRelativeToImageCoordinates(relative *vec2.T, width, height *float64) vec2.T {
	return vec2.T{(relative[0] + 1.0) * *width * 0.5, (relative[1] + 1.0) * *height * 0.5}
}
