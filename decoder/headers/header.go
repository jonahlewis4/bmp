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
}
type InfoHeader interface {
	size() uint32

	//Stringer is used for testing and debugging. It allows us to customize what this object looks like
	//when %s is used to format it
	fmt.Stringer
}

func GetHeaderFromReader(reader io.Reader) (*Header, error) {
	fileHeader := &BITMAPFILEHEADER{}
	err := binary.Read(reader, binary.LittleEndian, fileHeader)
	if err != nil {
		return nil, fmt.Errorf("error reading file header: %w", err)
	}

	//TODO: determine info header type and decode based on the type.
	var infoHeader InfoHeader
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
