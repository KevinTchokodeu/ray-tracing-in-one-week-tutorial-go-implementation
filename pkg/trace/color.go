package trace

import "ray-tracing-in-one-week-tutorial-go-implementation/pkg/geom"

// Represents an RGB color value
type Color struct {
	geom.Vec3
}

// Creates a Color from 3 float values
func NewColor(e0, e1, e2 float64) (c Color) {
	c.El[0] = e0
	c.El[1] = e1
	c.El[2] = e2
	return
}

// Returns the first element (Red element)
func (c Color) R() float64 {
	return c.El[0]
}

// Returns the second element (Green element)
func (c Color) G() float64 {
	return c.El[1]
}

// Returns the third element (Blue element)
func (c Color) B() float64 {
	return c.El[2]
}

// Returns sum of two colors
func (c Color) Plus(c2 Color) Color {
	return Color{Vec3: c.Vec3.Plus(c2.Vec3)}
}

// Returns the color scaled
func (c Color) Scaled(n float64) Color {
	return Color{Vec3: c.Vec3.Scaled(n)}
}
