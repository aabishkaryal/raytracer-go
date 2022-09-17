package app

import (
	"github.com/aabishkaryal/raytracer-go/models"
	"github.com/aabishkaryal/raytracer-go/utils"
)

type Worker struct {
	works                     <-chan Work
	result                    chan<- Result
	imageWidth, imageHeight   int
	samplesPerPixel, maxDepth float64
	world                     models.Hittable
	cam                       models.Camera
}

func newWorker(works <-chan Work, result chan<- Result, imageWidth, imageHeight int, samplesPerPixel, maxDepth float64, world models.Hittable, cam models.Camera) Worker {
	return Worker{
		works:           works,
		result:          result,
		imageWidth:      imageWidth,
		imageHeight:     imageHeight,
		samplesPerPixel: samplesPerPixel,
		maxDepth:        maxDepth,
		world:           world,
		cam:             cam,
	}
}

func (w Worker) Start() {
	for work := range w.works {
		i, j := float64(work.i), float64(work.j)
		pixelColor := models.Color{X: 0, Y: 0, Z: 0}
		for s := 0.0; s < w.samplesPerPixel; s++ {
			u := (i + utils.RandomRange(0, 1)) / float64(w.imageWidth-1)
			v := (j + utils.RandomRange(0, 1)) / float64(w.imageHeight-1)
			r := w.cam.GetRay(u, v)
			pixelColor = models.AddVectors(pixelColor, RayColor(r, w.world, w.maxDepth))
		}
		w.result <- Result{i: work.i, j: work.j, color: pixelColor}
	}
}
