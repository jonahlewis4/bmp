package headers

import "fmt"

// BitmapSignature is the bitmap magic number or file signature.
// The first two bytes of a BMP file will be the characters B and M.
// B is ascii 0x42. M is ascii 0x4D
// 0x424D will be the first two bytes of a bmp file
// However, we will read this in Little Endian order,
// Meaning the left bytes will be the least significant,
// hence the swapping to 0x4D42
const (
	BitmapSignature = 0x4D42
)

// BitmapReserved will always be 0 for BMPs.
const (
	BitmapReserved = 0
)

type BITMAPFILEHEADER struct {
	Signature int16  //'BM'
	FileSize  uint32 //File size in bytes
	Reserved  uint32 //unused(=0)
	DataSize  uint32 //File offset to Raster Data (offset until there are pixels)
}

const fileHeaderSize = 14

// size returns the size of the file header, which is always 14
func (h *BITMAPFILEHEADER) size() uint32 {
	return fileHeaderSize
}
func (h *BITMAPFILEHEADER) String() string {
	return fmt.Sprintf("{Signature:%d FileSize:%d Reserved:%d DataSize:%d }", h.Signature, h.FileSize, h.Reserved, h.DataSize)
}
