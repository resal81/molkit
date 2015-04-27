package blocks

import (
	"fmt"
	"math"
)

/**********************************************************
* PairType
**********************************************************/

type PTSetting int64

const (
	PT_NULL       PTSetting = 1 << iota
	PT_TYPE_CHM_1           // CHARMM NBFIX
	PT_TYPE_GMX_1           // equivalent to CHARMM NBFIX
	PT_HAS_LJ_DISTANCE_SET
	PT_HAS_LJ_ENERGY_SET
	PT_HAS_LJ_DISTANCE_14_SET
	PT_HAS_LJ_ENERGY_14_SET
)

// PairType is used to describe custom parameters for non-bonding interactions
type PairType struct {
	aType1     string
	aType2     string
	ljDist     float64
	ljEnergy   float64
	ljDist14   float64
	ljEnergy14 float64
	setting    PTSetting
}

/* new pairtype */

func NewPairType(at1, at2 string, t PTSetting) *PairType {
	return &PairType{
		aType1:  at1,
		aType2:  at2,
		setting: t,
	}
}

/* types */

func (pt *PairType) AType1() string {
	return pt.aType1
}

func (pt *PairType) AType2() string {
	return pt.aType2
}

/* ljDist */

func (pt *PairType) SetLJDistance(v float64) {
	pt.setting |= PT_HAS_LJ_DISTANCE_SET
	pt.ljDist = v
}
func (pt *PairType) HasLJDistanceSet() bool {
	return pt.setting&PT_HAS_LJ_DISTANCE_SET != 0
}
func (pt *PairType) LJDistance() float64 {
	return pt.ljDist
}

/* ljDist14 */

func (pt *PairType) SetLJDistance14(v float64) {
	pt.setting |= PT_HAS_LJ_DISTANCE_14_SET
	pt.ljDist14 = v
}
func (pt *PairType) HasLJDistance14Set() bool {
	return pt.setting&PT_HAS_LJ_DISTANCE_14_SET != 0
}
func (pt *PairType) LJDistance14() float64 {
	return pt.ljDist14
}

/* ljEnergy */

func (pt *PairType) SetLJEnergy(v float64) {
	pt.setting |= PT_HAS_LJ_ENERGY_SET
	pt.ljEnergy = v
}
func (pt *PairType) HasLJEnergySet() bool {
	return pt.setting&PT_HAS_LJ_ENERGY_SET != 0
}
func (pt *PairType) LJEnergy() float64 {
	return pt.ljEnergy
}

/* ljEnergy14 */

func (pt *PairType) SetLJEnergy14(v float64) {
	pt.setting |= PT_HAS_LJ_ENERGY_14_SET
	pt.ljEnergy14 = v
}
func (pt *PairType) HasLJEnergy14Set() bool {
	return pt.setting&PT_HAS_LJ_ENERGY_14_SET != 0
}
func (pt *PairType) LJEnergy14() float64 {
	return pt.ljEnergy14
}

/* setting */

func (pt *PairType) Setting() PTSetting {
	return pt.setting
}

/* convert */

func (pt *PairType) ConvertTo(to PTSetting) (*PairType, error) {

	// check to is valid
	if to&PT_TYPE_CHM_1 == 0 && to&PT_TYPE_GMX_1 == 0 {
		return nil, fmt.Errorf("'to' parameter is not known")
	}

	// check if both types are the same
	if to&pt.setting != 0 {
		return pt, nil
	}

	// if we are CHM type
	if pt.setting&PT_TYPE_CHM_1 != 0 {

		switch {
		case to&PT_TYPE_GMX_1 != 0:

			npt := NewPairType(pt.AType1(), pt.AType2(), PT_TYPE_GMX_1)

			if pt.HasLJEnergySet() {
				npt.SetLJEnergy(math.Abs(pt.LJEnergy()) * 4.184)
			}

			if pt.HasLJDistanceSet() {
				npt.SetLJDistance(pt.LJDistance() * 0.1 / math.Pow(2.0, 1.0/6.0))
			}

			if pt.HasLJEnergy14Set() {
				npt.SetLJEnergy14(math.Abs(pt.LJEnergy14()) * 4.184)
			}

			if pt.HasLJDistance14Set() {
				// TODO double check this coversion
				npt.SetLJDistance14(pt.LJDistance14() * 2 * 0.1 / math.Pow(2.0, 1.0/6.0))
			}
			return npt, nil
		}

	}

	// TODO conversion for other types

	return nil, nil
}

/**********************************************************
* Pair
**********************************************************/

type Pair struct {
	atom1 *Atom
	atom2 *Atom
	tipe  *PairType
}

/* new pair */

func NewPair(a1, a2 *Atom) *Pair {
	return &Pair{
		atom1: a1,
		atom2: a2,
	}
}

/* atoms */

func (p *Pair) Atom1() *Atom {
	return p.atom1
}

func (p *Pair) Atom2() *Atom {
	return p.atom2
}

/* type */

func (p *Pair) SetType(pt *PairType) {
	p.tipe = pt
}

func (p *Pair) Type() *PairType {
	return p.tipe
}
