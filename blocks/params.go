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

// --------------------------------------------------------

type PairType struct {
	atyp1   string
	atyp2   string
	sigma   float32
	epsilon float32
}

// --------------------------------------------------------

type BOND_TYPE int64

const (
	B_TYPE1 BOND_TYPE = 1 << iota
	B_TYPE2
	B_TYPE3
	B_TYPE4
	B_TYPE5
)

type BondType struct {
	atype1 string
	atype2 string

	btype BOND_TYPE
	k_r   float32
	r0    float32
}

// --------------------------------------------------------

type ANGLE_TYPE int64

const (
	A_TYPE1 ANGLE_TYPE = 1 << iota // GMX 1
	A_TYPE2
	A_TYPE3
	A_TYPE4
	A_TYPE5
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

// --------------------------------------------------------

type DIHEDRAL_TYPE int64

const (
	D_TYPE1 DIHEDRAL_TYPE = 1 << iota // GMX 1
	D_TYPE2
	D_TYPE3
	D_TYPE4
	D_TYPE5
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
