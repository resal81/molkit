package ff

/**********************************************************
* TopAtom
**********************************************************/

type TopAtom struct {
	name     string
	atype    string
	serial   int64
	charge   float64
	mass     float64
	fragment *TopFragment
}

func NewTopAtom() *TopAtom {
	return &TopAtom{}
}

//
func (a *TopAtom) SetName(name string) {
	a.name = name
}

func (a *TopAtom) Name() string {
	return a.name
}

//
func (a *TopAtom) SetAtomType(atype string) {
	a.atype = atype
}

func (a *TopAtom) AtomType() string {
	return a.atype
}

//
func (a *TopAtom) SetSerial(ser int64) {
	a.serial = ser
}

func (a *TopAtom) Serial() int64 {
	return a.serial
}

//
func (a *TopAtom) SetMass(m float64) {
	a.mass = m
}

func (a *TopAtom) Mass() float64 {
	return a.mass
}

//
func (a *TopAtom) SetCharge(ch float64) {
	a.charge = ch
}

func (a *TopAtom) Charge() float64 {
	return a.charge
}

//
func (a *TopAtom) SetTopFragment(frag *TopFragment) {
	a.fragment = frag
}

func (a *TopAtom) Fragment() *TopFragment {
	return a.fragment
}

/**********************************************************
* TopFragment
**********************************************************/

type TopFragment struct {
	serial int64
	name   string
	atoms  []*TopAtom
}

/**********************************************************
* TopPolymer
**********************************************************/

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

/**********************************************************
* TopSystem
**********************************************************/

type TopSystem struct {
	polymers []*TopPolymer
}

/**********************************************************
* TopBond
**********************************************************/

type TopBond struct {
	atom1, atom2 *TopAtom
	bondType     *BondType
}

/**********************************************************
* TopPair
**********************************************************/

type TopPair struct {
	atom1, atom2 *TopAtom
	pairType     *PairType
}

/**********************************************************
* TopAngle
**********************************************************/

type TopAngle struct {
	atom1, atom2, atom3 *TopAtom
	angleType           *AngleType
}

/**********************************************************
* TopDihedral
**********************************************************/

type TopDihedral struct {
	atom1, atom2, atom3, atom4 *TopAtom
	dihedralType               *DihedralType
}

/**********************************************************
* TopConstraint
**********************************************************/

type TopConstraint struct {
	constraintType *ConstraintType
}

/**********************************************************
* TopExclusion
**********************************************************/

type TopExclusion struct {
	atoms []*TopAtom
}

/**********************************************************
* TopSettle
**********************************************************/

type TopSettle struct {
	d_OH float64
	d_HH float64
}

/**********************************************************
* TopPositionRestraint
**********************************************************/

type TopPositionRestraint struct {
}

/**********************************************************
* TopDistanceRestraint
**********************************************************/

type TopDistanceRestraint struct {
}

/**********************************************************
* TopAngleRestraint
**********************************************************/

type TopAngleRestraint struct {
}

/**********************************************************
* TopDihedralRestraint
**********************************************************/

type TopDihedralRestraint struct {
}

/**********************************************************
* TopOrientationRestraint
**********************************************************/

type TopOrientationRestraint struct {
}
