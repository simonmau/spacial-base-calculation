package plane

import "github.com/ungerik/go3d/float64/vec3"

func (p *Plane) CaluclateIntersection(pt1, pt2 *vec3.T) vec3.T {
	dX := pt2[0] - pt1[0]
	dY := pt2[1] - pt1[1]
	dZ := pt2[2] - pt1[2]

	subX := p.PointOnPlane[0] - pt1[0]
	subY := p.PointOnPlane[1] - pt1[1]
	subZ := p.PointOnPlane[2] - pt1[2]

	nominator := p.NormalVector[0]*subX + p.NormalVector[1]*subY + p.NormalVector[2]*subZ
	denominator := p.NormalVector[0]*dX + p.NormalVector[1]*dY + p.NormalVector[2]*dZ

	fac := nominator / denominator

	dX *= fac
	dY *= fac
	dZ *= fac

	return vec3.T{pt1[0] + dX, pt1[1] + dY, pt1[2] + dZ}
}
