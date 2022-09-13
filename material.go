package main

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
}

func (m Metal) Scatter(rIn Ray, rec HitRecord) (bool, Ray, Color) {
	reflected := Reflect(UnitVector(rIn.Direction), rec.Normal)
	scattered := Ray{rec.P, reflected}
	attenuation := m.Albedo
	return (VectorDot(scattered.Direction, rec.Normal) > 0), scattered, attenuation
}
