package trace

import (
	"math"
	"ray-tracing-in-one-week-tutorial-go-implementation/pkg/geom"
)

var (
	lowerLeft  = geom.NewVec(-2, -1, -1)
	horizontal = geom.NewVec(4, 0, 0)
	vertical   = geom.NewVec(0, 2, 0)
	origin     = geom.NewVec(0, 0, 0)
)

// Camera originates Rays.
type Camera struct {
	lowerLeft  geom.Vec3
	horizontal geom.Vec3
	vertical   geom.Vec3
	origin     geom.Vec3
}

func NewCamera(lookFrom, lookAt geom.Vec3, vup geom.Unit, vfov, aspect float64) (c Camera) {
	theta := vfov * math.Pi / 180
	halfH := math.Tan(theta / 2)
	halfW := aspect * halfH

	w := lookFrom.Minus(lookAt).Unit()
	u := vup.Cross(w.Vec3).Unit()
	v := w.Cross(u.Vec3).Unit()

	c.origin = lookFrom
	c.lowerLeft = c.origin.Minus(u.Scaled(halfW)).Minus(v.Scaled(halfH)).Minus(w.Vec3)
	c.horizontal = u.Scaled(2 * halfW)
	c.vertical = v.Scaled(2 * halfH)
	return
}

// Ray returns a Ray passing through a given u, v coordinate.
func (c Camera) Ray(u, v float64) geom.Ray {
	return geom.NewRay(
		c.origin,
		c.lowerLeft.Plus((c.horizontal.Scaled(u)).Plus(c.vertical.Scaled(v))).Minus(c.origin).Unit(),
	)
}
