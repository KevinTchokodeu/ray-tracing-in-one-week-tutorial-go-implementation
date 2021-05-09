package geom

// Defines a ray with origin and direction
type Ray struct {
	Or  Vec3
	Dir Unit
}

// Creates a new ray given an origin and a direction
func NewRay(origin Vec3, direction Unit) Ray {
	return Ray{Or: origin, Dir: direction}
}

// Returns the ray at point t
func (r Ray) At(t float64) Vec3 {
	return r.Or.Plus(r.Dir.Scaled(t))
}
