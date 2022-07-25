package camera

import (
	"testing"

	"github.com/simonmau/spacial-base-calculation/dto"
	"github.com/stretchr/testify/assert"
	vec2 "github.com/ungerik/go3d/float64/vec2"
	"github.com/ungerik/go3d/float64/vec3"
)

func TestSimpleProjectionFrom3To2(t *testing.T) {
	camDto := CameraDto{
		Eye:                  dto.JsonVector3{X: 0, Y: 0, Z: 200},
		LooksThrough:         dto.JsonVector3{X: 0, Y: 0, Z: 0},
		Rotation:             0,
		SensorWidth:          36,
		FocalLength:          50,
		ProjectionWidth:      100,
		LensCorrection:       nil,
		LensCorrectionCenter: nil,
		Width:                1920,
		Height:               1080,
	}

	cam := FromDto(&camDto)

	projectedPoint := cam.ProjectPointIntoImage(&vec3.T{0, 0, 0})

	assert.InDelta(t, 1920/2, projectedPoint[0], 0.00001)
	assert.InDelta(t, 1080/2, projectedPoint[1], 0.00001)
}

func TestSimpleProjectionFrom2To3(t *testing.T) {
	camDto := CameraDto{
		Eye:                  dto.JsonVector3{X: 0, Y: 0, Z: 200},
		LooksThrough:         dto.JsonVector3{X: 0, Y: 0, Z: 0},
		Rotation:             0,
		SensorWidth:          36,
		FocalLength:          50,
		ProjectionWidth:      100,
		LensCorrection:       nil,
		LensCorrectionCenter: nil,
		Width:                1920,
		Height:               1080,
	}

	cam := FromDto(&camDto)

	projectedPoint := cam.ConvertImagePointToPlanePoint(&vec2.T{0, 0})

	assert.InDelta(t, -50, projectedPoint[0], 0.00001)
	assert.InDelta(t, -(1080.0*100.0/1920.0)/2.0, projectedPoint[1], 0.00001)

	backConverted := cam.ConvertPlanePointToImagePoint(&projectedPoint)

	assert.InDelta(t, 0, backConverted[0], 0.00001)
	assert.InDelta(t, 0, backConverted[1], 0.00001)
}

func TestProjection(t *testing.T) {
	camDto := CameraDto{
		Eye:                  dto.JsonVector3{X: 20, Y: 20, Z: 200},
		LooksThrough:         dto.JsonVector3{X: 0, Y: 0, Z: 0},
		Rotation:             0,
		SensorWidth:          36,
		FocalLength:          50,
		ProjectionWidth:      100,
		LensCorrection:       nil,
		LensCorrectionCenter: nil,
		Width:                1920,
		Height:               1080,
	}

	cam := FromDto(&camDto)

	projectedPoint := cam.ConvertImagePointToPlanePoint(&vec2.T{1920 / 2, 1080 / 2})

	backConverted := cam.ConvertPlanePointToImagePoint(&projectedPoint)

	assert.InDelta(t, 1920/2, backConverted[0], 0.00001)
	assert.InDelta(t, 1080/2, backConverted[1], 0.00001)
}
