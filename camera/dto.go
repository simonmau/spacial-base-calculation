package camera

import (
	"github.com/simonmau/spacial-base-calculation/dto"
)

type CameraDto struct {
	//where the camera is positioned
	Eye dto.JsonVector3 `json:"CameraEye" validate:"required"`

	//the sight-direction will be calculated starting from the eye
	LooksThrough dto.JsonVector3 `json:"LooksThrough" validate:"required"`

	//rotation around the sight-axis (lt - eye)
	Rotation float64 `json:"Rotation" validate:"required"`

	//a sensor size can be set here, changes will result in zoom effect, default should be 36mm for 'vollformat'
	SensorWidth float64 `json:"SensorWidth" validate:"required"`

	//the zoom-length of the camera, will move the projection-layer from the camera back and forwards
	//with sensor-width 36mm should be at least 30mm (extremely-wide) to max 200mm (extreme-zoom)
	FocalLength float64 `json:"FocalLength" validate:"required"`

	//the projection-width is the width of the projection-area where the image in px will be projected on
	//this should be in normal cases 100mm
	//with a fixed projection-width, every width/height px produces the same image, just less sharp (or more)
	ProjectionWidth float64 `json:"ProjectionWidth" validate:"required"`

	//lenscorrection warps the image before/after the necessary transformations are finished, so its on top of everything else
	//can be nil -> no lenscorrection will be applied
	LensCorrection *float64 `json:"LensCorrection"`

	//can be nil when lenscorrection is set -> will then be width/2 and height/2
	LensCorrectionCenter *dto.JsonVector2 `json:"LensCorrectionCenter"`

	//the rendered image sizes
	Width int32 `json:"Width" validate:"required"`

	//the rendered image sizes
	Height int32 `json:"Height" validate:"required"`
}
