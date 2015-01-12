package ff

/**********************************************************
* AtomType
**********************************************************/

type atSetting int32

const (
	at_sett_PROTONS_SET atSetting = 1 << iota
	at_sett_MASS_SET
	at_sett_CHARGE_SET
	at_sett_SIGMA_SET
	at_sett_SIGMA14_SET
	at_sett_EPSILON_SET
	at_sett_EPSILON14_SET
)

type AtomType struct {
	atype     string
	protons   int8
	mass      float64
	sigma     float64
	epsilon   float64
	sigma14   float64
	epsilon14 float64
	charge    float64

	setting atSetting
}

// Constructor
func NewAtomType(atype string) *AtomType {
	return &AtomType{
		atype: atype,
	}
}

//
func (a *AtomType) SetProtons(u int8) {
	a.setting |= at_sett_PROTONS_SET
	a.protons = u
}

func (a *AtomType) HasProtonsSet() bool {
	return a.setting&at_sett_PROTONS_SET != 0
}

func (a *AtomType) Protons() int8 {
	return a.protons
}

//
func (a *AtomType) SetMass(v float64) {
	a.setting |= at_sett_MASS_SET
	a.mass = v
}

func (a *AtomType) HasMassSet() bool {
	return a.setting&at_sett_MASS_SET != 0
}

func (a *AtomType) Mass() float64 {
	return a.mass
}

//
func (a *AtomType) SetCharge(v float64) {
	a.setting |= at_sett_CHARGE_SET
	a.charge = v
}

func (a *AtomType) HasChargeSet() bool {
	return a.setting&at_sett_CHARGE_SET != 0
}

func (a *AtomType) Charge() float64 {
	return a.charge
}

//
func (a *AtomType) SetSigma(v float64) {
	a.setting |= at_sett_SIGMA_SET
	a.sigma = v
}

func (a *AtomType) HasSigmaSet() bool {
	return a.setting&at_sett_SIGMA_SET != 0
}

func (a *AtomType) Sigma() float64 {
	return a.sigma
}

//
func (a *AtomType) SetEpsilon(v float64) {
	a.setting |= at_sett_EPSILON_SET
	a.epsilon = v
}

func (a *AtomType) HasEpsilonSet() bool {
	return a.setting&at_sett_EPSILON_SET != 0
}

func (a *AtomType) Epsilon() float64 {
	return a.epsilon
}

//
func (a *AtomType) SetSigma14(v float64) {
	a.setting |= at_sett_SIGMA14_SET
	a.sigma14 = v
}

func (a *AtomType) HasSigma14Set() bool {
	return a.setting&at_sett_SIGMA14_SET != 0
}

func (a *AtomType) Sigma14() float64 {
	return a.sigma14
}

//
func (a *AtomType) SetEpsilon14(v float64) {
	a.setting |= at_sett_EPSILON14_SET
	a.epsilon14 = v
}

func (a *AtomType) HasEpsilon14Set() bool {
	return a.setting&at_sett_EPSILON14_SET != 0
}

func (a *AtomType) Epsilon14() float64 {
	return a.epsilon14
}

/**********************************************************
* NonBondedType
**********************************************************/

type nonbondedtypeSetting int32

const (
	nbt_sett_SIGMA_SET nonbondedtypeSetting = 1 << iota
	nbt_sett_EPSILON_SET
)

type NonBondedType struct {
	atype1 string
	atype2 string

	nbtype nonbondedtypeSetting

	sigma   float64
	epsilon float64

	setting nonbondedtypeSetting
	kind    ffTypes
}

func NewNonBondedType(atype1, atype2 string, nbt ffTypes) *NonBondedType {

	if nbt&FF_NON_BONDED_TYPE_1 == 0 {
		panic("unsupported nonbonded-param")
	}
	return &NonBondedType{
		atype1: atype1,
		atype2: atype2,
		kind:   nbt,
	}
}

//
func (n *NonBondedType) AType1() string {
	return n.atype1
}

func (n *NonBondedType) AType2() string {
	return n.atype2
}

//
func (n *NonBondedType) SetSigma(v float64) {
	n.setting |= nbt_sett_SIGMA_SET
	n.sigma = v
}

func (n *NonBondedType) HasSigmaSet() bool {
	return n.setting&nbt_sett_SIGMA_SET != 0
}

func (n *NonBondedType) Sigma() float64 {
	return n.sigma
}

//
func (n *NonBondedType) SetEpsilon(v float64) {
	n.setting |= nbt_sett_EPSILON_SET
	n.epsilon = v
}

func (n *NonBondedType) HasEpsilonSet() bool {
	return n.setting&nbt_sett_EPSILON_SET != 0
}

func (n *NonBondedType) Epsilon() float64 {
	return n.epsilon
}

/**********************************************************
* PairType
**********************************************************/

type pairtypeSetting int32

