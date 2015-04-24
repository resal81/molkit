package blocks

import (
	"github.com/resal81/molkit/utils"
)

var (
	dihedralHash = utils.NewComponentHash()
)

/**********************************************************
* DihedralType
**********************************************************/

type DTSetting int64

const (
	DT_TYPE_CHM_1 DTSetting = 1 << iota // proper
	DT_TYPE_GMX_1                       // proper
	DT_TYPE_GMX_9                       // prper muliple
	DT_HAS_PHI_SET
	DT_HAS_PHI_CONSTANT_SET
	DT_HAS_MULTIPLICITY_SET
)

type DihedralType struct {
	aType1   string
	aType2   string
	aType3   string
	aType4   string
	phi      float64
	phiConst float64
	mult     int
	setting  DTSetting
}

/* new dihedraltype */

func NewDihedralType(at1, at2, at3, at4 string, t DTSetting) *DihedralType {
	return &DihedralType{
		aType1:  at1,
		aType2:  at2,
		aType3:  at3,
		aType4:  at4,
		setting: t,
	}
}

/* atom types */

func (dt *DihedralType) AType1() string {
	return dt.aType1
}

func (dt *DihedralType) AType2() string {
	return dt.aType2
}

func (dt *DihedralType) AType3() string {
	return dt.aType3
}

func (dt *DihedralType) AType4() string {
	return dt.aType4
}

/* phi */

func (dt *DihedralType) SetPhi(v float64) {
	dt.setting |= DT_HAS_PHI_SET
	dt.phi = v
}

func (dt *DihedralType) HasPhiSet() bool {
	return dt.setting&DT_HAS_PHI_SET != 0
}

func (dt *DihedralType) Phi() float64 {
	return dt.phi
}

/* phi const */

func (dt *DihedralType) SetPhiConstant(v float64) {
	dt.setting |= DT_HAS_PHI_CONSTANT_SET
	dt.phiConst = v
}

func (dt *DihedralType) HasPhiConstantSet() bool {
	return dt.setting&DT_HAS_PHI_CONSTANT_SET != 0
}

func (dt *DihedralType) PhiConstant() float64 {
	return dt.phiConst
}

/* mult */

func (dt *DihedralType) SetMultiplicity(v int) {
	dt.setting |= DT_HAS_MULTIPLICITY_SET
	dt.mult = v
}

func (dt *DihedralType) HasMultiplicitySet() bool {
	return dt.setting&DT_HAS_MULTIPLICITY_SET != 0
}

func (dt *DihedralType) Multiplicity() int {
	return dt.mult
}

/* setting */

func (dt *DihedralType) Setting() DTSetting {
	return dt.setting
}

/**********************************************************
* Dihedral
**********************************************************/

type Dihedral struct {
	id    int64
	atom1 *Atom
	atom2 *Atom
	atom3 *Atom
	atom4 *Atom
	tipe  *DihedralType
}

/* new dihedral */

func NewDihedral(atom1, atom2, atom3, atom4 *Atom) *Dihedral {
	dih := &Dihedral{
		atom1: atom1,
		atom2: atom2,
		atom3: atom3,
		atom4: atom4,
	}

	dih.id = dihedralHash.Add(dih)
	return dih
}

/* id */

func (d *Dihedral) Id() int64 {
	return d.id
}

/* atoms */

func (d *Dihedral) Atom1() *Atom {
	return d.atom1
}

func (d *Dihedral) Atom2() *Atom {
	return d.atom2
}

func (d *Dihedral) Atom3() *Atom {
	return d.atom3
}

func (d *Dihedral) Atom4() *Atom {
	return d.atom4
}

/* type */

func (d *Dihedral) SetType(dt *DihedralType) {
	d.tipe = dt
}

func (d *Dihedral) Type() *DihedralType {
	return d.tipe
}
