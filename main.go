package main

import (
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
	_, err = decoder.Decode()
	if err != nil {
		log.Fatal(err)
	}

}
