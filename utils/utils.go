package utils

import (
	"math"
	"math/rand"
)

// Utility Functions

// DegreesToRadians converts degrees to radians
func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

// RandomRange returns a random number between min and max
func RandomRange(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}

// Clamp clamps a value between min and max
func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
