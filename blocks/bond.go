package blocks

import (
	"github.com/resal81/molkit/utils"
)

var (
	bondHash = utils.NewComponentHash()
)

/**********************************************************
* BondType
**********************************************************/

type BTSetting int64

const (
	BT_ORDER_SINGLE BTSetting = 1 << iota
	BT_ORDER_DOUBLE
	BT_ORDER_TRIPLE
	BT_TYPE_GMX_1
	BT_TYPE_CHM_1
	BT_HAS_HARMONIC_CONSTANT_SET
	BT_HAS_HARMONIC_DISTANCE_SET
)

type BondType struct {
	aType1    string
	aType2    string
	harmConst float64
	harmDist  float64
	setting   BTSetting
}

/* new bondtype */

func NewBondType(at1, at2 string) *BondType {
	return &BondType{
		aType1: at1,
		aType2: at2,
	}
}

/* types */

func (bt *BondType) AType1() string {
	return bt.aType1
}

func (bt *BondType) AType2() string {
	return bt.aType2
}

/* harmConst */

func (bt *BondType) SetHarmonicConstant(v float64) {
	bt.setting |= BT_HAS_HARMONIC_CONSTANT_SET
	bt.harmConst = v
}

func (bt *BondType) HasHarmonicConstantSet() bool {
	return bt.setting&BT_HAS_HARMONIC_CONSTANT_SET != 0
}

func (bt *BondType) HarmonicConstant() float64 {
	return bt.harmConst
}

/* harmDist */

func (bt *BondType) SetHarmonicDistance(v float64) {
	bt.setting |= BT_HAS_HARMONIC_DISTANCE_SET
	bt.harmDist = v
}

func (bt *BondType) HasHarmonicDistanceSet() bool {
	return bt.setting&BT_HAS_HARMONIC_DISTANCE_SET != 0
}

func (bt *BondType) HarmonicDistance() float64 {
	return bt.harmDist
}

/* setting */
func (bt *BondType) Setting() BTSetting {
	return bt.setting
}

/* order */
func (bt *BondType) SetOrderSingle() {
	bt.setting |= BT_ORDER_SINGLE
}

func (bt *BondType) IsSingle() bool {
	return bt.setting&BT_ORDER_SINGLE != 0
}

/**********************************************************
* Bond
**********************************************************/

type Bond struct {
	id int64

	atom1 *Atom
	atom2 *Atom
	tipe  *BondType
}

func NewBond(atom1, atom2 *Atom) *Bond {
	bnd := &Bond{
		atom1: atom1,
		atom2: atom2,
	}
	atom1.AddBond(bnd)
	atom2.AddBond(bnd)
	bnd.id = bondHash.Add(bnd)
	return bnd
}

func (b *Bond) Id() int64 {
	return b.id
}

/* atoms */
func (b *Bond) Atom1() *Atom {
	return b.atom1
}

func (b *Bond) Atom2() *Atom {
	return b.atom2
}

/* type */

func (b *Bond) SetType(bt *BondType) {
	b.tipe = bt
}

func (b *Bond) Type() *BondType {
	return b.tipe
}
