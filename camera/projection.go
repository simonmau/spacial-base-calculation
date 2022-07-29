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

//WATCH OUT; POINT MUST ALREADY BE IN PLANE, USE ProjectPointIntoImage instead
func (c *Camera) ConvertPlanePointToImagePoint(pt *vec3.T) vec2.T {
	ptNew := pt.Subed(c.projectionArea.PointOnPlane)

	lx := (ptNew[0]*c.xAxis[0] + ptNew[1]*c.xAxis[1] + ptNew[2]*c.xAxis[2])
	ly := (ptNew[0]*c.yAxis[0] + ptNew[1]*c.yAxis[1] + ptNew[2]*c.yAxis[2])

	lx = lx * c.ProjectionScale
	ly = ly * c.ProjectionScale

	tmp := vec2.T{lx + c.widthOffsetCenter, ly + c.heightOffsetCenter}

	if c.LensCorrection != nil && c.LensCorrectionCenter != nil {
		tmp = *advancedlenscorrection.CorrectSexyToRaw(c.LensCorrection, &tmp, c.LensCorrectionCenter, &c.Width, &c.Height)
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
	lx = lx * c.ProjectionScaleInv

	ly := (pt[1] - c.heightOffsetCenter)
	ly = ly * c.ProjectionScaleInv

	xP := c.xAxis.Scaled(lx)
	yP := c.yAxis.Scaled(ly)

	partSum := vec3.Add(&xP, &yP)

	return vec3.Add(&partSum, c.projectionArea.PointOnPlane)
}

func (c *Camera) ProjectPointIntoImage(pt *vec3.T) vec2.T {
	intersection := c.projectionArea.CaluclateIntersection(&c.Eye, pt)
	return c.ConvertPlanePointToImagePoint(&intersection)
}
