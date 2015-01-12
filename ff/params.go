package ff

/**********************************************************
* AtomType
**********************************************************/

type atSetting int32

const (
	AT_PROTONS_SET atSetting = 1 << iota
	AT_MASS_SET
	AT_CHARGE_SET
	AT_SIGMA_SET
	AT_SIGMA14_SET
	AT_EPSILON_SET
	AT_EPSILON14_SET
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
	a.setting |= AT_PROTONS_SET
	a.protons = u
}

func (a *AtomType) HasProtonsSet() bool {
	return a.setting&AT_PROTONS_SET != 0
}

func (a *AtomType) Protons() int8 {
	return a.protons
}

//
func (a *AtomType) SetMass(v float64) {
	a.setting |= AT_MASS_SET
	a.mass = v
}

func (a *AtomType) HasMassSet() bool {
	return a.setting&AT_MASS_SET != 0
}

func (a *AtomType) Mass() float64 {
	return a.mass
}

//
func (a *AtomType) SetCharge(v float64) {
	a.setting |= AT_CHARGE_SET
	a.charge = v
}

func (a *AtomType) HasChargeSet() bool {
	return a.setting&AT_CHARGE_SET != 0
}

func (a *AtomType) Charge() float64 {
	return a.charge
}

//
func (a *AtomType) SetSigma(v float64) {
	a.setting |= AT_SIGMA_SET
	a.sigma = v
}

func (a *AtomType) HasSigmaSet() bool {
	return a.setting&AT_SIGMA_SET != 0
}

func (a *AtomType) Sigma() float64 {
	return a.sigma
}

//
func (a *AtomType) SetEpsilon(v float64) {
	a.setting |= AT_EPSILON_SET
	a.epsilon = v
}

func (a *AtomType) HasEpsilonSet() bool {
	return a.setting&AT_EPSILON_SET != 0
}

func (a *AtomType) Epsilon() float64 {
	return a.epsilon
}

//
func (a *AtomType) SetSigma14(v float64) {
	a.setting |= AT_SIGMA14_SET
	a.sigma14 = v
}

func (a *AtomType) HasSigma14Set() bool {
	return a.setting&AT_SIGMA14_SET != 0
}

func (a *AtomType) Sigma14() float64 {
	return a.sigma14
}

//
func (a *AtomType) SetEpsilon14(v float64) {
	a.setting |= AT_EPSILON14_SET
	a.epsilon14 = v
}

func (a *AtomType) HasEpsilon14Set() bool {
	return a.setting&AT_EPSILON14_SET != 0
}

func (a *AtomType) Epsilon14() float64 {
	return a.epsilon14
}

/**********************************************************
* NonBondedType
**********************************************************/

type nonbondedtypeSetting int32

const (
	NBT_TYPE_1 nonbondedtypeSetting = 1 << iota

	NBT_SIGMA_SET
	NBT_EPSILON_SET
)

type NonBondedType struct {
	atype1 string
	atype2 string

	nbtype nonbondedtypeSetting

	sigma   float64
	epsilon float64

	setting nonbondedtypeSetting
}

func NewNonBondedType(atype1, atype2 string, nbt nonbondedtypeSetting) *NonBondedType {

	if nbt&NBT_TYPE_1 == 0 {
		panic("unsupported nonbonded-param")
	}
	return &NonBondedType{
		atype1:  atype1,
		atype2:  atype2,
		setting: nbt,
	}
}

//
func (n *NonBondedType) SetSigma(v float64) {
	n.setting |= NBT_SIGMA_SET
	n.sigma = v
}

func (n *NonBondedType) HasSigmaSet() bool {
	return n.setting&NBT_SIGMA_SET != 0
}

func (n *NonBondedType) Sigma() float64 {
	return n.sigma
}

//
func (n *NonBondedType) SetEpsilon(v float64) {
	n.setting |= NBT_EPSILON_SET
	n.epsilon = v
}

