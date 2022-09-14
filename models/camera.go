package models

import (
	"math"

	"github.com/aabishkaryal/raytracer-go/utils"
)

type Camera struct {
	Origin          Point3
	LowerLeftCorner Point3
	Horizontal      Vec3
	Vertical        Vec3
	U, V, W         Vec3
	LensRadius      float64
}

func NewCamera(lookFrom, lookAt Point3,
	vup Vec3,
	verticalFOV, aspectRatio float64,
	aperture, focusDistance float64,
) Camera {
	theta := utils.DegreesToRadians(verticalFOV)
	h := math.Tan(theta / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := UnitVector(SubtractVectors(lookFrom, lookAt))
	u := UnitVector(VectorCross(vup, w))
	v := VectorCross(w, u)

	origin := lookFrom
	horizontal := MultiplyScalar(u, viewportWidth*focusDistance)
	vertical := MultiplyScalar(v, viewportHeight*focusDistance)
	// lower_left_corner = origin - horizontal/2 - vertical/2 - focusDistance*w;
	lowerLeftCorner := SubtractVectors(
		SubtractVectors(
			SubtractVectors(
				origin,
				DivideScalar(horizontal, 2),
			),
			DivideScalar(vertical, 2),
		),
		MultiplyScalar(w, focusDistance),
	)

	return Camera{origin, lowerLeftCorner, horizontal, vertical, u, v, w, aperture / 2}
}

func (c Camera) GetRay(s, t float64) Ray {
	rd := MultiplyScalar(RandomVectorInUnitDisk(), c.LensRadius)
	offset := AddVectors(
		MultiplyScalar(c.U, rd.X),
		MultiplyScalar(c.V, rd.Y),
	)

	direction := SubtractVectors(
		SubtractVectors(
			AddVectors(
				AddVectors(
					c.LowerLeftCorner,
					MultiplyScalar(c.Horizontal, s),
				),
				MultiplyScalar(c.Vertical, t),
			),
			c.Origin,
		),
		offset,
	)

	return Ray{
		AddVectors(c.Origin, offset),
		direction,
	}
}
