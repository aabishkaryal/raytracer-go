package utils

import "math"

// Constants
const (
	INFINITY          = 1e8       // infinity
	ASPECT_RATIO      = 3.0 / 2.0 // aspect ratio
	IMAGE_WIDTH       = 1200      // image width
	SAMPLES_PER_PIXEL = 500       // samples per pixel
	MAX_DEPTH         = 50        // max depth
)

var IMAGE_HEIGHT int = int(math.Floor(float64(IMAGE_WIDTH) / ASPECT_RATIO)) // image height
