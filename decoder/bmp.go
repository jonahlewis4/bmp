package decoder

import (
	"encoding/binary"
	"fmt"
	"github.com/jonahlewis4/bmp/decoder/headers"
	"image"
	"io"
)

type bmp struct {
	header    *headers.Header
	pixelData *[]byte
}
type Decoder struct {
	bmp    *bmp
	reader io.Reader
}

func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{
		bmp:    &bmp{},
		reader: reader,
	}
}
func (d *Decoder) Decode() (image.Image, error) {
	header, err := headers.GetHeaderFromReader(d.reader)
	if err != nil {
		return nil, fmt.Errorf("error decoding bmp header: %w", err)
	}
	d.bmp.header = header
	pixelData := make([]byte, header.PixelDataSize)
	err = binary.Read(d.reader, binary.LittleEndian, &pixelData)
	if err != nil {
		return nil, fmt.Errorf("error decoding bmp pixel data: %w", err)
	}
	d.bmp.pixelData = &pixelData
	//TODO verify that this works, and add image.Image functions to bmp.
	return d.bmp, nil
}
