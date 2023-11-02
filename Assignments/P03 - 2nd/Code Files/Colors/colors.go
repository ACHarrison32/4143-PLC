package main

import (
	"fmt"
	"image"
	_ "image/jpeg" // Import this package to decode JPEG images
	"io"
	"os"
)

func main() {
	// Open the image file
	reader, err := os.Open("C:/Users/ACHar/4143-PLC/Assignments/P03 - 2nd/img_get/Img_get/downloaded_image.jpg")
	if err != nil {
		fmt.Printf("Error opening image file: %v\n", err)
		return
	}
	defer reader.Close() // Ensure the file is closed when we're done with it

	// Decode the image
	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error decoding image: %v\n", err)
		return
	}

	// Create an output file for writing pixel information
	outputFile, err := os.Create("pixel_info.txt")
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outputFile.Close() // Ensure the output file is closed when we're done

	// Get image dimensions
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Create a writer for the output file
	writer := io.Writer(outputFile)

	// Loop through each pixel and write its RGB values to the output file
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()
			pixelInfo := fmt.Sprintf("Pixel at (%d, %d) - R: %d, G: %d, B: %d\n", x, y, r>>8, g>>8, b>>8)

			// Write the pixel information to the output file
			_, err := writer.Write([]byte(pixelInfo))
			if err != nil {
				fmt.Printf("Error writing to output file: %v\n", err)
				return
			}
		}
	}

	fmt.Println("Pixel information has been written to pixel_info.txt")
}
