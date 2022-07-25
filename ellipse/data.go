package ellipse

const (
	_ERROR_OFFSET = 0.00001
)

// Ellipse-Data for the general ellipse-equation in Canonical form
// (A x2 + Bxy + Cy2 + Dx + Ey + F = 0)
type Data struct {
	A float64 `json:"A"`
	B float64 `json:"B"`
	C float64 `json:"C"`
	D float64 `json:"D"`
	E float64 `json:"E"`
	F float64 `json:"F"`
}
