package headers

import (
	"fmt"
)

type Header struct {
	*BITMAPFILEHEADER
	InfoHeader
	PixelDataSize uint32
	fmt.Stringer  //literally all this means is that this struct has a String() function

}
type InfoHeader interface {
	size() uint32
	PixelDataSize() uint32
}

func (h *Header) String() string {
	return fmt.Sprintf("File Header: %+v,\nInfo Header: %+v\n", h.BITMAPFILEHEADER, h.InfoHeader)
}
