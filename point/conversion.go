package point

import (
	"github.com/ungerik/go3d/float64/vec2"
)

func FromVector(pt *vec2.T) *T {
	return &T{int32(pt[0]), int32(pt[1])}
}

func (t *T) ToIndex(width *int32, channels int32) int32 {
	return (t[1]**width + t[0]) * channels
}

func (t *T) InRange(width, height *int32) bool {
	return t[0] >= 0 && t[0] < *width && t[1] >= 0 && t[1] < *height
}
