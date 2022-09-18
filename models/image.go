package models

import (
	"fmt"
	"io"
	"math"

	"github.com/aabishkaryal/raytracer-go/utils"
)

type Color = Vec3

// Color Utility functions

// WriteColor writes the color to the given writer
func WriteColor(out io.Writer, color Color, samplesPerPixel float64) {
	c := DivideScalar(color, samplesPerPixel)
	r, g, b := math.Sqrt(c.X), math.Sqrt(c.Y), math.Sqrt(c.Z)

	r = 256.0 * utils.Clamp(r, 0, 0.999)
	g = 256.0 * utils.Clamp(g, 0, 0.999)
	b = 256.0 * utils.Clamp(b, 0, 0.999)
	fmt.Fprintf(out, "%v %v %v\n", int(r), int(g), int(b))
}

// Image structure
type Image struct {
	Width, Height int
	Pixels        [][]Color
}

// NewImage creates a new image
func NewImage(width, height int) Image {
	pixels := make([][]Color, height)
	for i := range pixels {
		pixels[i] = make([]Color, width)
	}

	return Image{
		Width:  width,
		Height: height,
		Pixels: pixels,
	}
}

// SetPixel sets the pixel at the given coordinates
func (i *Image) SetPixel(x, y int, color Color) {
	i.Pixels[y][x] = color
}

// Write writes the image to the given writer
func (img Image) Write(out io.Writer, samplesPerPixel float64) {
	fmt.Fprintf(out, "P3\n%d %d\n255\n", img.Width, img.Height)
	for j := img.Height - 1; j >= 0; j-- {
		for i := 0; i < img.Width; i++ {
			WriteColor(out, img.Pixels[j][i], samplesPerPixel)
		}
	}
}
