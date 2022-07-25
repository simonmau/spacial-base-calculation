package ellipse

import "math"

func (d *Data) CalcTheta() float64 {
	if d.B > -_ERROR_OFFSET && d.B < _ERROR_OFFSET {
		if d.A < d.C {
			return 0.0
		} else {
			return math.Pi * 0.5
		}
	}

	wurzel := (d.A-d.C)*(d.A-d.C) + (d.B * d.B)

	return math.Atan((d.C - d.A - math.Sqrt(wurzel)) / d.B)
}
