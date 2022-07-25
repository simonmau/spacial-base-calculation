package triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ungerik/go3d/float64/vec3"
)

func TestSimpleExample_1(t *testing.T) {
	tri := Triangle{Pt0: &vec3.T{0, 3, 0}, Pt1: &vec3.T{-0.5, 3, -1}, Pt2: &vec3.T{0.5, 3, -1}}

	cut := tri.CaluclateIntersection(&vec3.T{0, 2, -1}, &vec3.T{0, 3, -1})

	assert.NotNil(t, cut)

	if cut != nil {
		assert.Equal(t, float64(0.0), cut[0])
		assert.Equal(t, float64(3.0), cut[1])
		assert.Equal(t, float64(-1.0), cut[2])
	}
}

func TestIntersection(t *testing.T) {
	cases := []struct {
		Pt1   *vec3.T
		Pt2   *vec3.T
		CutPt *vec3.T
	}{
		{
			Pt1:   &vec3.T{1.0, 1.0, 1.0},
			Pt2:   &vec3.T{1.0, 1.0, -1.0},
			CutPt: &vec3.T{1.0, 1.0, 0.0},
		},
		{
			Pt1:   &vec3.T{-1.0, 1.0, 1.0},
			Pt2:   &vec3.T{-1.0, 1.0, -1.0},
			CutPt: nil,
		},
		{
			Pt1:   &vec3.T{10.0, 10.0, 1.0},
			Pt2:   &vec3.T{10.0, 10.0, -1.0},
			CutPt: nil,
		},
	}

	triangle := GenTriangle(&vec3.T{0.0, 0.0, 0.0},
		&vec3.T{10.0, 0.0, 0.0},
		&vec3.T{0.0, 10.0, 0.0},
	)

	for _, tc := range cases {
		result := triangle.CaluclateIntersection(tc.Pt1, tc.Pt2)

		if result == nil {
			assert.Nil(t, tc.CutPt)
		} else {
			assert.Equal(t, result[0], tc.CutPt[0])
			assert.Equal(t, result[1], tc.CutPt[1])
			assert.Equal(t, result[2], tc.CutPt[2])
		}
	}
}
