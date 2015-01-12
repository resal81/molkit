package ff

/**********************************************************
* TopAtom
**********************************************************/

type topAtomSetting int32

const (
	TOP_ATOM_NAME_SET topAtomSetting = 1 << iota
	TOP_ATOM_TYPE_SET
	TOP_ATOM_SERIAL_SET
	TOP_ATOM_CHARGE_SET
	TOP_ATOM_MASS_SET
	TOP_ATOM_FRAGMENT_SET
	TOP_ATOM_CGNR_SET
)

type TopAtom struct {
	name     string
	atype    string
	serial   int64
	charge   float64
	mass     float64
	cgnr     int64
	fragment *TopFragment
	setting  topAtomSetting
}

func NewTopAtom() *TopAtom {
	return &TopAtom{}
}

//
func (a *TopAtom) SetName(name string) {
	a.setting |= TOP_ATOM_NAME_SET
	a.name = name
}

func (a *TopAtom) HasNameSet() bool {
	return a.setting&TOP_ATOM_NAME_SET != 0
}

func (a *TopAtom) Name() string {
	return a.name
}

//
func (a *TopAtom) SetAtomType(atype string) {
	a.setting |= TOP_ATOM_TYPE_SET
	a.atype = atype
}

func (a *TopAtom) HasAtomTypeSet() bool {
	return a.setting&TOP_ATOM_TYPE_SET != 0
}

func (a *TopAtom) AtomType() string {
	return a.atype
}

//
func (a *TopAtom) SetSerial(ser int64) {
	a.setting |= TOP_ATOM_SERIAL_SET
	a.serial = ser
}

func (a *TopAtom) HasSerialSet() bool {
	return a.setting&TOP_ATOM_SERIAL_SET != 0
}

func (a *TopAtom) Serial() int64 {
	return a.serial
}

//
func (a *TopAtom) SetCGNR(cgnr int64) {
	a.setting |= TOP_ATOM_CGNR_SET
	a.cgnr = cgnr
}

func (a *TopAtom) HasCGNRSet() bool {
	return a.setting&TOP_ATOM_CGNR_SET != 0
}

func (a *TopAtom) CGNR() int64 {
	return a.cgnr
}

//
func (a *TopAtom) SetMass(m float64) {
	a.setting |= TOP_ATOM_MASS_SET
	a.mass = m
}

func (a *TopAtom) HasMassSet() bool {
	return a.setting&TOP_ATOM_MASS_SET != 0
}

func (a *TopAtom) Mass() float64 {
	return a.mass
}

//
func (a *TopAtom) SetCharge(ch float64) {
	a.setting |= TOP_ATOM_CHARGE_SET
	a.charge = ch
}

func (a *TopAtom) HasChargeSet() bool {
	return a.setting&TOP_ATOM_CHARGE_SET != 0
}

func (a *TopAtom) Charge() float64 {
	return a.charge
}

//
func (a *TopAtom) setTopFragment(frag *TopFragment) {
	a.setting |= TOP_ATOM_FRAGMENT_SET
	a.fragment = frag
}

func (a *TopAtom) HasFragmentSet() bool {
	return a.setting&TOP_ATOM_FRAGMENT_SET != 0
}

func (a *TopAtom) Fragment() *TopFragment {
	return a.fragment
}

/**********************************************************
* TopFragment
**********************************************************/

type topFragmentSetting int32

const (
	TOP_FRAGMENT_NAME_SET topFragmentSetting = 1 << iota
	TOP_FRAGMENT_SERIAL_SET
	TOP_FRAGMENT_POLYMER_SET
)

type TopFragment struct {
	serial  int64
	name    string
	atoms   []*TopAtom
	polymer *TopPolymer
	setting topFragmentSetting
}

func NewTopFragment() *TopFragment {
	return &TopFragment{}
}

//
func (f *TopFragment) SetName(name string) {
	f.setting |= TOP_FRAGMENT_NAME_SET
	f.name = name
}

func (f *TopFragment) HasNameSet() bool {
	return f.setting&TOP_FRAGMENT_NAME_SET != 0
}

func (f *TopFragment) Name() string {
	return f.name
}

//
func (f *TopFragment) SetSerial(ser int64) {
	f.setting |= TOP_FRAGMENT_SERIAL_SET
	f.serial = ser
}

