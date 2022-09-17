package app

import (
	"fmt"
	"os"

	"github.com/aabishkaryal/raytracer-go/models"
)

type Work struct {
	i, j int // Pixel Position
}

type Result struct {
	i, j  int          // Pixel Position
	color models.Color // Pixel color
}

type WorkerManager struct {
	workers                 []Worker
	imageWidth, imageHeight int
	works                   chan Work
	result                  chan Result
}

func NewWorkerManager(imageWidth, imageHeight int, samplesPerPixel, maxDepth float64,
	numWorkers int,
	world models.Hittable,
	cam models.Camera,
) WorkerManager {
	works := make(chan Work, imageWidth*imageHeight)
	result := make(chan Result, numWorkers)
	workers := make([]Worker, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workers[i] = newWorker(works, result, imageWidth, imageHeight,
			samplesPerPixel, maxDepth, world, cam)
	}
	return WorkerManager{
		workers:     workers,
		imageWidth:  imageWidth,
		imageHeight: imageHeight,
		works:       works,
		result:      result,
	}
}

func (wm WorkerManager) Start() [][]models.Color {
	for _, worker := range wm.workers {
		go worker.Start()
	}

	for j := wm.imageHeight - 1; j >= 0; j-- {
		for i := 0; i < wm.imageWidth; i++ {
			wm.works <- Work{i, j}
		}
	}

	close(wm.works)

	image := make([][]models.Color, wm.imageHeight)
	for i := range image {
		image[i] = make([]models.Color, wm.imageWidth)
	}

	totalWork := wm.imageWidth * wm.imageHeight
	for i := 0; i < totalWork; i++ {
		result := <-wm.result
		image[result.j][result.i] = result.color
		fmt.Fprintf(os.Stderr, "\033[2K\rPixels done: %f", (float64(i) / float64(totalWork) * 100.0))
	}
	fmt.Fprintf(os.Stderr, "\033[2K\rPixels done: %f", 100.0)

	return image
}
