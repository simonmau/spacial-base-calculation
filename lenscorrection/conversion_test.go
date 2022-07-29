package lenscorrection

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ungerik/go3d/float64/vec2"
)

func TestConversions(t *testing.T) {
	testCases := []struct {
		P     float64
		Point vec2.T
	}{
		{
			P:     0.31,
			Point: vec2.T{302, 483},
		},
		{
			P:     0.31,
			Point: vec2.T{0, 0},
		},
		{
			P:     1,
			Point: vec2.T{0, 0},
		},
		{
			P:     0.51,
			Point: vec2.T{402, 133},
		},
		{
			P:     2.11,
			Point: vec2.T{851, 591},
		},
		{
			P:     2.0,
			Point: vec2.T{100, 100},
		},
	}

	width := float64(1920)
	height := float64(1080)

	for i, item := range testCases {

		item := item //copy the value local for the parallel testing

		t.Run("tc "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			sexyPt := CorrectRawToSexy(&item.P, &item.Point, &width, &height)

			assert.NotNil(t, sexyPt)

			if sexyPt != nil {
				imgPt := CorrectSexyToRaw(&item.P, sexyPt, &width, &height)

				assert.NotNil(t, imgPt)

				assert.InDelta(t, item.Point[0], imgPt[0], 0.001)
				assert.InDelta(t, item.Point[1], imgPt[1], 0.001)
			}
		})
	}
}

func TestConversions3(t *testing.T) {
	testCases := []struct {
		P     float64
		Point vec2.T
	}{
		{
			P:     3.0,
			Point: vec2.T{100, 100},
		},
	}

	width := float64(1920)
	height := float64(1080)

	for i, item := range testCases {

		item := item //copy the value local for the parallel testing

		t.Run("tc "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			sexyPt := CorrectRawToSexy(&item.P, &item.Point, &width, &height)

			assert.NotNil(t, sexyPt)

			if sexyPt != nil {
				imgPt := CorrectSexyToRaw(&item.P, sexyPt, &width, &height)

				assert.NotNil(t, imgPt)

				assert.InDelta(t, item.Point[0], imgPt[0], 0.001)
				assert.InDelta(t, item.Point[1], imgPt[1], 0.001)
			}
		})
	}
}
