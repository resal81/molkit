package ff

/**********************************************************
* TopAtom
**********************************************************/

type topAtomSetting int32

const (
	t_at_sett_NAME_SET topAtomSetting = 1 << iota
	t_at_sett_TYPE_SET
	t_at_sett_SERIAL_SET
	t_at_sett_CHARGE_SET
	t_at_sett_MASS_SET
	t_at_sett_FRAGMENT_SET
	t_at_sett_CGNR_SET
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

	posrest *TopPositionRestraint
}

func NewTopAtom() *TopAtom {
	return &TopAtom{}
}

//
func (a *TopAtom) SetName(name string) {
	a.setting |= t_at_sett_NAME_SET
	a.name = name
}

func (a *TopAtom) HasNameSet() bool {
	return a.setting&t_at_sett_NAME_SET != 0
}

func (a *TopAtom) Name() string {
	return a.name
}

//
func (a *TopAtom) SetAtomType(atype string) {
	a.setting |= t_at_sett_TYPE_SET
	a.atype = atype
}

func (a *TopAtom) HasAtomTypeSet() bool {
	return a.setting&t_at_sett_TYPE_SET != 0
}

func (a *TopAtom) AtomType() string {
	return a.atype
}

//
func (a *TopAtom) SetSerial(ser int64) {
	a.setting |= t_at_sett_SERIAL_SET
	a.serial = ser
}

func (a *TopAtom) HasSerialSet() bool {
	return a.setting&t_at_sett_SERIAL_SET != 0
}

func (a *TopAtom) Serial() int64 {
	return a.serial
}

//
func (a *TopAtom) SetCGNR(cgnr int64) {
	a.setting |= t_at_sett_CGNR_SET
	a.cgnr = cgnr
}

func (a *TopAtom) HasCGNRSet() bool {
	return a.setting&t_at_sett_CGNR_SET != 0
}

func (a *TopAtom) CGNR() int64 {
	return a.cgnr
}

//
func (a *TopAtom) SetMass(m float64) {
	a.setting |= t_at_sett_MASS_SET
	a.mass = m
}

func (a *TopAtom) HasMassSet() bool {
	return a.setting&t_at_sett_MASS_SET != 0
}

func (a *TopAtom) Mass() float64 {
	return a.mass
}

//
func (a *TopAtom) SetCharge(ch float64) {
	a.setting |= t_at_sett_CHARGE_SET
	a.charge = ch
}

func (a *TopAtom) HasChargeSet() bool {
	return a.setting&t_at_sett_CHARGE_SET != 0
}

func (a *TopAtom) Charge() float64 {
	return a.charge
}

//
func (a *TopAtom) setTopFragment(frag *TopFragment) {
	a.setting |= t_at_sett_FRAGMENT_SET
	a.fragment = frag
}

func (a *TopAtom) HasFragmentSet() bool {
	return a.setting&t_at_sett_FRAGMENT_SET != 0
}

func (a *TopAtom) Fragment() *TopFragment {
	return a.fragment
}

/**********************************************************
* TopFragment
**********************************************************/

type topFragmentSetting int32

const (
	t_frag_sett_NAME_SET topFragmentSetting = 1 << iota
	t_frag_sett_SERIAL_SET
	t_frag_sett_POLYMER_SET
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
	f.setting |= t_frag_sett_NAME_SET
	f.name = name
}

func (f *TopFragment) HasNameSet() bool {
	return f.setting&t_frag_sett_NAME_SET != 0
}

func (f *TopFragment) Name() string {
	return f.name
}

//
func (f *TopFragment) SetSerial(ser int64) {
	f.setting |= t_frag_sett_SERIAL_SET
	f.serial = ser
}

func (f *TopFragment) HasSerialSet() bool {
	return f.setting&t_frag_sett_SERIAL_SET != 0
}

func (f *TopFragment) Serial() int64 {
	return f.serial
}

//
func (f *TopFragment) setTopPolymer(pol *TopPolymer) {
	f.setting |= t_frag_sett_POLYMER_SET
	f.polymer = pol
}

func (f *TopFragment) HasTopPolymerSet() bool {
	return f.setting&t_frag_sett_POLYMER_SET != 0
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
	t_pol_sett_NAME_SET topPolymerSetting = 1 << iota
)

