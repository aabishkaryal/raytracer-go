package models

type Ray struct {
	Origin    Point3
	Direction Vec3
}

func (r Ray) At(t float64) Point3 {
	return AddVectors(r.Origin, MultiplyScalar(r.Direction, t))
}
