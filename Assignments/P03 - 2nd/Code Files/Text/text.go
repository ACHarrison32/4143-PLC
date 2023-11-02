package main

import (
	"fmt"
	"image"
	"os"

	"github.com/fogleman/gg"
)

func main() {
	// Open the image file
	imageFile, err := os.Open("C:/Users/ACHar/4143-PLC/Assignments/P03 - 2nd/img_get/Img_get/downloaded_image.jpg")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer imageFile.Close()

	// Decode the image
	img, _, err := image.Decode(imageFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Get the image dimensions
	imgWidth := img.Bounds().Max.X
	imgHeight := img.Bounds().Max.Y

	// Create a canvas with the same dimensions as the image
	dc := gg.NewContext(imgWidth, imgHeight)

	// Set text color
	dc.SetRGB(0.5, 0, 0)

	// Load a font with a larger font size (e.g., 48)
	if err := dc.LoadFontFace("C:/Windows/Fonts/arial.ttf", 48); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Draw the original image
	dc.DrawImage(img, 0, 0)

	// Draw the text in the center of the canvas
	dc.DrawStringAnchored("Hello, world!", float64(imgWidth)/2, float64(imgHeight)/2, 0.5, 0.5)

	// Stroke the text
	dc.Stroke()

	// Save the image as PNG
	if err := dc.SavePNG("output_image.png"); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Image saved as output_image.png.")
}
