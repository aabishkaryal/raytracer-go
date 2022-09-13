package main

type HitRecord struct {
	P         Point3
	Normal    Vec3
	T         float64
	FrontFace bool
}

func (h *HitRecord) SetFaceNormal(r Ray, outwardNormal Vec3) {
	h.FrontFace = VectorDot(r.Direction, outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Negative()
	}
}

type Hittable interface {
	Hit(r Ray, tMin float64, tMax float64) (bool, HitRecord)
}

type HittableList struct {
	Objects []Hittable
}

func (h *HittableList) Clear() {
	h.Objects = []Hittable{}
}

func (h *HittableList) Add(object Hittable) {
	h.Objects = append(h.Objects, object)
}

func (h HittableList) Hit(r Ray, tMin float64, tMax float64) (bool, HitRecord) {
	var rec HitRecord
	hitAnything := false
	closestSoFar := tMax

	for _, object := range h.Objects {
		hit, tempRec := object.Hit(r, tMin, closestSoFar)
		if hit {
			hitAnything = true
			closestSoFar = tempRec.T
			rec = tempRec
		}
	}
	return hitAnything, rec
}
