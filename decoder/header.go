package decoder

type Header struct {
	height int
	width  int
}

func getHeader() *Header {
	return &Header{
		height: 0,
		width:  0,
	}
}
