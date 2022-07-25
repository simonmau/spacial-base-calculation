package ellipse

import "github.com/ungerik/go3d/float64/vec2"

func (d *Data) CalcCenter() vec2.T {
	x0Zähler := 2.0*d.C*d.D - d.B*d.E
	y0Zähler := 2.0*d.A*d.E - d.B*d.D

	nenner := d.B*d.B - 4.0*d.A*d.C

	return vec2.T{x0Zähler / nenner, y0Zähler / nenner}
}
