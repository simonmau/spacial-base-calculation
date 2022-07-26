package test

import (
	"testing"

	"github.com/simonmau/spacial-base-calculation/camera"
	"github.com/simonmau/spacial-base-calculation/dto"
	"github.com/stretchr/testify/assert"
	"github.com/ungerik/go3d/float64/vec2"
)

func TestFullTransformation(t *testing.T) {
	srcPt := vec2.T{26.166, 168.76}

	//#region GENERATE CAMERA

	camDto := camera.CameraDto{
		Eye:             dto.JsonVector3{X: 0.0, Y: 487.053, Z: 350.3981},
		LooksThrough:    dto.JsonVector3{X: 0.0, Y: 30.0, Z: 0.0},
		SensorWidth:     36,
		FocalLength:     247.9126,
		ProjectionWidth: 100,
		Rotation:        0,
		Width:           640,
		Height:          360,
	}

	cam := camera.FromDto(&camDto)

	//#endregion GENERATE CAMERA

	//#region CALC TRANSFORMATIONS

	spacePt := cam.ConvertImagePointToPlanePoint(&srcPt)

	dstPt := cam.ProjectPointIntoImage(&spacePt)

	assert.InDelta(t, srcPt[0], dstPt[0], 0.0001)
	assert.InDelta(t, srcPt[1], dstPt[1], 0.0001)

	//#endregion CALC TRANSFORMATIONS
}
