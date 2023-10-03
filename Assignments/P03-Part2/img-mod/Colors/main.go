// Colors/main.go

// Package Colors provides functionality to extract colors from an image.
package Colors

import (
	"fmt"
	"image"
	_ "image/png" // Import this package to decode PNGs
	"os"
)

// ExtractColors extracts the RGB values of each pixel in an image.
// It takes a string argument specifying the path to the image.
func ExtractColors(imagePath string) {
	reader, err := os.Open(imagePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()
			fmt.Printf("Pixel at (%d, %d) - R: %d, G: %d, B: %d\n", x, y, r>>8, g>>8, b>>8)
		}
	}
}
