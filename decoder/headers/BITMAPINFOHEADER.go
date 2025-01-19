package headers

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

const expectedInfoHeaderSize = 40

func (h *BITMAPINFOHEADER) size() uint32 {
	return h.size()
}
func (h *BITMAPINFOHEADER) pixelDataSize() uint32 {
	return h.ImageSize
}
