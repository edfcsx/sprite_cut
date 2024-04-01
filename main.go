package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
)

type Parameters struct {
	path       string
	output     string
	gap        int
	imageSize  Vec2
	spriteSize Vec2
}

type Vec2 struct {
	x int
	y int
}

func (p *Parameters) RequestParams() {
	fmt.Println("Enter the path to the image file:")
	_, err := fmt.Scanln(&p.path)

	dir, err := filepath.Abs(p.path)

	if err != nil {
		fmt.Println("Error reading the path to the image file.")
		os.Exit(1)
	}

	p.path = dir

	if err != nil {
		fmt.Println("Error reading the path to the image file.")
		os.Exit(1)
	}

	fmt.Println("Enter the path to the output folder:")
	_, err = fmt.Scanln(&p.output)

	fmt.Println("Gap between sprites (in pixels):")
	_, err = fmt.Scanln(&p.gap)

	fmt.Println("Enter the sprite size (width, in pixels):")
	_, err = fmt.Scanln(&p.spriteSize.x)

	fmt.Println("Enter the sprite size (height, in pixels):")
	_, err = fmt.Scanln(&p.spriteSize.y)
}

func main() {
	fmt.Println("Sprites Cutout Tool")

	var p = Parameters{}
	p.RequestParams()

	if !checkFileExists(p.path) {
		fmt.Println("File not found.")
		os.Exit(1)
	}

	dimensions, err := GetImageSize(p.path)

	if err != nil {
		fmt.Println("Error reading the image metadata.")
		fmt.Println(err)
		os.Exit(1)
	}

	p.imageSize = dimensions

	outputPath, err := filepath.Abs(fmt.Sprintf("./%s", p.output))

	if err != nil {
		fmt.Println("Error reading the output path.")
		os.Exit(1)
	}

	err = os.MkdirAll(outputPath, os.ModePerm)

	if err != nil {
		fmt.Println("Error creating the output folder.")
		os.Exit(1)
	}

	fmt.Println("Path to the image file:", p.path)
}

func checkFileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func GetImageSize(path string) (Vec2, error) {
	file, err := os.Open(path)

	if err != nil {
		return Vec2{}, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing the file.")
		}
	}(file)

	imageConfig, _, err := image.DecodeConfig(file)
	if err != nil {
		return Vec2{}, err
	}

	return Vec2{imageConfig.Width, imageConfig.Height}, nil
}

func CutImageArea() {

}
