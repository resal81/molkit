package blocks

// --------------------------------------------------------

type AtomType struct {
	atype   string
	protons int8
	mass    float32
	sigma   float32
	epsilon float32
	charge  float32
}

func NewGMXAtomType(atype string, protons int8, mass, charge, sigma, epsilon float32) *AtomType {
	at := AtomType{
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

type NON_BONDED_TYPE int8

const (
	NB_TYPE1 NON_BONDED_TYPE = 1 << iota
	NB_TYPE2
)

type NonBondedType struct {
	atype1 string
	atype2 string

	nbtype NON_BONDED_TYPE

	v1 float32
	v2 float32
	v3 float32
}

// --------------------------------------------------------

type PAIR_TYPE int8

const (
	P_TYPE1 PAIR_TYPE = 1 << iota
	P_TYPE2
)

type PairType struct {
	atype1  string
	atype2  string
	ptype   PAIR_TYPE
	sigma   float32
	epsilon float32
}

func NewGMXPairType(atype1, atype2 string, fn int8, sigma, epsilon float32) *PairType {
	pt := PairType{
		atype1:  atype1,
		atype2:  atype2,
		sigma:   sigma,
		epsilon: epsilon,
	}

	return &pt
}

// --------------------------------------------------------

type BOND_TYPE int64

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

type BondType struct {
	atype1 string
	atype2 string

	btype BOND_TYPE
	k_r   float32
	r0    float32
}

func NewGMXBondType(atype1, atype2 string, fn int8, k_r, r0 float32) *BondType {
	return nil
}

// --------------------------------------------------------

type ANGLE_TYPE int64

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

type AngleType struct {
	atype1 string
	atype2 string
	atype3 string

	k_theta float32
	theta   float32
	r13     float32
	k_ub    float32
}

func NewGMXAngleType(atype1, atype2, atype3 string, fn int8, k_theta, theta, r13, k_ub float32) *AngleType {
	return nil
}

// --------------------------------------------------------

type DIHEDRAL_TYPE int64

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

type DihedralType struct {
	atype1 string
	atype2 string
	atype3 string
	atype4 string

	k_phi float32
	phi   float32
	mult  int8
}

func NewGMXDihedralType(atype1, atype2, atype3, atype4 string, fn int8, k_phi, phi float32, mult int8) *DihedralType {
	return nil
}

// --------------------------------------------------------

type CONSTRAINT_TYPE int8

const (
	CNT_TYPE1 CONSTRAINT_TYPE = 1 << iota
	CNT_TYPE2
)

type ConstraintType struct {
	atype1 string
	atype2 string

	cnttype CONSTRAINT_TYPE
	b0      float32
}

// --------------------------------------------------------

//
