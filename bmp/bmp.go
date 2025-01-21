package bmp

import (
	"github.com/jonahlewis4/bmp/bmp/headers"
	"image"
	"image/color"
)

type Bmp struct {
	Header        *headers.Header
	PixelData     *[]byte
	RowSize       int
	BytesPerPixel int
}

func (bmp *Bmp) ColorModel() color.Model {
	//TODO set this based on info header
	return color.NRGBAModel
}
func (bmp *Bmp) Bounds() image.Rectangle {
	return image.Rect(0, 0, bmp.Header.WidthNoPadding(), bmp.Header.HeightInPixels())
}
func (bmp *Bmp) At(x, y int) color.Color {
	//TODO set this function based on the info header
	idx := x*bmp.RowSize + y
	colorBytes := (*bmp.PixelData)[idx : idx+bmp.BytesPerPixel]

	//TODO make this more generic where we call an info header, which will give a At function or something
	return color.NRGBA{R: colorBytes[0], G: colorBytes[1], B: colorBytes[2], A: 255}
}
