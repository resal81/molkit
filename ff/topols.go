package ff

// --------------------------------------------------------

type TopAtom struct {
	serial   int64
	name     string
	atype    string
	charge   float32
	mass     float32
	fragment *TopFragment
}

// --------------------------------------------------------

type TopFragment struct {
	serial int64
	name   string
	atoms  []*TopAtom
}

// --------------------------------------------------------

type TopPolymer struct {
	atoms       []*TopAtom
	fragments   []*TopFragment
	bonds       []*TopBond
	angles      []*TopAngle
	dihedrals   []*TopDihedral
	impropers   []*TopDihedral
	pairs       []*TopPair
	rest_pos    []*TopPositionRestraint
	rest_dist   []*TopDistanceRestraint
	rest_ang    []*TopAngleRestraint
	rest_dih    []*TopDihedralRestraint
	rest_orient []*TopOrientationRestraint
	exclusions  []*TopExclusion
	settle      *TopSettle
}

// --------------------------------------------------------

type TopBond struct {
}

// --------------------------------------------------------

type TopAngle struct {
}

// --------------------------------------------------------

type TopDihedral struct {
}

// --------------------------------------------------------

type TopImproper struct {
}

// --------------------------------------------------------

type TopConstraint struct {
}

// --------------------------------------------------------

type TopPair struct {
}

// --------------------------------------------------------

type TopExclusion struct {
	atoms []*TopAtom
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
