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

type ParamsAtomType struct {
	atype   string
	protons int8
	mass    float32
	sigma   float32
	epsilon float32
	charge  float32
}

func NewGMXParamsAtomType(atype string, protons int8, mass, charge, sigma, epsilon float32) *ParamsAtomType {
	at := ParamsAtomType{
		atype:   atype,
		protons: protons,
		mass:    mass,
		charge:  charge,
		sigma:   sigma,
		epsilon: epsilon,
	}

	return &at
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

//
