package advancedlenscorrection

import "github.com/ungerik/go3d/float64/vec2"

func convertImageToRelativeCoordiantes(point, center *vec2.T, width, height *float64) vec2.T {
	invW := 1.0 / *width
	invH := 1.0 / *height

	posCX := center[0] * invW
	posCY := center[1] * invH

	return vec2.T{(point[0] * invW) - posCX, (point[1] * invH) - posCY}
}

func convertRelativeToImageCoordinates(relative, center *vec2.T, width, height *float64) vec2.T {
	posCX := center[0] / *width
	posCY := center[1] / *height

	return vec2.T{(relative[0] + posCX) * *width, (relative[1] + posCY) * *height}
}
