package ff

// --------------------------------------------------------

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
	mass      float32
	sigma     float32
	epsilon   float32
	sigma14   float32
	epsilon14 float32
	charge    float32

	setting atSetting
}

func NewAtomType(atype string) *AtomType {
	return &AtomType{
		atype: atype,
	}
}

func (a *AtomType) SetProtons(u int8) {
	a.setting |= AT_PROTONS_SET
	a.protons = u
}

func (a *AtomType) Protons() int8 {
	return a.protons
}

func (a *AtomType) SetMass(v float32) {
	a.setting |= AT_MASS_SET
	a.mass = v
}

func (a *AtomType) Mass() float32 {
	return a.mass
}

func (a *AtomType) SetCharge(v float32) {
	a.setting |= AT_CHARGE_SET
	a.charge = v
}

func (a *AtomType) Charge() float32 {
	return a.charge
}

func (a *AtomType) SetSigma(v float32) {
	a.setting |= AT_SIGMA_SET
	a.sigma = v
}

func (a *AtomType) Sigma() float32 {
	return a.sigma
}

func (a *AtomType) SetEpsilon(v float32) {
	a.setting |= AT_EPSILON_SET
	a.epsilon = v
}

func (a *AtomType) Epsilon() float32 {
	return a.epsilon
}

func (a *AtomType) SetSigma14(v float32) {
	a.setting |= AT_SIGMA14_SET
	a.sigma14 = v
}

func (a *AtomType) Sigma14() float32 {
	return a.sigma14
}

func (a *AtomType) SetEpsilon14(v float32) {
	a.setting |= AT_EPSILON14_SET
	a.epsilon14 = v
}

func (a *AtomType) Epsilon14() float32 {
	return a.epsilon14
}

// --------------------------------------------------------

type nonbondedtypeSetting int32

const (
	NBT_TYPE_1 nonbondedtypeSetting = 1 << iota
)

type NonBondedType struct {
	atype1 string
	atype2 string

	nbtype nonbondedtypeSetting

	sigma   float32
	epsilon float32
}

func NewNonBondedType(atype1, atype2 string) *NonBondedType {
	return &NonBondedType{
		atype1: atype1,
		atype2: atype2,
	}
}

func (n *NonBondedType) SetSigma(v float32) {
	n.sigma = v
}

func (n *NonBondedType) Sigma() float32 {
	return n.sigma
}

func (n *NonBondedType) SetEpsilon(v float32) {
	n.epsilon = v
}

func (n *NonBondedType) Epsilon() float32 {
	return n.epsilon
}

// --------------------------------------------------------

type pairtypeSetting int32

const (
	PT_TYPE_1 pairtypeSetting = 1 << iota
)

type PairType struct {
	atype1    string
	atype2    string
	sigma14   float32
	epsilon14 float32

	setting pairtypeSetting
}

func (p *PairType) NewPairType(atype1, atype2 string) *PairType {
	return &PairType{
		atype1: atype1,
		atype2: atype2,
	}
}

func (p *PairType) SetSigma14(v float32) {
	p.sigma14 = v
}

func (p *PairType) Sigma14() float32 {
	return p.sigma14
}

func (p *PairType) SetEpsilon14(v float32) {
	p.epsilon14 = v
}

func (p *PairType) Epsilon14() float32 {
	return p.epsilon14
}

// --------------------------------------------------------

type bondtypeSetting int64

const (
	BT_TYPE_1 bondtypeSetting = 1 << iota // harmonic bond
)

type BondType struct {
	atype1 string
	atype2 string

	kr float32
	r0 float32

	setting bondtypeSetting
}

func NewBondType(atype1, atype2 string) *BondType {
	return &BondType{
		atype1: atype1,
		atype2: atype2,
	}
}

func (b *BondType) SetHarmonicConstant(v float32) {
	b.kr = v
}

func (b *BondType) HarmonicConstant() float32 {
	return b.kr
}

func (b *BondType) SetHarmonicDistance(v float32) {
	b.r0 = v
}

func (b *BondType) HarmonicDistance() float32 {
	return b.r0
}

// --------------------------------------------------------

type AngleType struct {
	atype1 string
	atype2 string
	atype3 string

	k_theta float32
	theta   float32
	r13     float32
	k_ub    float32
}

// --------------------------------------------------------

type DihedralType struct {
	atype1 string
	atype2 string
	atype3 string
	atype4 string

	k_phi float32
	phi   float32
	mult  int8
}

// --------------------------------------------------------

type ConstraintType struct {
	atype1 string
	atype2 string

	b0 float32
}

// --------------------------------------------------------

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
