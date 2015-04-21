package blocks

import (
	"github.com/resal81/molkit/utils"
)

var (
	bondHash     = utils.NewComponentHash()
	angleHash    = utils.NewComponentHash()
	dihedralHash = utils.NewComponentHash()
	improperHash = utils.NewComponentHash()
)

// --------------------------------------------------------

type PTSetting int64

const (
	PT_TYPE_CHM PTSetting = 1 << iota
	PT_HAS_LJ_DIST_SET
	PT_HAS_LJ_ENERGY_SET
	PT_HAS_LJ_DIST14_SET
	PT_HAS_LJ_ENERGY14_SET
)

type PairType struct {
	AType1     string
	Atype2     string
	LJDist     float64
	LJEnergy   float64
	LJDist14   float64
	LJEnergy14 float64
	Setting    PTSetting
}

type Pair struct {
	Atom1 *Atom
	Atom2 *Atom
	Type  *PairType
}

// --------------------------------------------------------

type BTSetting int64

const (
	BT_ORDER_SINGLE BTSetting = 1 << iota
	BT_ORDER_DOUBLE
	BT_ORDER_TRIPLE
	BT_TYPE_GMX_1
	BT_TYPE_CHM_1
	BT_HAS_HARM_CONST_SET
	BT_HAS_HARM_DIST_SET
)

type BondType struct {
	AType1    string
	AType2    string
	HarmConst float64
	HarmDist  float64
	Setting   BTSetting
}

type Bond struct {
	id int64

	Atom1 *Atom
	Atom2 *Atom
	Type  *BondType
}

func NewBond(atom1, atom2 *Atom) *Bond {
	bnd := &Bond{
		Atom1: atom1,
		Atom2: atom2,
	}
	bnd.id = bondHash.Add(bnd)
	return bnd
}

func (b *Bond) Id() int64 {
	return b.id
}

// --------------------------------------------------------
type NTSetting int64

const (
	NT_TYPE_CHM   NTSetting = 1 << iota // Harmonic
	NT_TYPE_GMX_1                       // Harmonic
	NT_TYPE_GMX_5                       // UB
	NT_HAS_THETA_CONST_SET
	NT_HAS_THETA_ANGLE_SET
	NT_HAS_UB_CONST_SET
	NT_HAS_R13_SET
)

type AngleType struct {
	AType1        string
	AType2        string
	AType3        string
	ThetaConstant float64
	Theta         float64
	R13           float64
	UBConst       float64
	Setting       NTSetting
}

type Angle struct {
	id    int64
	Atom1 *Atom
	Atom2 *Atom
	Atom3 *Atom
	Type  *AngleType
}

func NewAngle(atom1, atom2, atom3 *Atom) *Angle {
	ang := &Angle{
		Atom1: atom1,
		Atom2: atom2,
		Atom3: atom3,
	}
	ang.id = angleHash.Add(ang)
	return ang
}

func (a *Angle) Id() int64 {
	return a.id
}

// --------------------------------------------------------

type DTSetting int64

const (
	DT_TYPE_CHM   DTSetting = 1 << iota // proper
	DT_TYPE_GMX_1                       // proper
	DT_TYPE_GMX_9                       // prper muliple
	DT_HAS_PHI_ANGLE_SET
	DT_HAS_PHI_CONST_SET
	DT_HAS_MULT_SET
)

type DihedralType struct {
	AType1   string
	AType2   string
	AType3   string
	AType4   string
	PhiAngle float64
	PhiConst float64
	Mult     float64
	Setting  DTSetting
}

type Dihedral struct {
	id    int64
	Atom1 *Atom
	Atom2 *Atom
	Atom3 *Atom
	Atom4 *Atom
	Type  *DihedralType
}

func NewDihedral(atom1, atom2, atom3, atom4 *Atom) *Dihedral {
	dih := &Dihedral{
		Atom1: atom1,
		Atom2: atom2,
		Atom3: atom3,
		Atom4: atom4,
	}

	dih.id = dihedralHash.Add(dih)
	return dih
}

func (d *Dihedral) Id() int64 {
	return d.id
}

// --------------------------------------------------------

type ITSetting int64

const (
	IT_TYPE_CHM ITSetting = 1 << iota
	IT_TYPE_GMX_1
	IT_HAS_PSI_ANGLE_SET
	IT_HAS_PSI_CONST_SET
)

type ImproperType struct {
	AType1   string
	AType2   string
	AType3   string
	AType4   string
	PsiAngle float64
	PsiConst float64
	Setting  ITSetting
}

type Improper struct {
	id    int64
	Atom1 *Atom
	Atom2 *Atom
	Atom3 *Atom
	Atom4 *Atom
	Type  ImproperType
}

func NewImproper(atom1, atom2, atom3, atom4 *Atom) *Improper {
	imp := &Improper{
		Atom1: atom1,
		Atom2: atom2,
		Atom3: atom3,
		Atom4: atom4,
	}

	imp.id = improperHash.Add(imp)
	return imp
}

func (d *Improper) Id() int64 {
	return d.id
}

// --------------------------------------------------------
