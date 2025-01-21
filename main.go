package main

import (
	"fmt"
	"github.com/jonahlewis4/bmp/decoder"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test/bmps/original/rgb-triangle.bmp")
	if err != nil {
		log.Fatal(err)
	}
	decoder := decoder.NewDecoder(file)
	img, err := decoder.Decode()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("first pixel: ", img.At(0, 0))

}