type TopPolymer struct {
	atoms      []*TopAtom
	atomsMap   map[int64]*TopAtom
	fragments  []*TopFragment
	bonds      []*TopBond
	angles     []*TopAngle
	dihedrals  []*TopDihedral
	impropers  []*TopDihedral
	pairs      []*TopPair
	exclusions []*TopExclusion
	settle     *TopSettle

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
	p.setting |= t_pol_sett_NAME_SET
	p.name = name
}

func (p *TopPolymer) HasNameSet() bool {
	return p.setting&t_pol_sett_NAME_SET != 0
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

func (p *TopPolymer) TopAtoms() []*TopAtom {
	return p.atoms
}

// Returns *TopAtom or nil.
func (p *TopPolymer) AtomBySerial(i int64) *TopAtom {
	return p.atomsMap[i]
}

//
func (p *TopPolymer) AddTopBond(b *TopBond) {
	p.bonds = append(p.bonds, b)
}

func (p *TopPolymer) TopBonds() []*TopBond {
	return p.bonds
}

//
func (p *TopPolymer) AddTopPair(v *TopPair) {
	p.pairs = append(p.pairs, v)
}

func (p *TopPolymer) TopPairs() []*TopPair {
	return p.pairs
}

//
func (p *TopPolymer) AddTopAngle(v *TopAngle) {
	p.angles = append(p.angles, v)
}

func (p *TopPolymer) TopAngles() []*TopAngle {
	return p.angles
}

//
func (p *TopPolymer) AddTopDihedral(v *TopDihedral) {
	p.dihedrals = append(p.dihedrals, v)
}

func (p *TopPolymer) TopDihedrals() []*TopDihedral {
	return p.dihedrals
}

//
func (p *TopPolymer) AddTopImproper(v *TopDihedral) {
	p.impropers = append(p.impropers, v)
}

func (p *TopPolymer) TopImpropers() []*TopDihedral {
	return p.impropers
}

//

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

//
func (s *TopSystem) RegisterTopPolymer(p *TopPolymer) {
	if !p.HasNameSet() {
		panic("polymer name must be set")
	}

	if _, found := s.polymersMap[p.Name()]; found {
		panic("a polymer with the same name has been already registered")
	}
	s.polymersMap[p.Name()] = p
}

func (s *TopSystem) RegisteredTopPolymers() map[string]*TopPolymer {
	return s.polymersMap
}

func (s *TopSystem) TopPolymerByName(name string) *TopPolymer {
	return s.polymersMap[name]
}

//
func (s *TopSystem) AddTopPolymer(p *TopPolymer) {
	s.polymers = append(s.polymers, p)
}

func (s *TopSystem) TopPolymers() []*TopPolymer {
	return s.polymers
}

/**********************************************************
* TopBond
**********************************************************/
type topBondSetting int32

const (
	t_bnd_sett_CUSTOM_BOND_TYPE_SET topBondSetting = 1 << iota
	t_bnd_sett_RESTRAINT_SET
)

type TopBond struct {
	atom1 *TopAtom
	atom2 *TopAtom
	kind  ffTypes

	customBondType *BondType
	bondrest       *TopDistanceRestraint
	setting        topBondSetting
}

//
func NewTopBond(atom1, atom2 *TopAtom, kind ffTypes) *TopBond {
	if kind&FF_BOND_TYPE_1 == 0 {
		panic("bond type is not supported")
	}

	return &TopBond{
		atom1: atom1,
		atom2: atom2,
		kind:  kind,
	}

}

//
func (b *TopBond) TopAtom1() *TopAtom {
	return b.atom1
}

func (b *TopBond) TopAtom2() *TopAtom {
	return b.atom2
}

//
func (b *TopBond) SetCustomBondType(bt *BondType) {
	b.setting |= t_bnd_sett_CUSTOM_BOND_TYPE_SET
	b.customBondType = bt
}

func (b *TopBond) HasCustomBondTypeSet() bool {
	return b.setting&t_bnd_sett_CUSTOM_BOND_TYPE_SET != 0
}

func (b *TopBond) CustomBondType() *BondType {
	return b.customBondType
}

/**********************************************************
* TopPair
**********************************************************/

type topPairSetting int32

const (
	t_pair_sett_CUSTOM_PAIR_TYPE_SET topPairSetting = 1 << iota
)

type TopPair struct {
	atom1 *TopAtom
	atom2 *TopAtom
	kind  ffTypes

	customPairType *PairType
	setting        topPairSetting
}

func NewTopPair(atom1, atom2 *TopAtom, kind ffTypes) *TopPair {
	if kind&FF_PAIR_TYPE_1 == 0 {
		panic("pairtype is not supported")
	}
	return &TopPair{
		atom1: atom1,
		atom2: atom2,
		kind:  kind,
	}
}

//
func (p *TopPair) TopAtom1() *TopAtom {
	return p.atom1
}

func (p *TopPair) TopAtom2() *TopAtom {
	return p.atom2
}

//
func (p *TopPair) SetCustomPairType(pt *PairType) {
	p.setting |= t_pair_sett_CUSTOM_PAIR_TYPE_SET
	p.customPairType = pt
}

func (p *TopPair) HasCustomPairTypeSet() bool {
	return p.setting&t_pair_sett_CUSTOM_PAIR_TYPE_SET != 0
}

func (p *TopPair) CustomPairType() *PairType {
	return p.customPairType
}

/**********************************************************
* TopAngle
**********************************************************/

type topAngleSetting int32

const (
	t_ang_sett_CUSTOM_ANGLE_TYPE_SET topAngleSetting = 1 << iota
)

type TopAngle struct {
	atom1 *TopAtom
	atom2 *TopAtom
	atom3 *TopAtom
	kind  ffTypes

	customAngleType *AngleType
	angleRest       *TopAngleRestraint
	setting         topAngleSetting
}

//
func NewTopAngle(atom1, atom2, atom3 *TopAtom, kind ffTypes) *TopAngle {
	if kind&FF_ANGLE_TYPE_1 == 0 && kind&FF_ANGLE_TYPE_5 == 0 {
		panic("angletype is not supported")
	}

	return &TopAngle{
		atom1: atom1,
		atom2: atom2,
		atom3: atom3,
		kind:  kind,
	}

}

//
func (a *TopAngle) TopAtom1() *TopAtom {
	return a.atom1
}

func (a *TopAngle) TopAtom2() *TopAtom {
	return a.atom2
}

func (a *TopAngle) TopAtom3() *TopAtom {
	return a.atom3
}

//
func (a *TopAngle) SetCustomAngleType(at *AngleType) {
	a.setting |= t_ang_sett_CUSTOM_ANGLE_TYPE_SET
	a.customAngleType = at
}

func (a *TopAngle) HasCustomAngleTypeSet() bool {
	return a.setting&t_ang_sett_CUSTOM_ANGLE_TYPE_SET != 0
}

func (a *TopAngle) CustomAngleType() *AngleType {
	return a.customAngleType
}

/**********************************************************
* TopDihedral
**********************************************************/

type topDihedralSetting int32

const (
	t_dih_sett_CUSTOM_DIHEDRAL_TYPE_SET topDihedralSetting = 1 << iota
)

type TopDihedral struct {
	atom1 *TopAtom
	atom2 *TopAtom
	atom3 *TopAtom
	atom4 *TopAtom
	kind  ffTypes

	customDihedralType *DihedralType
	setting            topDihedralSetting
}

//
func NewTopDihedral(atom1, atom2, atom3, atom4 *TopAtom, kind ffTypes) *TopDihedral {
	if kind&FF_DIHEDRAL_TYPE_1 == 0 && kind&FF_DIHEDRAL_TYPE_2 == 0 && kind&FF_DIHEDRAL_TYPE_9 == 0 {
		panic("unsupported dihedral type")
	}

	return &TopDihedral{
		atom1: atom1,
		atom2: atom2,
		atom3: atom3,
		atom4: atom4,
		kind:  kind,
	}
}

//
func (d *TopDihedral) TopAtom1() *TopAtom {
	return d.atom1
}

func (d *TopDihedral) TopAtom2() *TopAtom {
	return d.atom2
}

func (d *TopDihedral) TopAtom3() *TopAtom {
	return d.atom3
}

func (d *TopDihedral) TopAtom4() *TopAtom {
	return d.atom4
}

//
func (d *TopDihedral) SetCustomDihedralType(dt *DihedralType) {
	d.setting |= t_dih_sett_CUSTOM_DIHEDRAL_TYPE_SET
	d.customDihedralType = dt
}

func (d *TopDihedral) HasCustomDihedralTypeSet() bool {
	return d.setting&t_dih_sett_CUSTOM_DIHEDRAL_TYPE_SET != 0
}

func (d *TopDihedral) CustomDihedralType() *DihedralType {
	return d.customDihedralType
}

//
func (d *TopDihedral) Kind() ffTypes {
	return d.kind
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
