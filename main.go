package main

import (
	"ImageCut/file"
	"ImageCut/img"
	"ImageCut/sprite"
	"ImageCut/vector"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Sprites Cutout Tool")

	spr := sprite.Sprite{}
	//spr.Initialize()
	spr.Path = "./sprites.png"
	spr.Output = "./portraits"

	path, err := filepath.Abs(spr.Path)

	if err != nil {
		fmt.Println("Error on get sprite path.")
		os.Exit(1)
	}

	if !file.Exists(path) {
		fmt.Println("Sprite not found.")
		os.Exit(1)
	}

	sprImg := img.Image{}
	err = sprImg.Open(spr.Path)
	defer sprImg.Close()

	if err != nil {
		fmt.Println("Error on open sprite image.")
		os.Exit(1)
	}

	outputPath, err := filepath.Abs(fmt.Sprintf("./%s", spr.Output))

	if err != nil {
		fmt.Println("Error reading the output path.")
		os.Exit(1)
	}

	err = os.MkdirAll(outputPath, os.ModePerm)

	if err != nil {
		fmt.Println("Error creating the output folder.")
		os.Exit(1)
	}

	for i := 0; i < 4; i++ {
		imgSizeX := 74
		gap := 6
		x := imgSizeX * i

		imgSizeY := 74

		if x > 0 {
			x += gap
		}

		sprImg.CutArea(vector.Vec4{
			x,
			imgSizeY,
			x + imgSizeX,
			imgSizeY + imgSizeY + gap,
		}, &spr)
	}
}
