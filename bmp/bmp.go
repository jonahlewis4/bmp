package bmp

import (
	"github.com/jonahlewis4/bmp/bmp/headers"
	"image"
	"image/color"
)

type Bmp struct {
	Header    *headers.Header
	PixelData *[]byte
}

func (bmp *Bmp) ColorModel() color.Model {
	//TODO set this function based on the info header
	return color.RGBAModel
}
func (bmp *Bmp) Bounds() image.Rectangle {
	//TODO set this function based on the info header
	return image.Rect(0, 0, 0, 0)
}
func (bmp *Bmp) At(x, y int) color.Color {
	//TODO set this function based on the info header
	return color.RGBA{R: uint8(x), G: uint8(y)}
}
