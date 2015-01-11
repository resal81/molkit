package ff

// --------------------------------------------------------

type NON_BONDED_TYPE int8
type PAIR_TYPE int8
type BOND_TYPE int64
type ANGLE_TYPE int64
type DIHEDRAL_TYPE int64
type CONSTRAINT_TYPE int8

const (
	NB_TYPE1 NON_BONDED_TYPE = 1 << iota
	NB_TYPE2
)

const (
	P_TYPE1 PAIR_TYPE = 1 << iota
	P_TYPE2
)

const (
	B_TYPE1 BOND_TYPE = 1 << iota
	B_TYPE2
	B_TYPE3
	B_TYPE4
	B_TYPE5
	B_TYPE6
	B_TYPE7
	B_TYPE8
	B_TYPE9
	B_TYPE10
)

const (
	A_TYPE1 ANGLE_TYPE = 1 << iota // GMX 1
	A_TYPE2
	A_TYPE3
	A_TYPE4
	A_TYPE5
	A_TYPE6
	A_TYPE7
	A_TYPE8
	A_TYPE9
	A_TYPE10
)

const (
	D_TYPE1 DIHEDRAL_TYPE = 1 << iota // GMX 1
	D_TYPE2
	D_TYPE3
	D_TYPE4
	D_TYPE5
	D_TYPE6
	D_TYPE7
	D_TYPE8
	D_TYPE9
	D_TYPE10
	D_TYPE11
)

const (
	CNT_TYPE1 CONSTRAINT_TYPE = 1 << iota
	CNT_TYPE2
)

// --------------------------------------------------------

type atSetting int32

const (
	AT_GMX_ATOMTYPE atSetting = 1 << iota
	AT_CHM_ATOMTYPE
	AT_PROTONS_SET
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

func NewAtomType(atype string, src atSetting) *AtomType {
	return &AtomType{
		atype:   atype,
		setting: src,
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

type ParamsNonBondedType struct {
	atype1 string
	atype2 string

	nbtype NON_BONDED_TYPE

	v1 float32
	v2 float32
	v3 float32
}

// --------------------------------------------------------

type ParamsPairType struct {
	atype1  string
	atype2  string
	ptype   PAIR_TYPE
	sigma   float32
	epsilon float32
}

func NewGMXParamsPairType(atype1, atype2 string, fn int8, sigma, epsilon float32) *ParamsPairType {
	pt := ParamsPairType{
		atype1:  atype1,
		atype2:  atype2,
		sigma:   sigma,
		epsilon: epsilon,
	}

	return &pt
}

// --------------------------------------------------------

type ParamsBondType struct {
	atype1 string
	atype2 string

	btype BOND_TYPE
	k_r   float32
	r0    float32
}

func NewGMXParamsBondType(atype1, atype2 string, fn int8, k_r, r0 float32) *ParamsBondType {
	return nil
}

// --------------------------------------------------------

type ParamsAngleType struct {
	atype1 string
	atype2 string
	atype3 string

	k_theta float32
	theta   float32
	r13     float32
	k_ub    float32
}

func NewGMXParamsAngleType(atype1, atype2, atype3 string, fn int8, k_theta, theta, r13, k_ub float32) *ParamsAngleType {
	return nil
}

// --------------------------------------------------------

type ParamsDihedralType struct {
	atype1 string
	atype2 string
	atype3 string
	atype4 string

	k_phi float32
	phi   float32
	mult  int8
}

func NewGMXParamsDihedralType(atype1, atype2, atype3, atype4 string, fn int8, k_phi, phi float32, mult int8) *ParamsDihedralType {
	return nil
}

// --------------------------------------------------------

type ParamsConstraintType struct {
	atype1 string
	atype2 string

	cnttype CONSTRAINT_TYPE
	b0      float32
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
