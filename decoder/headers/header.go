package headers

import (
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
	fileHeader := &BITMAPFILEHEADER{}
	err := binary.Read(reader, binary.LittleEndian, fileHeader)
	if err != nil {
		return nil, fmt.Errorf("error reading file header: %w", err)
	}

	//TODO: determine info header type and decode based on the type.

	var infoHeader InfoHeader
	var infoHeaderSize uint32
	err = binary.Read(reader, binary.LittleEndian, &infoHeaderSize)
	if err != nil {
		return nil, fmt.Errorf("failed to determine size of bmp image: %w", err)
	}

	//TODO this won't work because because there are potentially color tables and palleting, meaning the DataSize
	//which is also known as offset, is not a reliable way of determining the info header type.
	switch fileHeader.DataSize - fileHeader.size() {
	case expectedInfoHeaderSize:
		infoHeader = &BITMAPINFOHEADER{}
		err = binary.Read(reader, binary.LittleEndian, infoHeader)
		if err != nil {
			return nil, fmt.Errorf("error reading info header: %w", err)
		}
	default:
		return nil, fmt.Errorf("unknown Info Header size: %d", fileHeader.DataSize)
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
