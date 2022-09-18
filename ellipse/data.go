package ellipse

const (
	_ERROR_OFFSET = 0.00001
)

// Ellipse-Data for the general ellipse-equation in Canonical form
// (A x2 + Bxy + Cy2 + Dx + Ey + F = 0)
type Data struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
	C float64 `json:"c"`
	D float64 `json:"d"`
	E float64 `json:"e"`
	F float64 `json:"f"`
}
