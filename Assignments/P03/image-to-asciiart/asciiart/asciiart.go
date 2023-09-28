package asciiart

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/nfnt/resize"
)

const charset = "@%#*+=-:. "

func imageToGray(img image.Image) *image.Gray {
	grayImage := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			grayImage.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}
	return grayImage
}

func ConvertToASCII(imagePath string, cols, rows int) (string, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	img = resize.Resize(uint(cols), 0, img, resize.Lanczos3)

	grayImg := imageToGray(img)

	var asciiArt string
	step := grayImg.Bounds().Max.Y / rows

	for y := 0; y < grayImg.Bounds().Max.Y; y += step {
		for x := 0; x < grayImg.Bounds().Max.X; x++ {
			g := grayImg.GrayAt(x, y).Y
			index := int(float64(g) / 255.0 * float64(len(charset)-1))
			asciiArt += string(charset[index])
		}
		asciiArt += "\n"
	}

	return asciiArt, nil
}
