package main

import (
	"fmt"
	"os"
)

func main() {
	// Image
	ASPECT_RATIO := 16.0 / 9.0
	IMAGE_WIDTH := 800
	IMAGE_HEIGHT := int(float64(IMAGE_WIDTH) / ASPECT_RATIO)

	// World
	world := HittableList{}
	world.Add(Sphere{Center: Point3{0, 0, -1}, Radius: 0.5})
	world.Add(Sphere{Center: Point3{0, -100.5, -1}, Radius: 100})

	// Camera
	viewportHeight := 2.0
	viewportWidth := ASPECT_RATIO * viewportHeight
	focalLength := 1.0

	origin := Point3{0, 0, 0}
	horizontal := Vec3{viewportWidth, 0, 0}
	vertical := Vec3{0, viewportHeight, 0}
	// lower_left_corner = origin - horizontal/2 - vertical/2 - vec3(0, 0, focal_length)
	lowerLeftCorner := SubtractVectors(
		SubtractVectors(
			SubtractVectors(origin,
				DivideScalar(horizontal, 2)),
			DivideScalar(vertical, 2)),
		Vec3{0, 0, focalLength})

	// Render
	fmt.Printf("P3\n%d %d\n255\n", IMAGE_WIDTH, IMAGE_HEIGHT)

	for j := IMAGE_HEIGHT - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d", j)
		for i := 0; i < IMAGE_WIDTH; i++ {
			u := float64(i) / float64(IMAGE_WIDTH-1)
			v := float64(j) / float64(IMAGE_HEIGHT-1)
			// lower_left_corner + u*horizontal + v*vertical - origin
			direction := SubtractVectors(
				AddVectors(
					AddVectors(lowerLeftCorner,
						MultiplyScalar(horizontal, u)),
					MultiplyScalar(vertical, v)),
				origin)

			r := Ray{origin, direction}
			pixelColor := ray_color(r, world)
			WriteColor(os.Stdout, pixelColor)
		}
	}
	fmt.Fprintf(os.Stderr, "\nDone.\n")
}

func ray_color(r Ray, world Hittable) Color {
	if ok, rec := world.Hit(r, 0, INFINITY); ok {
		return MultiplyScalar(AddVectors(rec.Normal, Color{1, 1, 1}), 0.5)
	}

	unitDirection := UnitVector(r.Direction)
	t := 0.5 * (unitDirection.Y + 1.0)
	return AddVectors(MultiplyScalar(Color{1.0, 1.0, 1.0}, 1.0-t), MultiplyScalar(Color{0.5, 0.7, 1.0}, t))
}
