package main

import (
	"fmt"
	"os"
)

const (
	IMAGE_WIDTH  = 256
	IMAGE_HEIGHT = 256
)

func main() {
	fmt.Printf("P3\n%d %d\n255\n", IMAGE_WIDTH, IMAGE_HEIGHT)

	for j := IMAGE_HEIGHT - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d", j)
		for i := 0; i < IMAGE_WIDTH; i++ {
			pixelColor := Color{float64(i) / (IMAGE_WIDTH - 1), float64(j) / (IMAGE_HEIGHT - 1), 0.25}
			WriteColor(os.Stdout, pixelColor)
		}
	}
	fmt.Fprintf(os.Stderr, "\nDone.\n")
}
