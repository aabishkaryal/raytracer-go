package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	// World
	world := HittableList{}

	materialGround := Lambertian{Color{0.8, 0.8, 0.0}}
	materialCenter := Lambertian{Color{0.1, 0.2, 0.5}}
	materialLeft := Dielectric{1.5}
	materialRight := Metal{Color{0.8, 0.6, 0.2}, 1.0}

	world.Add(Sphere{Vec3{0, -100.5, -1}, 100, materialGround})
	world.Add(Sphere{Vec3{0, 0, -1}, 0.5, materialCenter})
	world.Add(Sphere{Vec3{-1, 0, -1}, 0.5, materialLeft})
	world.Add(Sphere{Vec3{-1, 0, -1}, -0.4, materialLeft})
	world.Add(Sphere{Vec3{1, 0, -1}, 0.5, materialRight})

	// Camera
	cam := NewCamera()

	// Render
	fmt.Printf("P3\n%d %d\n255\n", IMAGE_WIDTH, IMAGE_HEIGHT)

	for j := IMAGE_HEIGHT - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\033[2K\rScanlines remaining: %d/%d", j, IMAGE_HEIGHT)
		for i := 0; i < IMAGE_WIDTH; i++ {
			pixelColor := Color{0, 0, 0}
			for s := 0; s < SAMPLES_PER_PIXEL; s++ {
				u := (float64(i) + RandomRange(0, 1)) / float64(IMAGE_WIDTH-1)
				v := (float64(j) + RandomRange(0, 1)) / float64(IMAGE_HEIGHT-1)
				r := cam.GetRay(u, v)
				pixelColor = AddVectors(pixelColor, RayColor(r, world, MAX_DEPTH))
			}
			WriteColor(os.Stdout, pixelColor, SAMPLES_PER_PIXEL)
		}
	}
	fmt.Fprintf(os.Stderr, "\nDone.\n")
}

func RayColor(r Ray, world Hittable, depth int) Color {
	if depth <= 0 {
		return Color{0, 0, 0}
	}

	if ok, rec := world.Hit(r, 0.001, INFINITY); ok {
		if ok, scattered, attenuation := rec.MatPtr.Scatter(r, rec); ok {
			return MultiplyVectors(attenuation, RayColor(scattered, world, depth-1))
		}
		return Color{0, 0, 0}
	}

	unitDirection := UnitVector(r.Direction)
	t := 0.5 * (unitDirection.Y + 1.0)
	return AddVectors(MultiplyScalar(Color{1.0, 1.0, 1.0}, 1.0-t), MultiplyScalar(Color{0.5, 0.7, 1.0}, t))
}
