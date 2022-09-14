package app

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/aabishkaryal/raytracer-go/models"
	"github.com/aabishkaryal/raytracer-go/utils"
)

func Raytrace() {
	rand.Seed(time.Now().Unix())

	// World
	world := models.RandomScene()

	// Camera
	lookFrom := models.Point3{X: 13, Y: 2, Z: 3}
	lookAt := models.Point3{X: 0, Y: 0, Z: 0}
	vUp := models.Vec3{X: 0, Y: 1, Z: 0}
	distToFocus := 10.0
	aperture := 0.1
	cam := models.NewCamera(
		lookFrom,
		lookAt,
		vUp,
		20,
		utils.ASPECT_RATIO,
		aperture,
		distToFocus,
	)

	// Render
	fmt.Printf("P3\n%d %d\n255\n", utils.IMAGE_WIDTH, utils.IMAGE_HEIGHT)

	for j := utils.IMAGE_HEIGHT - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\033[2K\rScanlines remaining: %d/%d", j, utils.IMAGE_HEIGHT)
		for i := 0; i < utils.IMAGE_WIDTH; i++ {
			pixelColor := models.Color{X: 0, Y: 0, Z: 0}
			for s := 0; s < utils.SAMPLES_PER_PIXEL; s++ {
				u := (float64(i) + utils.RandomRange(0, 1)) / float64(utils.IMAGE_WIDTH-1)
				v := (float64(j) + utils.RandomRange(0, 1)) / float64(utils.IMAGE_HEIGHT-1)
				r := cam.GetRay(u, v)
				pixelColor = models.AddVectors(pixelColor, RayColor(r, world, utils.MAX_DEPTH))
			}
			models.WriteColor(os.Stdout, pixelColor, utils.SAMPLES_PER_PIXEL)
		}
	}
	fmt.Fprintf(os.Stderr, "\nDone.\n")
}

func RayColor(r models.Ray, world models.Hittable, depth int) models.Color {
	if depth <= 0 {
		return models.Color{X: 0, Y: 0, Z: 0}
	}

	if ok, rec := world.Hit(r, 0.001, utils.INFINITY); ok {
		if ok, scattered, attenuation := rec.MatPtr.Scatter(r, rec); ok {
			return models.MultiplyVectors(attenuation, RayColor(scattered, world, depth-1))
		}
		return models.Color{X: 0, Y: 0, Z: 0}
	}

	unitDirection := models.UnitVector(r.Direction)
	t := 0.5 * (unitDirection.Y + 1.0)
	return models.AddVectors(
		models.MultiplyScalar(models.Color{X: 1.0, Y: 1.0, Z: 1.0}, 1.0-t),
		models.MultiplyScalar(models.Color{X: 0.5, Y: 0.7, Z: 1.0}, t))
}
