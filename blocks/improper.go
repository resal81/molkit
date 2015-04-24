package blocks

import (
	"github.com/resal81/molkit/utils"
)

var (
	improperHash = utils.NewComponentHash()
)

/**********************************************************
* ImproperType
**********************************************************/

type ITSetting int64

const (
	IT_TYPE_CHM_1 ITSetting = 1 << iota
	IT_TYPE_GMX_1
	IT_HAS_PSI_SET
	IT_HAS_PSI_CONSTANT_SET
)

type ImproperType struct {
	aType1   string
	aType2   string
	aType3   string
	aType4   string
	psi      float64
	psiConst float64
	setting  ITSetting
}

/* new dihedraltype */

func NewImproprtType(at1, at2, at3, at4 string, t ITSetting) *ImproperType {
	return &ImproperType{
		aType1:  at1,
		aType2:  at2,
		aType3:  at3,
		aType4:  at4,
		setting: t,
	}
}

/* atom types */

func (dt *ImproperType) AType1() string {
	return dt.aType1
}

func (dt *ImproperType) AType2() string {
	return dt.aType2
}

func (dt *ImproperType) AType3() string {
	return dt.aType3
}

func (dt *ImproperType) AType4() string {
	return dt.aType4
}

/* psi */

func (dt *ImproperType) SetPsi(v float64) {
	dt.setting |= IT_HAS_PSI_SET
	dt.psi = v
}

func (dt *ImproperType) HasPsiSet() bool {
	return dt.setting&IT_HAS_PSI_SET != 0
}

func (dt *ImproperType) Psi() float64 {
	return dt.psi
}

/* psi const */

func (dt *ImproperType) SetPsiConstant(v float64) {
	dt.setting |= IT_HAS_PSI_CONSTANT_SET
	dt.psiConst = v
}

func (dt *ImproperType) HasPsiConstantSet() bool {
	return dt.setting&IT_HAS_PSI_CONSTANT_SET != 0
}

func (dt *ImproperType) PsiConstant() float64 {
	return dt.psiConst
}

/* setting */

func (dt *ImproperType) Setting() ITSetting {
	return dt.setting
}

/**********************************************************
* Improper
**********************************************************/

type Improper struct {
	id    int64
	atom1 *Atom
	atom2 *Atom
	atom3 *Atom
	atom4 *Atom
	tipe  *ImproperType
}

/* new improper */

func NewImproper(atom1, atom2, atom3, atom4 *Atom) *Improper {
	imp := &Improper{
		atom1: atom1,
		atom2: atom2,
		atom3: atom3,
		atom4: atom4,
	}

	imp.id = improperHash.Add(imp)
	return imp
}

/* id */

func (d *Improper) Id() int64 {
	return d.id
}

/* atoms */

func (d *Improper) Atom1() *Atom {
	return d.atom1
}

func (d *Improper) Atom2() *Atom {
	return d.atom2
}

func (d *Improper) Atom3() *Atom {
	return d.atom3
}

func (d *Improper) Atom4() *Atom {
	return d.atom4
}

/* type */

func (d *Improper) SetType(dt *ImproperType) {
	d.tipe = dt
}

func (d *Improper) Type() *ImproperType {
	return d.tipe
}
