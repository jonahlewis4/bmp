package decoder

import (
	"encoding/binary"
	"fmt"
	"github.com/jonahlewis4/bmp/bmp"
	"image"
	"io"
)

type Decoder struct {
	bmp    *bmp.Bmp
	reader io.Reader
}

func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{
		bmp:    &bmp.Bmp{},
		reader: reader,
	}
}
func (d *Decoder) Decode() (image.Image, error) {
	header, err := GetHeaderFromReader(d.reader)
	if err != nil {
		return nil, fmt.Errorf("error decoding bmp header: %w", err)
	}
	d.bmp.Header = header
	pixelData := make([]byte, header.PixelDataSize)
	err = binary.Read(d.reader, binary.LittleEndian, &pixelData)
	if err != nil {
		return nil, fmt.Errorf("error decoding bmp pixel data: %w", err)
	}
	d.bmp.PixelData = &pixelData
	//TODO verify that this works, and add image.Image functions to bmp.
	return d.bmp, nil
}
