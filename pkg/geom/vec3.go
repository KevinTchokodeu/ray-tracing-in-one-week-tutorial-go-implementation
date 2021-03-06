package geom

import (
	"fmt"
	"io"
	"math"
	"math/rand"
)

// 3-element vector
type Vec3 struct {
	El [3]float64
}

// constructor
func NewVec(el0, el1, el2 float64) (v Vec3) {
	return Vec3{El: [3]float64{el0, el1, el2}}
}

// RandVecInSphere creates a random Vec within a unit sphere
// TODO: I don't like rejection methods. Isn't there a way to generate 2 angles and accomplish the same thing reliably?
func RandVecInSphere() Vec3 {
	for {
		v := NewVec(rand.Float64(), rand.Float64(), rand.Float64()).Scaled(2).Minus(NewVec(1, 1, 1))
		if v.LenSq() < 1 {
			return v
		}
	}
}

// RandVecInDisk creates a random Vec within a unit disk
// TODO: more rejection methods :/
func RandVecInDisk() Vec3 {
	xy := NewVec(1, 1, 0)
	for {
		v := NewVec(rand.Float64(), rand.Float64(), 0).Scaled(2).Minus(xy)
		if v.Dot(v) < 1 {
			return v
		}
	}
}

// Return first element
func (v Vec3) X() float64 {
	return v.El[0]
}

// Return second element
func (v Vec3) Y() float64 {
	return v.El[1]
}

// Return third element
func (v Vec3) Z() float64 {
	return v.El[2]
}

// Return vector's inverse as a new vector
func (v Vec3) Inv() Vec3 {
	return NewVec(-v.El[0], -v.El[1], -v.El[2])
}

// Return the square of the vector's length
func (v Vec3) LenSq() float64 {
	return v.El[0]*v.El[0] + v.El[1]*v.El[1] + v.El[2]*v.El[2]
}

// Vector's lenght
func (v Vec3) Len() float64 {
	return math.Sqrt(v.LenSq())
}

// Unit converts vector to a unit vector
func (v Vec3) Unit() (u Unit) {
	k := 1.0 / v.Len()
	u.El[0] = v.El[0] * k
	u.El[1] = v.El[1] * k
	u.El[2] = v.El[2] * k
	return
}

// streams in space-separated vector values from a Reader
func (v Vec3) IStream(r io.Reader) error {
	_, err := fmt.Fscan(r, v.El[0], v.El[1], v.El[2])
	return err
}

// writes space-separated vector values from a Reader
func (v Vec3) OStream(w io.Writer) error {
	_, err := fmt.Fprint(w, v.El[0], v.El[1], v.El[2])
	return err
}

// Sum of two vectors
func (v Vec3) Plus(v1 Vec3) Vec3 {
	return NewVec(v.El[0]+v1.El[0], v.El[1]+v1.El[1], v.El[2]+v1.El[2])
}

// Difference of two vectors
func (v Vec3) Minus(v1 Vec3) Vec3 {
	return NewVec(v.El[0]-v1.El[0], v.El[1]-v1.El[1], v.El[2]-v1.El[2])
}

// Multiplication of two vectors
func (v Vec3) Times(v1 Vec3) Vec3 {
	return NewVec(v.El[0]*v1.El[0], v.El[1]*v1.El[1], v.El[2]*v1.El[2])
}

// Division of two vectors
func (v Vec3) Div(v1 Vec3) Vec3 {
	return NewVec(v.El[0]/v1.El[0], v.El[1]/v1.El[1], v.El[2]/v1.El[2])
}

// Vector scaled by a scalar
func (v Vec3) Scaled(n float64) Vec3 {
	return NewVec(v.El[0]*n, v.El[1]*n, v.El[2]*n)
}

// Dot product of two vectors
func (v Vec3) Dot(v1 Vec3) float64 {
	return v.El[0]*v1.El[0] + v.El[1]*v1.El[1] + v.El[2]*v1.El[2]
}

// Returns the cross product of two vectors
func (v Vec3) Cross(v1 Vec3) Vec3 {
	return NewVec(
		v.El[1]*v1.El[2]-v.El[2]*v1.El[1],
		v.El[2]*v1.El[0]-v.El[0]*v1.El[2],
		v.El[0]*v1.El[1]-v.El[1]*v1.El[0],
	)
}
