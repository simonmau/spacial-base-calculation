package lenscorrection

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ungerik/go3d/float64/vec2"
)

func TestModestLensCorrection_1(t *testing.T) {
	src := vec2.T{-1, -1}

	dst := modestLensCorrection(-0.2, &src)

	assert.InDelta(t, -0.8, dst[0], 0.2)
	assert.InDelta(t, -0.8, dst[1], 0.2)
}

func TestModestLensCorrection_Full(t *testing.T) {
	src := vec2.T{-1, 1}

	p := -0.2

	dst := modestLensCorrection(p, &src)

	back := modestLensCorrection(-p, &dst)

	assert.InDelta(t, src[0], back[0], 0.001)
	assert.InDelta(t, src[1], back[1], 0.001)
}