const (
	pt_sett_SIGMA14_SET pairtypeSetting = 1 << iota
	pt_sett_EPSILON14_SET
)

type PairType struct {
	atype1    string
	atype2    string
	sigma14   float64
	epsilon14 float64

	setting pairtypeSetting
	kind    ffTypes
}

func NewPairType(atype1, atype2 string, pt ffTypes) *PairType {
	if pt&FF_PAIR_TYPE_1 == 0 {
		panic("unsupported pairtype")
	}

	return &PairType{
		atype1: atype1,
		atype2: atype2,
		kind:   pt,
	}
}

//
func (p *PairType) AType1() string {
	return p.atype1
}

func (p *PairType) AType2() string {
	return p.atype2
}

//
func (p *PairType) SetSigma14(v float64) {
	p.setting |= pt_sett_SIGMA14_SET
	p.sigma14 = v
}

func (p *PairType) HasSigma14Set() bool {
	return p.setting&pt_sett_SIGMA14_SET != 0
}

func (p *PairType) Sigma14() float64 {
	return p.sigma14
}

//
func (p *PairType) SetEpsilon14(v float64) {
	p.setting |= pt_sett_EPSILON14_SET
	p.epsilon14 = v
}

func (p *PairType) HasEpsilon14Set() bool {
	return p.setting&pt_sett_EPSILON14_SET != 0
}

func (p *PairType) Epsilon14() float64 {
	return p.epsilon14
}

/**********************************************************
* BondType
**********************************************************/

type bondtypeSetting int32

const (
	bt_sett_HARMONIC_CONSTANT_SET bondtypeSetting = 1 << iota
	bt_sett_HARMONIC_DISTANCE_SET
)

type BondType struct {
	atype1 string
	atype2 string

	kr float64
	r0 float64

	setting bondtypeSetting
	kind    ffTypes
}

func NewBondType(atype1, atype2 string, bt ffTypes) *BondType {
	if bt&FF_BOND_TYPE_1 == 0 {
		panic("unsupported bondtype")
	}

	return &BondType{
		atype1: atype1,
		atype2: atype2,
		kind:   bt,
	}
}

//
func (b *BondType) AType1() string {
	return b.atype1
}

func (b *BondType) AType2() string {
	return b.atype2
}

//
func (b *BondType) SetHarmonicConstant(v float64) {
	b.setting |= bt_sett_HARMONIC_CONSTANT_SET
	b.kr = v
}

func (b *BondType) HasHarmonicConstantSet() bool {
	return b.setting&bt_sett_HARMONIC_CONSTANT_SET != 0
}

func (b *BondType) HarmonicConstant() float64 {
	return b.kr
}

//
func (b *BondType) SetHarmonicDistance(v float64) {
	b.setting |= bt_sett_HARMONIC_DISTANCE_SET
	b.r0 = v
}

func (b *BondType) HasHarmonicDistanceSet() bool {
	return b.setting&bt_sett_HARMONIC_DISTANCE_SET != 0
}

func (b *BondType) HarmonicDistance() float64 {
	return b.r0
}

/**********************************************************
* AngleType
**********************************************************/

type angletypeSetting int32

const (
	ang_sett_THETA_CONSTANT_SET angletypeSetting = 1 << iota
	ang_sett_THETA_SET
	ang_sett_UB_CONSTANT_SET
	ang_sett_R13_SET
)

type AngleType struct {
	atype1 string
	atype2 string
	atype3 string

	k_theta float64
	theta   float64
	r13     float64
	k_ub    float64

	setting angletypeSetting
	kind    ffTypes
}

func NewAngleType(atype1, atype2, atype3 string, angtype ffTypes) *AngleType {

	if angtype&FF_ANGLE_TYPE_1 == 0 && angtype&FF_ANGLE_TYPE_5 == 0 {
		panic("unsupported angle type")
	}

	return &AngleType{
		atype1: atype1,
		atype2: atype2,
		atype3: atype3,
		kind:   angtype,
	}
}

//
func (a *AngleType) AType1() string {
	return a.atype1
}

func (a *AngleType) AType2() string {
	return a.atype2
}

func (a *AngleType) AType3() string {
	return a.atype3
}

//
func (a *AngleType) SetThetaConstant(v float64) {
	a.setting |= ang_sett_THETA_CONSTANT_SET
	a.k_theta = v
}

func (a *AngleType) HasThetaConstantSet() bool {
	return a.setting&ang_sett_THETA_CONSTANT_SET != 0
}

func (a *AngleType) ThetaConstant() float64 {
	return a.k_theta
}

//
func (a *AngleType) SetTheta(v float64) {
	a.setting |= ang_sett_THETA_SET
	a.theta = v
}

func (a *AngleType) HasThetaSet() bool {
	return a.setting&ang_sett_THETA_SET != 0
}

