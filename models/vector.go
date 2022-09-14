package models

import (
	"fmt"
	"io"
	"math"
	"math/rand"

	"github.com/aabishkaryal/raytracer-go/utils"
)

// Equivalent of the vec3 class

type Vec3 struct {
	X, Y, Z float64
}

// Negative returns the negative equivalent of the vector
func (v Vec3) Negative() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

// Add adds the given vector to this vector in place
func (v *Vec3) Add(v2 Vec3) {
	v.X += v2.X
	v.Y += v2.Y
	v.Z += v2.Z
}

// Multiply multiplies this vector by the given scalar in place
func (v *Vec3) Multiply(s float64) {
	v.X *= s
	v.Y *= s
	v.Z *= s
}

// Divide divides this vector by the given scalar in place
func (v *Vec3) Divide(s float64) {
	v.X /= s
	v.Y /= s
	v.Z /= s
}

// Length returns the length of the vector
func (v Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// LengthSquared returns the squared length of the vector
func (v Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// At returns the value of the vector at the given index
func (v Vec3) At(i int) float64 {
	switch i {
	case 0:
		return v.X
	case 1:
		return v.Y
	case 2:
		return v.Z
	default:
		panic("index out of range")
	}
}

// Vec3 Utility functions

// Print prints the vector to the console
func (v Vec3) Print(out io.Writer) {
	fmt.Fprintf(out, "%v %v %v", v.X, v.Y, v.Z)
}

// Returns true if the vector is close to zero in all dimensions
func (v Vec3) NearZero() bool {
	s := 1e-8
	return (math.Abs(v.X) < s) && (math.Abs(v.Y) < s) && (math.Abs(v.Z) < s)
}

// AddVectors returns the sum of the two vectors
func AddVectors(v1, v2 Vec3) Vec3 {
	return Vec3{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

// SubtractVectors returns the difference of the two vectors
func SubtractVectors(v1, v2 Vec3) Vec3 {
	return Vec3{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

// MultiplyVectors returns the product of the two vectors
func MultiplyVectors(v1, v2 Vec3) Vec3 {
	return Vec3{v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z}
}

// MultiplyScalar returns the product of the vector and the scalar
func MultiplyScalar(v Vec3, s float64) Vec3 {
	return Vec3{v.X * s, v.Y * s, v.Z * s}
}

// DivideScalar returns the quotient of the vector and the scalar
func DivideScalar(v Vec3, s float64) Vec3 {
	return Vec3{v.X / s, v.Y / s, v.Z / s}
}

// VectorDot returns the dot product of the two vectors
func VectorDot(v1, v2 Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// VectorCross returns the cross product of the two vectors
func VectorCross(v1, v2 Vec3) Vec3 {
	return Vec3{v1.Y*v2.Z - v1.Z*v2.Y, v1.Z*v2.X - v1.X*v2.Z, v1.X*v2.Y - v1.Y*v2.X}
}

// UnitVector returns the unit vector of the given vector
func UnitVector(v Vec3) Vec3 {
	return DivideScalar(v, v.Length())
}

// RandomVector returns a random vector in the unit cube
func RandomVector() Vec3 {
	return Vec3{rand.Float64(), rand.Float64(), rand.Float64()}
}

// RandomVectorRange returns a random vector in the given range
func RandomVectorRange(min, max float64) Vec3 {
	return Vec3{utils.RandomRange(min, max), utils.RandomRange(min, max), utils.RandomRange(min, max)}
}

// RandomVectorInUnitSphere returns a random vector in the unit sphere
func RandomVectorInUnitSphere() Vec3 {
	for {
		p := RandomVectorRange(-1, 1)
		if p.LengthSquared() >= 1 {
			continue
		}
		return p
	}
}

// RandomVectorInUnitDisk returns a random unit vector in the unit sphere
func RandomUnitVector() Vec3 {
	return UnitVector(RandomVectorInUnitSphere())
}

// RandomVectorInHemisphere returns a random vector in the hemisphere
func RandomVectorInHemisphere(normal Vec3) Vec3 {
	inUnitSphere := RandomVectorInUnitSphere()
	if VectorDot(inUnitSphere, normal) > 0.0 {
		return inUnitSphere
	}
	return inUnitSphere.Negative()
}

// RandomVectorInUnitDisk returns a random vector in the unit disk
func RandomVectorInUnitDisk() Vec3 {
	for {
		p := Vec3{utils.RandomRange(-1, 1), utils.RandomRange(-1, 1), 0}
		if p.LengthSquared() >= 1 {
			continue
		}
		return p
	}
}

// Reflect returns the reflection of the vector about the normal
func Reflect(v, n Vec3) Vec3 {
	return SubtractVectors(v, MultiplyScalar(n, 2*VectorDot(v, n)))
}

// Refract returns the refraction of the vector about the normal
func Refract(uv, n Vec3, etaiOverEtat float64) Vec3 {
	cosTheta := math.Min(VectorDot(uv.Negative(), n), 1.0)
	rOutPerp := MultiplyScalar(
		AddVectors(
			uv,
			MultiplyScalar(n, cosTheta),
		),
		etaiOverEtat,
	)
	rOutParallel := MultiplyScalar(
		n,
		-math.Sqrt(math.Abs(1.0-rOutPerp.LengthSquared())),
	)
	return AddVectors(rOutPerp, rOutParallel)
}
