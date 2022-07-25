package advancedlenscorrection

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/ungerik/go3d/float64/vec2"
)

func TestConversions(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		P      float64
		Point  vec2.T
		Center vec2.T
	}{
		{
			P:      0.31,
			Point:  vec2.T{302, 483},
			Center: vec2.T{860, 540},
		},
		{
			P:      0.31,
			Point:  vec2.T{302, 483},
			Center: vec2.T{10, 10},
		},
		{
			P:      -0.31,
			Point:  vec2.T{302, 483},
			Center: vec2.T{10, 10},
		},
		{
			P:      0.31,
			Point:  vec2.T{0, 0},
			Center: vec2.T{860, 540},
		},
		{
			P:      1,
			Point:  vec2.T{0, 0},
			Center: vec2.T{860, 540},
		},
		{
			P:      0.51,
			Point:  vec2.T{402, 133},
			Center: vec2.T{860, 540},
		},
		{
			P:      -0.51,
			Point:  vec2.T{102, 137},
			Center: vec2.T{860, 540},
		},
		{
			P:      -2.11,
			Point:  vec2.T{851, 591},
			Center: vec2.T{860, 540},
		},
		{
			P:      2.11,
			Point:  vec2.T{851, 591},
			Center: vec2.T{860, 540},
		},
	}

	width := float64(1920)
	height := float64(1080)

	for i, item := range testCases {

		item := item //copy the value local for the parallel testing

		t.Run("tc "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel() //start parallel execution

			sexyPt := CorrectRawToSexy(&item.P, &item.Point, &item.Center, &width, &height)

			if sexyPt != nil {
				imgPt := CorrectSexyToRaw(&item.P, sexyPt, &item.Center, &width, &height)

				assert.NotNil(t, imgPt)

				assert.InDelta(t, item.Point[0], imgPt[0], 0.001)
				assert.InDelta(t, item.Point[1], imgPt[1], 0.001)
			}

			time.Sleep(time.Second)
		})
	}
}
