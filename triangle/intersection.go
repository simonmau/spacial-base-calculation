package triangle

import (
	"math"

	"github.com/ungerik/go3d/float64/vec3"
)

func (tri *Triangle) CaluclateIntersection(pt1, pt2 *vec3.T) *vec3.T {
	dX := pt2[0] - pt1[0]
	dY := pt2[1] - pt1[1]
	dZ := pt2[2] - pt1[2]

	p0p1X := tri.Pt1[0] - tri.Pt0[0]
	p0p1Y := tri.Pt1[1] - tri.Pt0[1]
	p0p1Z := tri.Pt1[2] - tri.Pt0[2]

	p0p2X := tri.Pt2[0] - tri.Pt0[0]
	p0p2Y := tri.Pt2[1] - tri.Pt0[1]
	p0p2Z := tri.Pt2[2] - tri.Pt0[2]

	nX := p0p1Y*p0p2Z - p0p1Z*p0p2Y
	nY := p0p1Z*p0p2X - p0p1X*p0p2Z
	nZ := p0p1X*p0p2Y - p0p1Y*p0p2X

	nDotRayDirection := nX*dX + nY*dY + nZ*dZ

	if math.Abs(nDotRayDirection) < roundingError {
		return nil
	}

	d := nX*tri.Pt0[0] + nY*tri.Pt0[1] + nZ*tri.Pt0[2]

	tmp := nX*pt1[0] + nY*pt1[1] + nZ*pt1[2]

	t := -((tmp - d) / nDotRayDirection)

	dX *= t
	dY *= t
	dZ *= t

	pX := pt1[0] + dX
	pY := pt1[1] + dY
	pZ := pt1[2] + dZ

	// Step 2: inside-outside test
	if !insideTest(tri.Pt1, tri.Pt0, &pX, &pY, &pZ, &nX, &nY, &nZ) {
		return nil
	}

	if !insideTest(tri.Pt2, tri.Pt1, &pX, &pY, &pZ, &nX, &nY, &nZ) {
		return nil
	}

	if !insideTest(tri.Pt0, tri.Pt2, &pX, &pY, &pZ, &nX, &nY, &nZ) {
		return nil
	}

	return &vec3.T{pX, pY, pZ}
}

func insideTest(pt1, pt2 *vec3.T, PX, PY, PZ, NX, NY, NZ *float64) bool {
	edgeX := pt1[0] - pt2[0]
	edgeY := pt1[1] - pt2[1]
	edgeZ := pt1[2] - pt2[2]

	vpX := *PX - pt2[0]
	vpY := *PY - pt2[1]
	vpZ := *PZ - pt2[2]

	cX := edgeY*vpZ - edgeZ*vpY
	cY := edgeZ*vpX - edgeX*vpZ
	cZ := edgeX*vpY - edgeY*vpX

	return *NX*cX+*NY*cY+*NZ*cZ >= 0 //0 -> is on the left side
}
