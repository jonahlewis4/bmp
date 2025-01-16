package decoder

import (
	"encoding/binary"
	"fmt"
	"os"
)

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

type Header struct {
	BITMAPFILEHEADER
	InfoHeader
}

type InfoHeader interface {
	size() int
}
type BITMAPFILEHEADER struct {
	Signature int16  //'BM'
	FileSize  uint32 //File size in bytes
	Reserved  uint32 //unused(=0)
	DataSize  uint32 //File offset to Raster Data
}

// size returns the size of the file header, which is always 14
func (h *BITMAPFILEHEADER) size() uint32 {
	return 14
}

type BITMAPINFOHEADER struct {
	Size   uint32
	Width  uint32
	Height uint32
	Planes int16
}

func (h *BITMAPINFOHEADER) size() uint32 {
	return h.Size
}

// GetHeaderFromFileName parses the header of a BMP and errors if it reaches end of file before reading all the data
func GetHeaderFromFileName(fileName string) (*Header, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file '%s': %w", fileName, err)
	}
	defer file.Close()
	//TODO only read the file header. Need to add custom handling for the different possible bitmap headers
	header := &Header{}
	err = binary.Read(file, binary.LittleEndian, header)
	if err != nil {
		return nil, fmt.Errorf("error reading file '%s': %w", fileName, err)
	}
	return header, nil
}
