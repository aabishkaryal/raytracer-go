package main

import "math"

type Camera struct {
	Origin          Point3
	LowerLeftCorner Point3
	Horizontal      Vec3
	Vertical        Vec3
}

func NewCamera(lookFrom, lookAt Point3, vup Vec3, verticalFOV, aspectRatio float64) Camera {
	theta := DegreesToRadians(verticalFOV)
	h := math.Tan(theta / 2)
	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := UnitVector(SubtractVectors(lookFrom, lookAt))
	u := UnitVector(VectorCross(vup, w))
	v := VectorCross(w, u)

	origin := lookFrom
	horizontal := MultiplyScalar(u, viewportWidth)
	vertical := MultiplyScalar(v, viewportHeight)
	// lower_left_corner = origin - horizontal/2 - vertical/2 - w;
	lowerLeftCorner := SubtractVectors(
		SubtractVectors(
			SubtractVectors(
				origin,
				DivideScalar(horizontal, 2),
			),
			DivideScalar(vertical, 2),
		),
		w,
	)

	return Camera{origin, lowerLeftCorner, horizontal, vertical}
}

func (c Camera) GetRay(s, t float64) Ray {
	direction := SubtractVectors(
		AddVectors(
			AddVectors(
				c.LowerLeftCorner,
				MultiplyScalar(c.Horizontal, s),
			),
			MultiplyScalar(c.Vertical, t),
		),
		c.Origin,
	)
	return Ray{c.Origin, direction}
}
