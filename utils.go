package main

import "math/rand"

// Constants
const (
	INFINITY          = 1e8
	PI                = 3.1415926535897932385
	ASPECT_RATIO      = 16.0 / 9.0
	IMAGE_WIDTH       = 800
	SAMPLES_PER_PIXEL = 100
	MAX_DEPTH         = 50
)

var IMAGE_HEIGHT int = int(float64(IMAGE_WIDTH) / ASPECT_RATIO)

func DegreesToRadians(degrees float64) float64 {
	return degrees * PI / 180.0
}

func RandomRange(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}

func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
