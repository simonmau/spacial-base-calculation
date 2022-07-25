package triangle

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ungerik/go3d/float64/vec3"
)

func TestWasCut(t *testing.T) {
	cases := []struct {
		Pt1    *vec3.T
		Pt2    *vec3.T
		WasCut bool
	}{
		{
			Pt1:    &vec3.T{1.0, 1.0, 1.0},
			Pt2:    &vec3.T{1.0, 1.0, -1.0},
			WasCut: true,
		},
		{
			Pt1:    &vec3.T{-1.0, 1.0, 1.0},
			Pt2:    &vec3.T{-1.0, 1.0, -1.0},
			WasCut: false,
		},
		{
			Pt1:    &vec3.T{10.0, 10.0, 1.0},
			Pt2:    &vec3.T{10.0, 10.0, -1.0},
			WasCut: false,
		},
	}

	triangle := GenTriangle(&vec3.T{0.0, 0.0, 0.0},
		&vec3.T{10.0, 0.0, 0.0},
		&vec3.T{0.0, 10.0, 0.0},
	)

	for _, tc := range cases {
		result := triangle.WasCut(tc.Pt1, tc.Pt2)
		assert.Equal(t, tc.WasCut, result)
	}
}

func TestWasCut_Close(t *testing.T) {
	Pt1 := vec3.T{1.0, 1.0, 1.0}

	triangle := GenTriangle(&vec3.T{0.0, 0.0, 0.0},
		&vec3.T{10.0, 0.0, 0.0},
		&vec3.T{0.0, 10.0, 0.0},
	)

	Pt2 := vec3.T{1.0, 1.0, -1.0}
	result := triangle.WasCut(&Pt1, &Pt2)
	assert.Equal(t, true, result)

	Pt2 = vec3.T{1.0, 1.0, -0.1}
	result = triangle.WasCut(&Pt1, &Pt2)
	assert.Equal(t, true, result)

	Pt2 = vec3.T{1.0, 1.0, 0.1}
	result = triangle.WasCut(&Pt1, &Pt2)
	assert.Equal(t, false, result)
}

func BenchmarkWasCut_Close(b *testing.B) {
	Pt1 := vec3.T{1.0, 1.0, 1.0}

	triangle := GenTriangle(&vec3.T{0.0, 0.0, 0.0},
		&vec3.T{10.0, 0.0, 0.0},
		&vec3.T{0.0, 10.0, 0.0},
	)

	Pt2 := vec3.T{1.0, 1.0, -1.0}

	for i := 0; i < b.N; i++ {
		triangle.WasCut(&Pt1, &Pt2)
	}

	result := triangle.WasCut(&Pt1, &Pt2)
	assert.Equal(b, true, result)

	Pt2 = vec3.T{1.0, 1.0, -0.1}
	result = triangle.WasCut(&Pt1, &Pt2)
	assert.Equal(b, true, result)

	Pt2 = vec3.T{1.0, 1.0, 0.1}
	result = triangle.WasCut(&Pt1, &Pt2)
	assert.Equal(b, false, result)
}
