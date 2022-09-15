package main

import (
	"flag"

	"github.com/aabishkaryal/raytracer-go/app"
	"github.com/aabishkaryal/raytracer-go/utils"
)

func main() {
	// Parse command line arguments
	imageWidth := flag.Int("width", utils.IMAGE_WIDTH, "Width of the image.")
	aspectRatio := flag.Float64("aspectRatio", utils.ASPECT_RATIO, "Aspect Ratio of the image.")
	samplesPerPixel := flag.Int("samplesPerPixel", utils.SAMPLES_PER_PIXEL, "Number of samples per pixel.")
	maxDepth := flag.Int("maxDepth", utils.MAX_DEPTH, "Maximum depth of the ray to trace.")
	verticalFieldOfView := flag.Int("verticalFOV", utils.VERTICAL_FOV, "Vertical field of view of the camera.")

	flag.Parse()

	app.Raytrace(
		float64(*imageWidth),
		float64(*samplesPerPixel),
		float64(*maxDepth),
		*aspectRatio,
		float64(*verticalFieldOfView),
	)
}
