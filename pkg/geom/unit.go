package geom

// unit vector (length 1)
type Unit struct {
	Vec3
}

// Dot returns the dot product of two unit vectors.
// Values above zero indicate vectors pointing in the same hemisphere.
// Values below zero indicate vectors pointing towards opposite hemispheres.
// TODO: check to see if that's generally true, or only for unit vectors
func (u Unit) Dot(u2 Unit) float64 {
	return u.Vec3.Dot(u2.Vec3)
}

// Inv inverts the unit vector.
func (u Unit) Inv() Unit {
	return Unit{Vec3: u.Vec3.Inv()}
}
