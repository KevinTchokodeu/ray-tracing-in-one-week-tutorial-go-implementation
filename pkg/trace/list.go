package trace

import "ray-tracing-in-one-week-tutorial-go-implementation/pkg/geom"

// List of Surfaces
type List struct {
	SS []Surface
}

// Creates a new list of Surfaces
func NewList(s ...Surface) List {
	return List{SS: s}
}

// Finds the first intersection (if any) between Ray r and any of the Surfaces in the List.
// If no intersection is found, t = 0.
func (l List) Hit(r geom.Ray, tMin, tMax float64) (t float64, p geom.Vec3, n geom.Unit) {
	closest := tMax
	for _, s := range l.SS {
		if st, sp, sn := s.Hit(r, tMin, closest); st > 0 {
			closest, t = st, st
			p = sp
			n = sn
		}
	}
	return
}
