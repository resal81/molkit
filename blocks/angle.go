package blocks

import (
	"github.com/resal81/molkit/utils"
)

var (
	angleHash = utils.NewComponentHash()
)

/**********************************************************
* AngleType
**********************************************************/

type NTSetting int64

const (
	NT_TYPE_CHM_1 NTSetting = 1 << iota // Harmonic
	NT_TYPE_CHM_2                       // UB
	NT_TYPE_GMX_1                       // Harmonic
	NT_TYPE_GMX_5                       // UB
	NT_HAS_THETA_CONSTANT_SET
	NT_HAS_THETA_SET
	NT_HAS_UB_CONSTANT_SET
	NT_HAS_R13_SET
)

type AngleType struct {
	aType1     string
	aType2     string
	aType3     string
	thetaConst float64
	theta      float64
	r13        float64
	ubConst    float64
	setting    NTSetting
}

/* new angletype */

func NewAngleType(at1, at2, at3 string, t NTSetting) *AngleType {
	return &AngleType{
		aType1:  at1,
		aType2:  at2,
		aType3:  at3,
		setting: t,
	}
}

/* atom types */

func (at *AngleType) AType1() string {
	return at.aType1
}

func (at *AngleType) AType2() string {
	return at.aType2
}

func (at *AngleType) AType3() string {
	return at.aType3
}

/* theta const */

func (at *AngleType) SetThetaConstant(th float64) {
	at.setting |= NT_HAS_THETA_CONSTANT_SET
	at.thetaConst = th
}

func (at *AngleType) HasThetaConstantSet() bool {
	return at.setting&NT_HAS_THETA_CONSTANT_SET != 0
}

func (at *AngleType) ThetaConstant() float64 {
	return at.thetaConst
}

/* theta */

func (at *AngleType) SetTheta(th float64) {
	at.setting |= NT_HAS_THETA_SET
	at.theta = th
}

func (at *AngleType) HasThetaSet() bool {
	return at.setting&NT_HAS_THETA_SET != 0
}

func (at *AngleType) Theta() float64 {
	return at.theta
}

/* ub const */

func (at *AngleType) SetUBConstant(ub float64) {
	at.setting |= NT_HAS_UB_CONSTANT_SET
	at.ubConst = ub
}

func (at *AngleType) HasUBConstantSet() bool {
	return at.setting&NT_HAS_UB_CONSTANT_SET != 0
}

func (at *AngleType) UBConstant() float64 {
	return at.ubConst
}

/* r13 */

func (at *AngleType) SetR13(th float64) {
	at.setting |= NT_HAS_R13_SET
	at.r13 = th
}

func (at *AngleType) HasR13Set() bool {
	return at.setting&NT_HAS_R13_SET != 0
}

func (at *AngleType) R13() float64 {
	return at.r13
}

/* setting */

func (at *AngleType) Setting() NTSetting {
	return at.setting
}

/**********************************************************
* Angle
**********************************************************/

type Angle struct {
	id    int64
	atom1 *Atom
	atom2 *Atom
	atom3 *Atom
	tipe  *AngleType
}

/* new angle */

func NewAngle(atom1, atom2, atom3 *Atom) *Angle {
	ang := &Angle{
		atom1: atom1,
		atom2: atom2,
		atom3: atom3,
	}
	ang.id = angleHash.Add(ang)
	return ang
}

/* id */

func (a *Angle) Id() int64 {
	return a.id
}

/* atoms */

func (a *Angle) Atom1() *Atom {
	return a.atom1
}

func (a *Angle) Atom2() *Atom {
	return a.atom2
}

func (a *Angle) Atom3() *Atom {
	return a.atom3
}

/* type */

func (a *Angle) SetType(at *AngleType) {
	a.tipe = at
}

func (a *Angle) Type() *AngleType {
	return a.tipe
}
