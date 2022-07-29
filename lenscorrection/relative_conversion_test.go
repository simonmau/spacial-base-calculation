package lenscorrection

import (
	"testing"

	"github.com/simonmau/spacial-base-calculation/mathext"
	rangegen "github.com/simonmau/spacial-base-calculation/range-gen"
	"github.com/stretchr/testify/assert"
	"github.com/ungerik/go3d/float64/vec2"
)

func TestCoordinateConversion(t *testing.T) {
	bigRange := rangegen.GenLinearRangeArray(mathext.GetPointer(-4000.0), mathext.GetPointer(4000.0), mathext.GetPointer(1000.0))

	w := 1920.0
	h := 1080.0

	for _, x := range bigRange {
		for _, y := range bigRange {
			pt := vec2.T{x, y}

			relPt := convertImageToRelativeCoordiantes(&pt, &w, &h)

			newPt := convertRelativeToImageCoordinates(&relPt, &w, &h)

			assert.InDelta(t, pt[0], newPt[0], 0.00001)
			assert.InDelta(t, pt[1], newPt[1], 0.00001)
		}
	}
}

func TestCoordinateConversion_TopLeft(t *testing.T) {
	w := 1920.0
	h := 1080.0

	pt := vec2.T{0, 0}

	relPt := convertImageToRelativeCoordiantes(&pt, &w, &h)

	assert.InDelta(t, -1.0, relPt[0], 0.00001)
	assert.InDelta(t, -(1080.0 / 1920.0), relPt[1], 0.00001)

	newPt := convertRelativeToImageCoordinates(&relPt, &w, &h)

	assert.InDelta(t, pt[0], newPt[0], 0.00001)
	assert.InDelta(t, pt[1], newPt[1], 0.00001)
}

func TestCoordinateConversion_BottomRight(t *testing.T) {
	w := 1920.0
	h := 1080.0

	pt := vec2.T{w, h}

	relPt := convertImageToRelativeCoordiantes(&pt, &w, &h)

	assert.InDelta(t, 1.0, relPt[0], 0.00001)
	assert.InDelta(t, 1080.0/1920.0, relPt[1], 0.00001)

	newPt := convertRelativeToImageCoordinates(&relPt, &w, &h)

	assert.InDelta(t, pt[0], newPt[0], 0.00001)
	assert.InDelta(t, pt[1], newPt[1], 0.00001)
}
