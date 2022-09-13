package main

import "math"

type Material interface {
	Scatter(rIn Ray, rec HitRecord) (bool, Ray, Color)
}

// Lambertian Material
type Lambertian struct {
	Albedo Color
}

func (l Lambertian) Scatter(rIn Ray, rec HitRecord) (bool, Ray, Color) {
	scatterDirection := AddVectors(rec.Normal, RandomUnitVector())
	if scatterDirection.NearZero() {
		scatterDirection = rec.Normal
	}
	scattered := Ray{rec.P, scatterDirection}
	attenuation := l.Albedo
	return true, scattered, attenuation
}

// Metal Material
type Metal struct {
	Albedo Color
	Fuzz   float64
}

func (m Metal) Scatter(rIn Ray, rec HitRecord) (bool, Ray, Color) {
	reflected := Reflect(UnitVector(rIn.Direction), rec.Normal)
	scattered := Ray{
		rec.P,
		AddVectors(reflected, MultiplyScalar(RandomVectorInUnitSphere(), m.Fuzz)),
	}
	attenuation := m.Albedo
	return (VectorDot(scattered.Direction, rec.Normal) > 0), scattered, attenuation
}

// Dielectric Material
type Dielectric struct {
	RefIdx float64
}

func (d Dielectric) Scatter(rIn Ray, rec HitRecord) (bool, Ray, Color) {
	attenuation := Color{1.0, 1.0, 1.0}
	refractionRatio := d.RefIdx
	if rec.FrontFace {
		refractionRatio = 1.0 / d.RefIdx
	}
	unitDirection := UnitVector(rIn.Direction)

	cosTheta := math.Min(VectorDot(unitDirection.Negative(), rec.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	var direction Vec3

	if refractionRatio*sinTheta > 1.0 {
		direction = Reflect(unitDirection, rec.Normal)
	} else {
		direction = Refract(unitDirection, rec.Normal, refractionRatio)
	}

	scattered := Ray{rec.P, direction}
	return true, scattered, attenuation
}
