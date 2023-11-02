package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg" // Use this package to decode JPEG images
	"os"
)

func main() {
	// Open the original image
	reader, err := os.Open("C:/Users/ACHar/4143-PLC/Assignments/P03 - 2nd/img_get/Img_get/downloaded_image.jpg")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer reader.Close()

	// Decode the image using the "image/jpeg" package
	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Get image bounds
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Create a new grayscale image
	grayImg := image.NewGray(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Get the color at pixel (x, y)
			oldColor := img.At(x, y)
			r, g, b, _ := oldColor.RGBA()

			// Convert to gray using the formula
			gray := uint8((0.3*float64(r) + 0.59*float64(g) + 0.11*float64(b)) / 256.0)

			// Set the gray color
			grayColor := color.Gray{Y: gray}
			grayImg.Set(x, y, grayColor)
		}
	}

	// Save the grayscale image to an output file
	grayFile, err := os.Create("gray_image.jpg") // Use ".jpg" for JPEG output
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer grayFile.Close()

	err = jpeg.Encode(grayFile, grayImg, nil) // Use jpeg.Encode for JPEG output
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Grayscale image saved as gray_image.jpg.")
}
