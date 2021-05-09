package trace

import "ray-tracing-in-one-week-tutorial-go-implementation/pkg/geom"

// List of Surfaces
type List struct {
	HH []Hittable
}

// Creates a new list of Surfaces
func NewList(s ...Hittable) List {
	return List{HH: s}
}

// Finds the first intersection (if any) between Ray r and any of the Surfaces in the List.
// If no intersection is found, t = 0.
func (l List) Hit(r geom.Ray, tMin, tMax float64) (t float64, s Surfacer) {
	closest := tMax
	for _, h := range l.HH {
		if ht, hs := h.Hit(r, tMin, closest); ht > 0 {
			closest, t = ht, ht
			s = hs
		}
	}
	return
}
