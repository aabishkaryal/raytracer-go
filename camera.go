package main

type Camera struct {
	Origin          Point3
	LowerLeftCorner Point3
	Horizontal      Vec3
	Vertical        Vec3
}

func NewCamera() Camera {
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

	return Camera{origin, lowerLeftCorner, horizontal, vertical}
}

func (c Camera) GetRay(u, v float64) Ray {
	direction := SubtractVectors(
		AddVectors(
			AddVectors(c.LowerLeftCorner,
				MultiplyScalar(c.Horizontal, u)),
			MultiplyScalar(c.Vertical, v)),
		c.Origin)
	return Ray{c.Origin, direction}
}
