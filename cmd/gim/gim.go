package main

import (
	"bytes"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	gim "github.com/Neo-Desktop/go-image-merge"
)

func main() {

	imgFile, err := os.Open("./192.jfif")
	if err != nil {
		log.Panicf(err.Error())
	}

    buf := new(bytes.Buffer)
	buf.ReadFrom(imgFile)
	image := buf.Bytes()

	grids := []*gim.Grid{
		{
			ImageBytes:   image,
			ImageType:    "jfif",
			BackgroundColor: color.White,
		},
		{
			ImageFilePath:   "./192.jfif",
			BackgroundColor: color.RGBA{R: 0x8b, G: 0xd0, B: 0xc6},
		},
	}
	rgba, err := gim.New(grids, 1, 2).Merge()
	if err != nil {
		log.Panicf(err.Error())
	}

	file, err := os.Create("merged.jpg")
	if err != nil {
		log.Panicf(err.Error())
	}

	err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 80})
	if err != nil {
		log.Panicf(err.Error())
	}
}
