package trace

import "ray-tracing-in-one-week-tutorial-go-implementation/pkg/geom"

// Lambert describes a diffuse material.
type Lambert struct {
	Albedo Color
}

// Scatter scatters incoming light rays in a hemisphere about the normal.
func (l Lambert) Scatter(in geom.Ray, p geom.Vec3, n geom.Unit) (out geom.Ray, attenuation Color, ok bool) {
	target := p.Plus(n.Vec3).Plus(geom.RandVecInSphere())
	out = geom.NewRay(p, target.Minus(p).ToUnit())
	return out, l.Albedo, true
}
