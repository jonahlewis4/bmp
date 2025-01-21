package decoder

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/jonahlewis4/bmp/bmp"
	"image"
	"io"
)

type Decoder struct {
	bmp       *bmp.Bmp
	bufReader *bufio.Reader
}

func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{
		bmp:       &bmp.Bmp{},
		bufReader: bufio.NewReader(reader),
	}
}
func (d *Decoder) Decode() (image.Image, error) {
	header, err := GetHeaderFromBufReader(d.bufReader)
	if err != nil {
		return nil, fmt.Errorf("error decoding bmp header: %w", err)
	}
	d.bmp.Header = header
	pixelData := make([]byte, header.PixelDataSize)
	err = binary.Read(d.bufReader, binary.LittleEndian, &pixelData)
	if err != nil {
		return nil, fmt.Errorf("error decoding bmp pixel data: %w", err)
	}
	d.bmp.PixelData = &pixelData
	d.bmp.RowSize = d.bmp.Header.RowSize()
	d.bmp.BytesPerPixel = d.bmp.Header.BitsPerSinglePixel() / 8
	//TODO verify that this works, and add image.Image functions to bmp.
	return d.bmp, nil
}
