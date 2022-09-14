package models

import "math/rand"

// RandomScene generates a random scene like the cover from the book
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