func (f *TopFragment) HasSerialSet() bool {
	return f.setting&TOP_FRAGMENT_SERIAL_SET != 0
}

func (f *TopFragment) Serial() int64 {
	return f.serial
}

//
func (f *TopFragment) setTopPolymer(pol *TopPolymer) {
	f.setting |= TOP_FRAGMENT_POLYMER_SET
	f.polymer = pol
}

func (f *TopFragment) HasTopPolymerSet() bool {
	return f.setting&TOP_FRAGMENT_POLYMER_SET != 0
}

func (f *TopFragment) TopPolymer() *TopPolymer {
	return f.polymer
}

//
func (f *TopFragment) AddTopAtom(a *TopAtom) {
	f.atoms = append(f.atoms, a)
	a.setTopFragment(f)
}

func (f *TopFragment) TopAtoms() []*TopAtom {
	return f.atoms
}

/**********************************************************
* TopPolymer
**********************************************************/

type topPolymerSetting int32

const (
	TOP_POLYMER_NAME_SET topPolymerSetting = 1 << iota
)

type TopPolymer struct {
	atoms       []*TopAtom
	atomsMap    map[int64]*TopAtom
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

	name    string
	nrexcl  int8
	setting topPolymerSetting
}

//
func NewTopPolymer() *TopPolymer {
	return &TopPolymer{
		atomsMap: map[int64]*TopAtom{},
	}
}

//
func (p *TopPolymer) SetName(name string) {
	if name == "" {
		panic("polymer name cannot be empty")
	}
	p.setting |= TOP_POLYMER_NAME_SET
	p.name = name
}

func (p *TopPolymer) HasNameSet() bool {
	return p.setting&TOP_POLYMER_NAME_SET != 0
}

func (p *TopPolymer) Name() string {
	return p.name
}

//
func (p *TopPolymer) SetNrExcl(nrexcl int8) {
	p.nrexcl = nrexcl
}

func (p *TopPolymer) NrExcl() int8 {
	return p.nrexcl
}

//
func (p *TopPolymer) AddTopFragment(f *TopFragment) {
	p.fragments = append(p.fragments, f)
	f.setTopPolymer(p)
}

func (p *TopPolymer) Fragments() []*TopFragment {
	return p.fragments
}

//
func (p *TopPolymer) AddTopAtom(a *TopAtom) {
	if !a.HasSerialSet() {
		panic("atom doesn't have serial set")
	}

	if _, ok := p.atomsMap[a.Serial()]; ok {
		panic("atom with the same serial exist in the map")
	}

	p.atomsMap[a.Serial()] = a
	p.atoms = append(p.atoms, a)
}

func (p *TopPolymer) Atoms() []*TopAtom {
	return p.atoms
}

/**********************************************************
* TopSystem
**********************************************************/

type TopSystem struct {
	polymersMap map[string]*TopPolymer // unique polymers the build the system
	polymers    []*TopPolymer          // all polymers
}

func NewTopSystem() *TopSystem {
	return &TopSystem{
		polymersMap: map[string]*TopPolymer{},
	}
}

func (s *TopSystem) RegisterTopPolymer(p *TopPolymer) {
	if !p.HasNameSet() {
		panic("polymer name must be set")
	}

	if _, found := s.polymersMap[p.Name()]; found {
		panic("a polymer with the same name has been already registered")
	}
	s.polymersMap[p.Name()] = p
}

func (s *TopSystem) AddTopPolymer(p *TopPolymer) {
	s.polymers = append(s.polymers, p)
}

func (s *TopSystem) RegisteredTopPolymers() map[string]*TopPolymer {
	return s.polymersMap
}

func (s *TopSystem) TopPolymers() []*TopPolymer {
	return s.polymers
}

/**********************************************************
* TopBond
**********************************************************/

type TopBond struct {
	atom1, atom2 *TopAtom
	bondType     *BondType
}

func NewTopBond(atom1, atom2 *TopAtom) *TopBond {
	return &TopBond{
		atom1: atom1,
		atom2: atom2,
	}
}

/**********************************************************
* TopPair
**********************************************************/

type TopPair struct {
	atom1, atom2 *TopAtom
	pairType     *PairType
}

func NewTopPair(atom1, atom2 *TopAtom) *TopPair {
	return &TopPair{
		atom1: atom1,
		atom2: atom2,
	}
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
