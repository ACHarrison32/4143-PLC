package imageManipulator

import (
	"github.com/fogleman/gg"
)

// ImageManipulator represents an image manipulation tool.
type ImageManipulator struct {
	Image     *gg.Context
	ImagePath string // Add a field to store the image path
}

func NewImageManipulator(width, height int) *ImageManipulator {
	img := gg.NewContext(width, height)
	return &ImageManipulator{Image: img}
}

// NewImageManipulatorWithImage loads an existing image and creates an ImageManipulator instance.
func NewImageManipulatorWithImage(imagePath string) (*ImageManipulator, error) {
	img, err := gg.LoadImage(imagePath)
	if err != nil {
		return nil, err
	}

	ctx := gg.NewContextForImage(img)
	return &ImageManipulator{Image: ctx, ImagePath: imagePath}, nil
}

func (im *ImageManipulator) SaveToFile(filename string) error {
	return im.Image.SavePNG(filename)
}

func (im *ImageManipulator) DrawRectangle(x, y, width, height float64) {
	im.Image.DrawRectangle(x, y, width, height)
	im.Image.Stroke()
}
