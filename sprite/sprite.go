package sprite

import (
	"ImageCut/vector"
	"fmt"
	"os"
	"path/filepath"
)

type Sprite struct {
	Path   string
	Output string
	Gap    int
	Size   vector.Vec2
}

func (s *Sprite) Initialize() {
	fmt.Println("Enter the path to the image file:")
	_, err := fmt.Scanln(&s.Path)

	dir, err := filepath.Abs(s.Path)

	if err != nil {
		fmt.Println("Error reading the path to the image file.")
		os.Exit(1)
	}

	s.Path = dir

	if err != nil {
		fmt.Println("Error reading the path to the image file.")
		os.Exit(1)
	}

	fmt.Println("Enter the path to the output folder:")
	_, err = fmt.Scanln(&s.Output)
	//
	//fmt.Println("Gap between sprites (in pixels):")
	//_, err = fmt.Scanln(&s.Gap)
	//
	//fmt.Println("Enter the sprite size (width, in pixels):")
	//_, err = fmt.Scanln(&s.Size.X)
	//
	//fmt.Println("Enter the sprite size (height, in pixels):")
	//_, err = fmt.Scanln(&s.Size.Y)
}
