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
	world := RandomScene()

	// Camera
	lookFrom := Point3{13, 2, 3}
	lookAt := Point3{0, 0, 0}
	vUp := Vec3{0, 1, 0}
	distToFocus := 10.0
	aperture := 0.1
	cam := NewCamera(
		lookFrom,
		lookAt,
		vUp,
		20,
		ASPECT_RATIO,
		aperture,
		distToFocus,
	)

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

func RandomScene() HittableList {
	world := HittableList{}

	groundMaterial := Lambertian{Color{0.5, 0.5, 0.5}}

	world.Add(Sphere{Point3{0, -1000, 0}, 1000, groundMaterial})

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := Point3{
				float64(a) + 0.9*rand.Float64(),
				0.2,
				float64(b) + 0.9*rand.Float64(),
			}

			if SubtractVectors(center, Point3{4, 0.2, 0}).Length() > 0.9 {
				if chooseMat < 0.8 {
					// diffuse
					albedo := MultiplyVectors(RandomVector(), RandomVector())
					sphereMaterial := Lambertian{albedo}
					world.Add(Sphere{center, 0.2, sphereMaterial})
				} else if chooseMat < 0.95 {
					// metal
					albedo := RandomVectorRange(0.5, 1)
					sphereMaterial := Lambertian{albedo}
					world.Add(Sphere{center, 0.2, sphereMaterial})
				} else {
					// glass
					sphereMaterial := Dielectric{1.5}
					world.Add(Sphere{center, 0.2, sphereMaterial})
				}
			}
		}
	}

	material1 := Dielectric{1.5}
	world.Add(Sphere{Point3{0, 1, 0}, 1.0, material1})

	material2 := Lambertian{Color{0.4, 0.2, 0.1}}
	world.Add(Sphere{Point3{-4, 1, 0}, 1.0, material2})

	material3 := Metal{Color{0.7, 0.6, 0.5}, 0.0}
	world.Add(Sphere{Point3{4, 1, 0}, 1.0, material3})

	return world
}
