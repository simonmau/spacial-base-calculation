package triangle

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ungerik/go3d/float64/vec3"
)

func TestSerialization(t *testing.T) {
	cases := []struct {
		triangle Triangle
		name     string
	}{
		{
			name: "all-set",
			triangle: GenTriangle(
				&vec3.T{0.0, 0.0, 0.0},
				&vec3.T{10.0, 0.0, 0.0},
				&vec3.T{0.0, 10.0, 0.0},
			),
		},
		{
			name: "one-unset-1",
			triangle: GenTriangle(
				nil,
				&vec3.T{10.0, 0.0, 0.0},
				&vec3.T{0.0, 10.0, 0.0},
			),
		},
		{
			name: "one-unset-2",
			triangle: GenTriangle(
				&vec3.T{10.0, 0.0, 0.0},
				nil,
				&vec3.T{0.0, 10.0, 0.0},
			),
		},
		{
			name: "one-unset-3",
			triangle: GenTriangle(
				&vec3.T{10.0, 0.0, 0.0},
				&vec3.T{0.0, 10.0, 0.0},
				nil,
			),
		},
		{
			name: "all-unset",
			triangle: GenTriangle(
				nil,
				nil,
				nil,
			),
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			data, err := json.Marshal(tc.triangle)

			assert.Nil(t, err)

			strData := string(data)

			var parsedTriangle Triangle
			err = json.Unmarshal([]byte(strData), &parsedTriangle)

			assert.Nil(t, err)

			assert.Equal(t, tc.triangle.Pt0, parsedTriangle.Pt0)
			assert.Equal(t, tc.triangle.Pt1, parsedTriangle.Pt1)
			assert.Equal(t, tc.triangle.Pt2, parsedTriangle.Pt2)
		})

	}
}
