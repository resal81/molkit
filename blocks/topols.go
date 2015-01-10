package blocks

type TopAtom struct {
}

type TopFragment struct {
}

type TopPolymer struct {
}

// --------------------------------------------------------

type TopConstraint struct {
}

// --------------------------------------------------------

type TopPair struct {
}

// --------------------------------------------------------

type TopExclusion struct {
	atoms []*Atom
}

// --------------------------------------------------------

type TopSettle struct {
	d_OH float32
	d_HH float32
}

// --------------------------------------------------------

type TopPositionRestraint struct {
}

// --------------------------------------------------------

type TopDistanceRestraint struct {
}

// --------------------------------------------------------

type TopAngleRestraint struct {
}

// --------------------------------------------------------

type TopDihedralRestraint struct {
}

// --------------------------------------------------------

type TopOrientationRestraint struct {
}
