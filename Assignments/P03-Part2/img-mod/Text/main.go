// Text/main.go

// Package Text provides functionality to add text to an image.
package Text

import (
	"io/ioutil"
	"os"

	"github.com/fogleman/gg"
	"golang.org/x/image/font/gofont/goregular"
)

// AddTextToImage adds text to an image and saves it.
// W and H are the width and height of the image.
func AddTextToImage(W int, H int, outputFilename string, text string) {
	// Create a temporary file and write the byte slice to it
	tempFile, err := ioutil.TempFile("", "font-*.ttf")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(goregular.TTF); err != nil {
		panic(err)
	}

	dc := gg.NewContext(W, H)

	if err := dc.LoadFontFace(tempFile.Name(), 72); err != nil {
		panic(err)
	}

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	dc.SetRGB(.5, 0, 0)
	dc.DrawStringAnchored(text, float64(W)/2, float64(H)/2, 0.5, 0.5)
	dc.Stroke()

	dc.SavePNG(outputFilename)
}