func (a *AngleType) Theta() float64 {
	return a.theta
}

//
func (a *AngleType) SetUBConstant(v float64) {
	a.setting |= ang_sett_UB_CONSTANT_SET
	a.k_ub = v
}

func (a *AngleType) HasUBConstantSet() bool {
	return a.setting&ang_sett_UB_CONSTANT_SET != 0
}

func (a *AngleType) UBConstant() float64 {
	return a.k_ub
}

//
func (a *AngleType) SetR13(v float64) {
	a.setting |= ang_sett_R13_SET
	a.r13 = v
}

func (a *AngleType) HasR13Set() bool {
	return a.setting&ang_sett_R13_SET != 0
}

func (a *AngleType) R13() float64 {
	return a.r13
}

/**********************************************************
* DihedralType
**********************************************************/

type dihedraltypeSetting int32

const (
	dih_sett_PHI_CONSTANT_SET dihedraltypeSetting = 1 << iota
	dih_sett_PHI_SET
	dih_sett_MULT_SET
	dih_sett_PSI_CONSTANT_SET
	dih_sett_PSI_SET
)

type DihedralType struct {
	atype1 string
	atype2 string
	atype3 string
	atype4 string

	k_phi float64
	phi   float64
	mult  int8

	// for improper
	k_psi float64
	psi   float64

	setting dihedraltypeSetting
	kind    ffTypes
}

func NewDihedralType(atype1, atype2, atype3, atype4 string, dht ffTypes) *DihedralType {
	if dht&FF_DIHEDRAL_TYPE_1 == 0 && dht&FF_DIHEDRAL_TYPE_2 == 0 && dht&FF_DIHEDRAL_TYPE_9 == 0 {
		panic("unsupported dihedraltype")
	}

	return &DihedralType{
		atype1: atype1,
		atype2: atype2,
		atype3: atype3,
		atype4: atype4,
		kind:   dht,
	}
}

//
func (d *DihedralType) Kind() ffTypes {
	return d.kind
}

//
func (d *DihedralType) AType1() string {
	return d.atype1
}

func (d *DihedralType) AType2() string {
	return d.atype2
}

func (d *DihedralType) AType3() string {
	return d.atype3
}

func (d *DihedralType) AType4() string {
	return d.atype4
}

//
func (d *DihedralType) SetPhiConstant(v float64) {
	d.setting |= dih_sett_PHI_CONSTANT_SET
	d.k_phi = v
}

func (d *DihedralType) HasPhiConstantSet() bool {
	return d.setting&dih_sett_PHI_CONSTANT_SET != 0
}

func (d *DihedralType) PhiConstant() float64 {
	return d.k_phi
}

//
func (d *DihedralType) SetPhi(v float64) {
	d.setting |= dih_sett_PHI_SET
	d.phi = v
}

func (d *DihedralType) HasPhiSet() bool {
	return d.setting&dih_sett_PHI_SET != 0
}

func (d *DihedralType) Phi() float64 {
	return d.phi
}

//
func (d *DihedralType) SetPsiConstant(v float64) {
	d.setting |= dih_sett_PSI_CONSTANT_SET
	d.k_psi = v
}

func (d *DihedralType) HasPsiConstantSet() bool {
	return d.setting&dih_sett_PSI_CONSTANT_SET != 0
}

func (d *DihedralType) PsiConstant() float64 {
	return d.k_psi
}

//
func (d *DihedralType) SetPsi(v float64) {
	d.setting |= dih_sett_PSI_SET
	d.psi = v
}

func (d *DihedralType) HasPsiSet() bool {
	return d.setting&dih_sett_PSI_SET != 0
}

func (d *DihedralType) Psi() float64 {
	return d.psi
}

//
func (d *DihedralType) SetMult(v int8) {
	d.setting |= dih_sett_MULT_SET
	d.mult = v
}

func (d *DihedralType) HasMultiSet() bool {
	return d.setting&dih_sett_MULT_SET != 0
}

func (d *DihedralType) Mult() int8 {
	return d.mult
}

func (d *DihedralType) Setting() dihedraltypeSetting {
	return d.setting
}

/**********************************************************
* ConstraintType
**********************************************************/

type ConstraintType struct {
	atype1 string
	atype2 string

	b0 float64
}

/**********************************************************
* GMXDefaults
**********************************************************/

type GMXDefaults struct {
	nbfunc   int8
	combrule int8
	genpairs bool
	fudgeLJ  float32
	fudgeQQ  float32
}

func NewGMXDefaults(nbfunc, combrule int8, genpairs bool, fudgeLJ, fudgeQQ float32) *GMXDefaults {
	gd := GMXDefaults{
		nbfunc:   nbfunc,
		combrule: combrule,
		genpairs: genpairs,
		fudgeLJ:  fudgeLJ,
		fudgeQQ:  fudgeQQ,
	}

	return &gd
}

//
