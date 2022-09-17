package app

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/aabishkaryal/raytracer-go/models"
	"github.com/aabishkaryal/raytracer-go/utils"
)

func Raytrace(imageWidth int,
	samplesPerPixel, maxDepth, aspectRatio, verticalFieldOfView, numCPUs float64,
	output io.Writer,
) {
	rand.Seed(time.Now().Unix())

	imageHeight := int(math.Floor(float64(imageWidth) / aspectRatio)) // image height

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
		verticalFieldOfView,
		aspectRatio,
		aperture,
		distToFocus,
	)

	// Render with workers
	workerManager := NewWorkerManager(imageWidth, imageHeight, samplesPerPixel, maxDepth, int(numCPUs), world, cam)
	image := workerManager.Start()

	fmt.Fprintf(output, "P3\n%d %d\n255\n", imageWidth, imageHeight)
	for j := imageHeight - 1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			models.WriteColor(output, image[j][i], samplesPerPixel)
		}
	}
	fmt.Fprintf(os.Stderr, "\nDone.\n")
}

func RayColor(r models.Ray, world models.Hittable, depth float64) models.Color {
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
