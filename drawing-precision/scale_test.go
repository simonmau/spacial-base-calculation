package drawingprecision

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ungerik/go3d/float64/vec3"
)

func TestScale_1(t *testing.T) {
	dp := GenDrawPrecByDistance(0.1)
	dir := vec3.T{0, 0, 1}
	dp.ScaleDirection(&dir)

	assert.InDelta(t, 0.1, dir[2], 0.0001)
}

func TestScale_2(t *testing.T) {
	dp := GenDrawPrecByPointsPerUnit(10)
	dir := vec3.T{0, 0, 1}
	dp.ScaleDirection(&dir)

	assert.InDelta(t, 0.1, dir[2], 0.0001)
}
