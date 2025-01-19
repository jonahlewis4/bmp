package main

import (
	"fmt"
	"github.com/jonahlewis4/bmp/decoder/headers"
	"time"
)

func main() {
	time.Sleep(3 * time.Second)
	header := headers.BITMAPINFOHEADER{}
	fmt.Printf("%+v\n", header)
}
