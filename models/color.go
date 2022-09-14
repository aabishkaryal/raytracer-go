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
func WriteColor(out io.Writer, color Color, samplesPerPixel int) {
	c := DivideScalar(color, float64(samplesPerPixel))
	r, g, b := math.Sqrt(c.X), math.Sqrt(c.Y), math.Sqrt(c.Z)

	r = 256.0 * utils.Clamp(r, 0, 0.999)
	g = 256.0 * utils.Clamp(g, 0, 0.999)
	b = 256.0 * utils.Clamp(b, 0, 0.999)
	fmt.Fprintf(out, "%v %v %v\n", int(r), int(g), int(b))
}
