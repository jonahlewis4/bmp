package decoder

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/jonahlewis4/bmp/bmp/headers"
	"io"
	"os"
)

func GetHeaderFromReader(reader io.Reader) (*headers.Header, error) {

	bufReader := bufio.NewReader(reader)

	fileHeader := &headers.BITMAPFILEHEADER{}
	err := binary.Read(bufReader, binary.LittleEndian, fileHeader)
	if err != nil {
		return nil, fmt.Errorf("error reading file header: %w", err)
	}

	//peek ahead to determine the size of the bitmap.

	var infoHeader headers.InfoHeader
	var infoHeaderSize uint32

	sizeBytes, err := bufReader.Peek(binary.Size(infoHeaderSize))
	if err != nil {
		return nil, fmt.Errorf("error determining size of header: %w", err)
	}
	err = binary.Read(bytes.NewReader(sizeBytes), binary.LittleEndian, &infoHeaderSize)
	if err != nil {
		return nil, fmt.Errorf("failed to determine size of bmp image: %w", err)
	}

	//determine the info header type based on the header size.
	switch infoHeaderSize {
	case headers.ExpectedInfoHeaderSize:
		infoHeader = &headers.BITMAPINFOHEADER{}
		err = binary.Read(bufReader, binary.LittleEndian, infoHeader)
		if err != nil {
			return nil, fmt.Errorf("error reading info header: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported BITMAP type. Specifically, no info header of size %d bytes is supported", infoHeaderSize)
	}

	return &headers.Header{
		InfoHeader:       infoHeader,
		BITMAPFILEHEADER: fileHeader,
		PixelDataSize:    infoHeader.PixelDataSize(),
	}, nil
}

// GetHeaderFromFileName parses the header of a BMP and errors if it reaches end of file before reading all the data
func GetHeaderFromFileName(fileName string) (*headers.Header, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file '%s': %w", fileName, err)
	}
	defer file.Close()
	return GetHeaderFromReader(file)
}
