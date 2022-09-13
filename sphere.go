package main

import "math"

type Sphere struct {
	Center Point3
	Radius float64
	MatPtr Material
}

func (s Sphere) Hit(r Ray, tMin float64, tMax float64) (bool, HitRecord) {
	oc := SubtractVectors(r.Origin, s.Center)
	a := VectorDot(r.Direction, r.Direction)
	halfB := VectorDot(oc, r.Direction)
	c := VectorDot(oc, oc) - s.Radius*s.Radius
	discriminant := halfB*halfB - a*c

	if discriminant < 0 {
		return false, HitRecord{}
	}

	sqrtD := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range.
	root := (-halfB - sqrtD) / a
	if root < tMin || tMax < root {
		root = (-halfB + sqrtD) / a
		if root < tMin || tMax < root {
			return false, HitRecord{}
		}
	}

	rec := HitRecord{
		T:      root,
		P:      r.At(root),
		Normal: DivideScalar(SubtractVectors(r.At(root), s.Center), s.Radius),
		MatPtr: s.MatPtr,
	}
	rec.SetFaceNormal(r, rec.Normal)

	return true, rec
}
