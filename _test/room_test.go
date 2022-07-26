package test

import (
	"math"
	"testing"
	"time"

	"github.com/simonmau/spacial-base-calculation/camera"
	drawingprecision "github.com/simonmau/spacial-base-calculation/drawing-precision"
	"github.com/simonmau/spacial-base-calculation/dto"
	"github.com/simonmau/spacial-base-calculation/line"
	"github.com/simonmau/spacial-base-calculation/mathext"
	"github.com/simonmau/spacial-base-calculation/point"
	rangegen "github.com/simonmau/spacial-base-calculation/range-gen"
	"github.com/simonmau/spacial-base-calculation/triangle"
	"github.com/ungerik/go3d/float64/vec3"
)

func TestRoom(t *testing.T) {
	//#region GENERATE GRAPHIC FOR DRAWING

	tmp := (20.0 / (math.Sqrt(3) / 2)) / 2 //gleichseitiges dreieck - halbe höhe

	t1 := triangle.GenTriangle(&vec3.T{10, 0, 0}, &vec3.T{-10, tmp, 0}, &vec3.T{-10, -tmp, 0})
	t2 := triangle.GenTriangle(&vec3.T{-10, 0, 10}, &vec3.T{10, tmp, 10}, &vec3.T{10, -tmp, 10})

	c := t1.Pt1.Added(t1.Pt2)
	c.Scale(0.5)

	l1 := line.GenerateWithTwoPoints(t1.Pt0, &c)

	dp := drawingprecision.GenDrawPrecByPointsPerUnit(10)

	edgePoints := make([]vec3.T, 0)

	l1.GenerateEdgePoints(dp, &edgePoints)
	t1.GenerateEdgePoints(dp, &edgePoints)
	t2.GenerateEdgePoints(dp, &edgePoints)

	//#endregion GENERATE GRAPHIC FOR DRAWING

	//#region GENERATE CAMERA

	minY := 0.0
	maxY := 800.0
	steps := 20.0

	yRange := rangegen.GenLinearRangeArray(&minY, &maxY, &steps)

	for _, y := range yRange {
		eyeVec := vec3.T{0, y, 500}
		eyeVec.Normalize()
		eyeVec.Scale(500.0)

		camDto := camera.CameraDto{
			Eye:             dto.JsonVector3{X: eyeVec[0], Y: eyeVec[1], Z: eyeVec[2]},
			LooksThrough:    dto.JsonVector3{X: 0.0, Y: 0.0, Z: 0.0},
			SensorWidth:     36,
			FocalLength:     70,
			ProjectionWidth: 100,
			Rotation:        0,
			Width:           1920,
			Height:          1080,
			LensCorrection:  mathext.GetPointer(1.0),
		}

		cam := camera.FromDto(&camDto)

		//#endregion GENERATE CAMERA

		//#region DRAW IMAGE

		img := make([]byte, camDto.Width*camDto.Height*3)

		for _, spacePt := range edgePoints {
			imgPtFloat := cam.ProjectPointIntoImage(&spacePt)
			imgPt := point.FromVector(&imgPtFloat)

			if imgPt.InRange(&camDto.Width, &camDto.Height) {
				index := imgPt.ToIndex(&camDto.Width, 3)

				img[index] = 255
				img[index+1] = 0
				img[index+2] = 255
			}
		}

		rgbaImage := byteToRgbImage(&img, int(camDto.Width), int(camDto.Height))

		saveImageRgba("output.png", rgbaImage)

		time.Sleep(time.Millisecond * 250)
	}

	//#endregion DRAW IMAGE
}
