package camera

import (
	"encoding/json"
	"testing"

	"github.com/simonmau/spacial-base-calculation/dto"
	"github.com/simonmau/spacial-base-calculation/mathext"
	"github.com/stretchr/testify/assert"
)

func TestSerialization(t *testing.T) {
	cases := []struct {
		camDto CameraDto
		name   string
	}{
		{
			name: "all-set (all different values)",
			camDto: CameraDto{
				Eye:                  dto.JsonVector3{X: 0.1, Y: 0.2, Z: 0.3},
				LooksThrough:         dto.JsonVector3{X: 0.4, Y: 0.5, Z: 0.6},
				Rotation:             0.7,
				SensorWidth:          36,
				FocalLength:          50,
				ProjectionWidth:      100,
				LensCorrection:       mathext.GetPointer(20.0),
				LensCorrectionCenter: &dto.JsonVector2{X: 100, Y: 80},
				Width:                1920,
				Height:               1080,
			},
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			data, err := json.Marshal(tc.camDto)

			assert.Nil(t, err)

			strData := string(data)

			parsedCamera, err := LoadFromJson(strData)

			assert.Nil(t, err)

			assert.Equal(t, tc.camDto, *parsedCamera)
		})
	}
}

func TestErrorSerialization(t *testing.T) {
	cases := []struct {
		camDto CameraDto
		name   string
	}{
		{
			name:   "nothing-set",
			camDto: CameraDto{},
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			data, err := json.Marshal(tc.camDto)

			assert.Nil(t, err)

			strData := string(data)

			cam, err := LoadFromJson(strData)

			assert.NotNil(t, err)
			assert.Nil(t, cam)
		})
	}
}
