package test

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func byteToRgbImage(dataPtr *[]uint8, width, height int) *image.RGBA {
	data := *dataPtr

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	size := img.Bounds().Size()

	for y := 0; y < size.Y; y++ {

		for x := 0; x < size.X; x++ {
			index := 3 * (y*width + x)

			col := color.RGBA{
				B: data[index],
				G: data[index+1],
				R: data[index+2],
				A: 255}

			img.Set(x, y, col)
		}
	}

	return img
}

func saveImage(path string, img *image.Image) error {
	f, err := os.Create(path)

	if err != nil {
		return err
	}
	defer f.Close()

	// Encode to `PNG` with `DefaultCompression` level
	// then save to file
	err = png.Encode(f, *img)
	if err != nil {
		return err
	}

	return nil
}

func saveImageRgba(path string, img *image.RGBA) error {
	imgVal := *img
	baseImg := imgVal.SubImage(img.Rect)

	return saveImage(path, &baseImg)
}