func (n *NonBondedType) HasEpsilonSet() bool {
	return n.setting&NBT_EPSILON_SET != 0
}

func (n *NonBondedType) Epsilon() float64 {
	return n.epsilon
}

/**********************************************************
* PairType
**********************************************************/

type pairtypeSetting int32

const (
	PT_TYPE_1 pairtypeSetting = 1 << iota

	PT_SIGMA14_SET
	PT_EPSILON14_SET
)

type PairType struct {
	atype1    string
	atype2    string
	sigma14   float64
	epsilon14 float64

	setting pairtypeSetting
}

func NewPairType(atype1, atype2 string, pt pairtypeSetting) *PairType {
	if pt&PT_TYPE_1 == 0 {
		panic("unsupported pairtype")
	}

	return &PairType{
		atype1: atype1,
		atype2: atype2,
	}
}

//
func (p *PairType) SetSigma14(v float64) {
	p.setting |= PT_SIGMA14_SET
	p.sigma14 = v
}

func (p *PairType) HasSigma14Set() bool {
	return p.setting&PT_SIGMA14_SET != 0
}

func (p *PairType) Sigma14() float64 {
	return p.sigma14
}

//
func (p *PairType) SetEpsilon14(v float64) {
	p.setting |= PT_EPSILON14_SET
	p.epsilon14 = v
}

func (p *PairType) HasEpsilon14Set() bool {
	return p.setting&PT_EPSILON14_SET != 0
}

func (p *PairType) Epsilon14() float64 {
	return p.epsilon14
}

/**********************************************************
* BondType
**********************************************************/

type bondtypeSetting int32

const (
	BT_TYPE_1 bondtypeSetting = 1 << iota // harmonic bond

	BT_HARMONIC_CONSTANT_SET
	BT_HARMONIC_DISTANCE_SET
)

type BondType struct {
	atype1 string
	atype2 string

	kr float64
	r0 float64

	setting bondtypeSetting
}

func NewBondType(atype1, atype2 string, bt bondtypeSetting) *BondType {
	if bt&BT_TYPE_1 == 0 {
		panic("unsupported bondtype")
	}

	return &BondType{
		atype1:  atype1,
		atype2:  atype2,
		setting: bt,
	}
}

//
func (b *BondType) SetHarmonicConstant(v float64) {
	b.setting |= BT_HARMONIC_CONSTANT_SET
	b.kr = v
}

func (b *BondType) HasHarmonicConstantSet() bool {
	return b.setting&BT_HARMONIC_CONSTANT_SET != 0
}

func (b *BondType) HarmonicConstant() float64 {
	return b.kr
}

//
func (b *BondType) SetHarmonicDistance(v float64) {
	b.setting |= BT_HARMONIC_DISTANCE_SET
	b.r0 = v
}

func (b *BondType) HasHarmonicDistanceSet() bool {
	return b.setting&BT_HARMONIC_DISTANCE_SET != 0
}

func (b *BondType) HarmonicDistance() float64 {
	return b.r0
}

/**********************************************************
* AngleType
**********************************************************/

type angletypeSetting int32

const (
	ANG_TYPE_1 angletypeSetting = 1 << iota // harmonic
	ANG_TYPE_5                              // UB
	ANG_THETA_CONSTANT_SET
	ANG_THETA_SET
	ANG_UB_CONSTANT_SET
	ANG_R13_SET
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
}

func NewAngleType(atype1, atype2, atype3 string, angtype angletypeSetting) *AngleType {

	if angtype&ANG_TYPE_1 == 0 && angtype&ANG_TYPE_5 == 0 {
		panic("unsupported angle type")
	}

	return &AngleType{
		atype1:  atype1,
		atype2:  atype2,
		atype3:  atype3,
		setting: angtype,
	}
}

//
func (a *AngleType) SetThetaConstant(v float64) {
	a.setting |= ANG_THETA_CONSTANT_SET
	a.k_theta = v
}

