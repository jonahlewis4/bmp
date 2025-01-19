package headers

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type Header struct {
	*BITMAPFILEHEADER
	InfoHeader
	fmt.Stringer //literally all this means is that this struct has a String() function
}
type InfoHeader interface {
	size() uint32
}

func GetHeaderFromReader(reader io.Reader) (*Header, error) {
	bufReader := bufio.NewReader(reader)

	fileHeader := &BITMAPFILEHEADER{}
	err := binary.Read(bufReader, binary.LittleEndian, fileHeader)
	if err != nil {
		return nil, fmt.Errorf("error reading file header: %w", err)
	}

	//peek ahead to determine the size of the bitmap.

	var infoHeader InfoHeader
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
	case expectedInfoHeaderSize:
		infoHeader = &BITMAPINFOHEADER{}
		err = binary.Read(bufReader, binary.LittleEndian, infoHeader)
		if err != nil {
			return nil, fmt.Errorf("error reading info header: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported BITMAP type. Specifically, no info header of size %d bytes is supported", infoHeaderSize)
	}

	return &Header{
		BITMAPFILEHEADER: fileHeader,
		InfoHeader:       infoHeader,
	}, nil
}

// GetHeaderFromFileName parses the header of a BMP and errors if it reaches end of file before reading all the data
func GetHeaderFromFileName(fileName string) (*Header, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file '%s': %w", fileName, err)
	}
	defer file.Close()
	return GetHeaderFromReader(file)
}

func (h *Header) String() string {
	return fmt.Sprintf("File Header: %+v,\nInfo Header: %+v\n", h.BITMAPFILEHEADER, h.InfoHeader)
}
