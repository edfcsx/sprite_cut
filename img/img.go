package img

import (
	"ImageCut/sprite"
	"ImageCut/vector"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
)

var cutIndex int

type Image struct {
	file   *os.File
	isOpen bool
	size   vector.Vec2
}

func (i *Image) Open(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	i.file = file

	imageConfig, _, err := image.DecodeConfig(i.file)

	if err != nil {
		return err
	}

	i.size = vector.Vec2{
		X: imageConfig.Width,
		Y: imageConfig.Height,
	}

	i.isOpen = true
	return nil
}

func (i *Image) Close() error {
	err := i.file.Close()

	if err != nil {
		return err
	}

	return nil
}

func (i *Image) GetSize() (vector.Vec2, error) {
	return i.size, nil
}

func (i *Image) CutArea(p vector.Vec4, s *sprite.Sprite) {
	if !i.isOpen {
		fmt.Println("Image is not open.")
		return
	}

	_, err := i.file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error on reset file cursor.")
		os.Exit(1)
	}

	imageDecoded, _, err := image.Decode(i.file)

	if err != nil {
		fmt.Println("Error on decode image.")
		os.Exit(1)
	}

	rect := image.Rect(p.X, p.Y, p.W, p.H)
	subImg := imageDecoded.(*image.NRGBA).SubImage(rect)

	outFile, err := os.Create(fmt.Sprintf("./%s/%s_%d.jpeg", s.Output, s.Output, cutIndex))

	if err != nil {
		fmt.Println("Error on create image file.")
		os.Exit(1)
	}

	cutIndex++

	err = png.Encode(outFile, subImg)

	if err != nil {
		fmt.Println("Error on encode image.")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Image was cropped and saved successfully.")
}
