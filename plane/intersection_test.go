package plane

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ungerik/go3d/float64/vec3"
)

func TestSimpleExample_1(t *testing.T) {
	plane := Plane{PointOnPlane: &vec3.T{0, 0, 0}, NormalVector: &vec3.T{0, 0, 1}}

	cut := plane.CaluclateIntersection(&vec3.T{100, 100, 100}, &vec3.T{10, 10, 10})

	assert.InDelta(t, 0.0, cut[2], 0.0000001)
}

func BenchmarkSimpleExample_1(t *testing.B) {
	plane := Plane{PointOnPlane: &vec3.T{0, 0, 0}, NormalVector: &vec3.T{0, 0, 1}}

	cut := plane.CaluclateIntersection(&vec3.T{100, 100, 100}, &vec3.T{10, 10, 10})
	assert.InDelta(t, 0.0, cut[2], 0.0000001)

	for i := 0; i < t.N; i++ {
		_ = plane.CaluclateIntersection(&vec3.T{100, 100, 100}, &vec3.T{10, 10, 10})
	}
}

func TestSimpleExample_2(t *testing.T) {
	plane := Plane{PointOnPlane: &vec3.T{0, 0, 0}, NormalVector: &vec3.T{0, 0, 1}}

	cut := plane.CaluclateIntersection(&vec3.T{77.41, 40.96, 116.25}, &vec3.T{-4.466, -30.59, 90.09})

	assert.InDelta(t, 0.0, cut[2], 0.0000001)
}

func BenchmarkSimpleExample_2(t *testing.B) {
	plane := Plane{PointOnPlane: &vec3.T{0, 0, 0}, NormalVector: &vec3.T{0, 0, 1}}

	cut := plane.CaluclateIntersection(&vec3.T{77.41, 40.96, 116.25}, &vec3.T{-4.466, -30.59, 90.09})

	assert.InDelta(t, 0.0, cut[2], 0.0000001)
}
