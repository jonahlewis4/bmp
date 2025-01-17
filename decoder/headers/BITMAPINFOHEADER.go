package headers

import "fmt"

type BITMAPINFOHEADER struct {
	Size   uint32
	Width  uint32
	Height uint32
	Planes int16
}

const expectedInfoHeaderSize = 40

func (h *BITMAPINFOHEADER) size() uint32 {
	return h.size()
}

func (h *BITMAPINFOHEADER) String() string {
	return fmt.Sprintf("{Size: %d, Width: %d, Height: %d, Planes: %d}", h.Size, h.Width, h.Height, h.Planes)
}
