package headers

import "math"

const (
	BI_RGB = iota
	BI_RLE8
	BI_RLE4
	BI_BITFIELDS
	BI_JPEG
	BI_PNG
	BI_ALPHABITFIELDS
	BI_CMYK
	BI_CMYKRLE8
	BI_CMYKRLE4
)

type BITMAPINFOHEADER struct {
	Size               uint32
	Width              int32
	Height             int32
	Planes             int16
	BitsPerPixel       uint16
	Compression        int32
	ImageSize          uint32
	HorizontalRes      int32
	VerticalRes        int32
	NumColors          uint32
	NumImportantColors uint32
}

const ExpectedInfoHeaderSize = 40

func (h *BITMAPINFOHEADER) size() uint32 {
	return h.size()
}
func (h *BITMAPINFOHEADER) PixelDataSize() uint64 {
	if h.ImageSize == 0 {
		//if h is zero the pixel data size must be calculated. This can be determined based
		//on the BitsPerPixel and width. Bmp pixels are stored as rows
		//A row must be a multiple of 4 bytes.

		//if the image has a Width of 5, that means there will be 5 pixels in each row.
		//however, suppose the BitsPerPixel has a size of 24, 5 pixels per row.
		//this would mean 24 * 5 = 120 bits per row, or 15 bytes. We need the padding to be a multiple
		//of 4, so the padding will be what we have to add to get to the next multiple of 4.
		// an easy way to calculate this is:

		// ceil((BitsPerPixel * Width)/32) * 4 - BitsPerPixel * Width / 8
		//padding is not used in this function but there are plans to come back to here later.
		return uint64(math.Ceil(float64(int64(h.BitsPerPixel)*int64(h.Width))/32)*4) * uint64(math.Abs(float64(h.Height)))
	}
	return uint64(h.ImageSize)
}
func (h *BITMAPINFOHEADER) RowSize() uint64 {
	return uint64(math.Ceil(float64(int64(h.BitsPerPixel)*int64(h.Width))/32) * 4)
}
