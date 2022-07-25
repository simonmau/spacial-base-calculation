package line

import (
	"testing"

	drawingprecision "github.com/simonmau/spacial-base-calculation/drawing-precision"
	"github.com/stretchr/testify/assert"
	"github.com/ungerik/go3d/float64/vec3"
)

func TestEdge_1(t *testing.T) {
	line := GenerateWithTwoPoints(&vec3.Zero, &vec3.UnitXYZ)

	i := 0

	res := make([]vec3.T, 0)

	dp := drawingprecision.GenDrawPrecByPointsPerUnit(1)

	line.GenerateEdgePoints(dp, &res)

	for _, pt := range res {
		if i == 0 {
			assert.Equal(t, 0.0, pt[0])
			assert.Equal(t, 0.0, pt[1])
			assert.Equal(t, 0.0, pt[2])
		}

		if i == 1 {
			assert.Equal(t, 0.5773502691896258, pt[0])
			assert.Equal(t, 0.5773502691896258, pt[1])
			assert.Equal(t, 0.5773502691896258, pt[2])

			assert.Equal(t, 1.0, vec3.Distance(line.Pt0, &pt))
		}

		if i == 2 {
			assert.Equal(t, 1.0, pt[0])
			assert.Equal(t, 1.0, pt[1])
			assert.Equal(t, 1.0, pt[2])
		}

		i++
	}

	assert.Equal(t, 3, i)
}

func TestEdge_2(t *testing.T) {
	line := GenerateWithTwoPoints(&vec3.Zero, &vec3.UnitXYZ)

	i := 0

	res := make([]vec3.T, 0)

	dp := drawingprecision.GenDrawPrecByDistance(0.000001)

	line.GenerateEdgePoints(dp, &res)

	for _, pt := range res {
		assert.NotNil(t, pt)
		i++
	}

	assert.Equal(t, 1732052, i)
}

func BenchmarkEdge(t *testing.B) {
	t.StopTimer()

	line := GenerateWithTwoPoints(&vec3.Zero, &vec3.UnitXYZ)

	i := 0

	res := make([]vec3.T, 0)

	dp := drawingprecision.GenDrawPrecByDistance(0.000001)

	line.GenerateEdgePoints(dp, &res)

	for _, pt := range res {
		assert.NotNil(t, pt)
		i++
	}

	assert.Equal(t, 1732052, i)

	t.StartTimer()

	for i := 0; i < t.N; i++ {
		res := make([]vec3.T, 0)
		t.StartTimer()
		line.GenerateEdgePoints(dp, &res)
		t.StopTimer()
		assert.Equal(t, 1732052, len(res))
	}
}
