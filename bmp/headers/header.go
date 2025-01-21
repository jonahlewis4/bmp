package headers

import (
	"fmt"
)

type Header struct {
	*BITMAPFILEHEADER
	InfoHeader
	PixelDataSize uint64
	fmt.Stringer  //literally all this means is that this struct has a String() function

}
type InfoHeader interface {
	size() uint32
	PixelDataSize() uint64
	RowSize() int
	WidthNoPadding() int
	HeightInPixels() int
	BitsPerSinglePixel() int
}

func (h *Header) String() string {
	return fmt.Sprintf("File Header: %+v,\nInfo Header: %+v\n", h.BITMAPFILEHEADER, h.InfoHeader)
}