func (a *AngleType) HasThetaConstantSet() bool {
	return a.setting&ANG_THETA_CONSTANT_SET != 0
}

func (a *AngleType) ThetaConstant() float64 {
	return a.k_theta
}

//
func (a *AngleType) SetTheta(v float64) {
	a.setting |= ANG_THETA_SET
	a.theta = v
}

func (a *AngleType) HasThetaSet() bool {
	return a.setting&ANG_THETA_SET != 0
}

func (a *AngleType) Theta() float64 {
	return a.theta
}

//
func (a *AngleType) SetUBConstant(v float64) {
	a.setting |= ANG_UB_CONSTANT_SET
	a.k_ub = v
}

func (a *AngleType) HasUBConstantSet() bool {
	return a.setting&ANG_UB_CONSTANT_SET != 0
}

func (a *AngleType) UBConstant() float64 {
	return a.k_ub
}

//
func (a *AngleType) SetR13(v float64) {
	a.setting |= ANG_R13_SET
	a.r13 = v
}

func (a *AngleType) HasR13Set() bool {
	return a.setting&ANG_R13_SET != 0
}

func (a *AngleType) R13() float64 {
	return a.r13
}

/**********************************************************
* DihedralType
**********************************************************/

type dihedraltypeSetting int32

const (
	DHT_TYPE_1 dihedraltypeSetting = 1 << iota // proper dihedral
	DHT_TYPE_2                                 // improper
	DHT_TYPE_9                                 // proper multiple
	DHT_PHI_CONSTANT_SET
	DHT_PHI_SET
	DHT_MULT_SET
	DHT_PSI_CONSTANT_SET
	DHT_PSI_SET
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
}

func NewDihedralType(atype1, atype2, atype3, atype4 string, dht dihedraltypeSetting) *DihedralType {
	if dht&DHT_TYPE_1 == 0 && dht&DHT_TYPE_2 == 0 && dht&DHT_TYPE_9 == 0 {
		panic("unsupported dihedraltype")
	}

	return &DihedralType{
		atype1:  atype1,
		atype2:  atype2,
		atype3:  atype3,
		atype4:  atype4,
		setting: dht,
	}
}

//
func (d *DihedralType) SetPhiConstant(v float64) {
	d.setting |= DHT_PHI_CONSTANT_SET
	d.k_phi = v
}

func (d *DihedralType) HasPhiConstantSet() bool {
	return d.setting&DHT_PHI_CONSTANT_SET != 0
}

func (d *DihedralType) PhiConstant() float64 {
	return d.k_phi
}

//
func (d *DihedralType) SetPhi(v float64) {
	d.setting |= DHT_PHI_SET
	d.phi = v
}

func (d *DihedralType) HasPhiSet() bool {
	return d.setting&DHT_PHI_SET != 0
}

func (d *DihedralType) Phi() float64 {
	return d.phi
}

//
func (d *DihedralType) SetPsiConstant(v float64) {
	d.setting |= DHT_PSI_CONSTANT_SET
	d.k_psi = v
}

func (d *DihedralType) HasPsiConstantSet() bool {
	return d.setting&DHT_PSI_CONSTANT_SET != 0
}

func (d *DihedralType) PsiConstant() float64 {
	return d.k_psi
}

//
func (d *DihedralType) SetPsi(v float64) {
	d.setting |= DHT_PSI_SET
	d.psi = v
}

func (d *DihedralType) HasPsiSet() bool {
	return d.setting&DHT_PSI_SET != 0
}

func (d *DihedralType) Psi() float64 {
	return d.psi
}

//
func (d *DihedralType) SetMult(v int8) {
	d.setting |= DHT_MULT_SET
	d.mult = v
}

func (d *DihedralType) HasMultiSet() bool {
	return d.setting&DHT_MULT_SET != 0
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
