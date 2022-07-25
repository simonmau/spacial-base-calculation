package camera

import (
	"math"

	"github.com/simonmau/spacial-base-calculation/mathext"
	"github.com/simonmau/spacial-base-calculation/plane"
	vec2 "github.com/ungerik/go3d/float64/vec2"
	vec3 "github.com/ungerik/go3d/float64/vec3"
)

type Camera struct {
	//where the camera is positioned
	Eye vec3.T

	//the sight-direction will be calculated starting from the eye - just any point on sight
	LooksThrough vec3.T

	//rotation around the sight-axis (lt - eye)
	Rotation float64

	//a sensor size can be set here, changes will result in zoom effect, default should be 36mm for 'vollformat'
	SensorWidth float64

	//the zoom-length of the camera, will move the projection-layer from the camera back and forwards
	//with sensor-width 36mm should be at least 30mm (extremely-wide) to max 200mm (extreme-zoom)
	FocalLength float64

	//see https://www.fotomagazin.de/technik/der-bildwinkel
	ViewingAngleWidth float64

	LToProjectionArea float64

	//the projection-width is the width of the projection-area where the image in px will be projected on
	//this should be in normal cases 100mm
	//with a fixed projection-width, every width/height px produces the same image, just less sharp (or more)
	ProjectionWidth float64

	//lenscorrection warps the image before/after the necessary transformations are finished, so its on top of everything else
	//can be nil -> no lenscorrection will be applied
	LensCorrection *float64

	//can be nil when lenscorrection is set -> will then be width/2 and height/2
	LensCorrectionCenter *vec2.T

	Width  float64
	Height float64

	widthOffsetCenter  float64
	heightOffsetCenter float64

	projectionArea plane.Plane
	sightDirection vec3.T

	xAxis vec3.T
	yAxis vec3.T
}

func FromDto(dto *CameraDto) *Camera {
	w := float64(dto.Width)
	h := float64(dto.Height)

	c := Camera{
		Eye:          dto.Eye.GenVec(),
		LooksThrough: dto.LooksThrough.GenVec(),

		Rotation:       dto.Rotation,
		LensCorrection: dto.LensCorrection,

		Width:              w,
		Height:             h,
		widthOffsetCenter:  w / 2.0,
		heightOffsetCenter: h / 2.0,

		ViewingAngleWidth: 2.0 * math.Atan(dto.SensorWidth/(2.0*dto.FocalLength)),
	}

	if c.LensCorrection != nil && dto.LensCorrectionCenter != nil {
		tmp := dto.LensCorrectionCenter.GenVec()
		c.LensCorrectionCenter = &tmp
	}

	//fill sight-direction
	sight := vec3.Sub(&c.LooksThrough, &c.Eye)
	c.sightDirection = sight.Normalized()

	c.LToProjectionArea = (dto.ProjectionWidth / 2.0) / math.Tan(c.ViewingAngleWidth/2.0)

	pointOnProjectionPlane := c.sightDirection.Scaled(c.LToProjectionArea / c.sightDirection.Length())

	//fill plane
	c.projectionArea = plane.Plane{
		PointOnPlane: &pointOnProjectionPlane,
		NormalVector: &c.sightDirection,
	}

	//y axis
	intersection := c.projectionArea.CaluclateIntersection(&c.Eye, &vec3.T{0.0, 0.0, -1.0})

	yAxis := vec3.Sub(&intersection, &c.LooksThrough)
	yAxis.Normalize()

	///rotate, xaxis will be rotated automatically as it depends on yAxis
	c.yAxis = *mathext.RotateVecAxis(&yAxis, &c.sightDirection, &c.Rotation)

	//x axis
	xAxis := vec3.Cross(&c.yAxis, &c.sightDirection)
	xAxis.Normalize()

	c.xAxis = xAxis

	return &c
}
