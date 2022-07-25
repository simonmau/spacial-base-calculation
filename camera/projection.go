package camera

import (
	advancedlenscorrection "github.com/simonmau/spacial-base-calculation/advanced-lenscorrection"
	"github.com/simonmau/spacial-base-calculation/lenscorrection"
	vec2 "github.com/ungerik/go3d/float64/vec2"
	vec3 "github.com/ungerik/go3d/float64/vec3"
)

func (c *Camera) ProjectPointIntoPlane(pt *vec3.T) vec3.T {
	return c.projectionArea.CaluclateIntersection(&c.Eye, pt)
}

func (c *Camera) ConvertPlanePointToImagePoint(pt *vec3.T) vec2.T {
	correctedPointInPlaneX := pt[0] - c.Eye[0]
	correctedPointInPlaneY := pt[1] - c.Eye[1]
	correctedPointInPlaneZ := pt[2] - c.Eye[2]

	lx := (correctedPointInPlaneX*c.xAxis[0] + correctedPointInPlaneY*c.xAxis[1] + correctedPointInPlaneZ*c.xAxis[2])
	ly := (correctedPointInPlaneX*c.yAxis[0] + correctedPointInPlaneY*c.yAxis[1] + correctedPointInPlaneZ*c.yAxis[2])

	tmp := vec2.T{lx + c.widthOffsetCenter, ly + c.heightOffsetCenter}

	if c.LensCorrection != nil && c.LensCorrectionCenter != nil {
		tmp = *advancedlenscorrection.CorrectRawToSexy(c.LensCorrection, &tmp, c.LensCorrectionCenter, &c.Width, &c.Height)
	} else if c.LensCorrection != nil {
		tmp = *lenscorrection.CorrectSexyToRaw(c.LensCorrection, &tmp, &c.Width, &c.Height)
	}

	return tmp
}

func (c *Camera) ConvertImagePointToPlanePoint(raw *vec2.T) vec3.T {
	var pt vec2.T

	if c.LensCorrection != nil && c.LensCorrectionCenter != nil {
		pt = *advancedlenscorrection.CorrectRawToSexy(c.LensCorrection, raw, c.LensCorrectionCenter, &c.Width, &c.Height)
	} else if c.LensCorrection != nil {
		pt = *lenscorrection.CorrectRawToSexy(c.LensCorrection, raw, &c.Width, &c.Height)
	} else {
		pt = *raw
	}

	lx := (pt[0] - c.widthOffsetCenter)
	ly := (pt[1] - c.heightOffsetCenter)

	xP := c.xAxis.Scaled(lx)
	yP := c.yAxis.Scaled(ly)

	partSum := vec3.Add(&xP, &yP)

	return vec3.Add(&partSum, &c.LooksThrough)
}

func (c *Camera) ProjectPointIntoImage(pt *vec3.T) vec2.T {
	intersection := c.projectionArea.CaluclateIntersection(&c.Eye, pt)
	return c.ConvertPlanePointToImagePoint(&intersection)
}
