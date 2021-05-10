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
	u, v, w    geom.Unit
	lensRadius float64
}

func NewCamera(lookFrom, lookAt geom.Vec3, vup geom.Unit, vfov, aspect, aperture, focus float64) (c Camera) {
	theta := vfov * math.Pi / 180
	halfH := math.Tan(theta / 2)
	halfW := aspect * halfH

	c.w = lookFrom.Minus(lookAt).Unit()
	c.u = vup.Cross(c.w.Vec3).Unit()
	c.v = c.w.Cross(c.u.Vec3).Unit()

	width := c.u.Scaled(halfW * focus)
	height := c.v.Scaled(halfH * focus)
	dist := c.w.Scaled(focus)

	c.lensRadius = aperture / 2
	c.origin = lookFrom
	c.lowerLeft = c.origin.Minus(width).Minus(height).Minus(dist)
	c.horizontal = width.Scaled(2)
	c.vertical = height.Scaled(2)
	return
}

// Ray returns a Ray passing through a given s, t coordinate.
func (c Camera) Ray(s, t float64) geom.Ray {
	rd := geom.RandVecInDisk().Scaled(c.lensRadius)
	offset := c.u.Scaled(rd.X()).Plus(c.v.Scaled(rd.Y()))
	source := c.origin.Plus(offset)
	dest := c.lowerLeft.Plus(c.horizontal.Scaled(s).Plus(c.vertical.Scaled(t)))
	return geom.NewRay(
		source,
		dest.Minus(source).Unit(),
	)
}
