package blocks

import (
	"strings"
)

const HASH_KEY_SEP = "_"

/**********************************************************
* Hash Keys
**********************************************************/

func GenerateHashKey(v interface{}) []string {

	switch v := v.(type) {

	case *AtomType:
		return []string{
			v.Label(),
		}
	case *BondType:
		return []string{
			v.AType1() + HASH_KEY_SEP + v.AType2(),
		}
	case *AngleType:
		return []string{
			v.AType1() + HASH_KEY_SEP + v.AType2() + HASH_KEY_SEP + v.AType3(),
		}
	case *DihedralType:
		return []string{
			v.AType1() + HASH_KEY_SEP + v.AType2() + HASH_KEY_SEP + v.AType3() + HASH_KEY_SEP + v.AType4(),
		}
	case *ImproperType:
		return []string{
			v.AType1() + HASH_KEY_SEP + v.AType2() + HASH_KEY_SEP + v.AType3() + HASH_KEY_SEP + v.AType4(),
		}
	case *PairType:
		return []string{
			v.AType1() + HASH_KEY_SEP + v.AType2(),
		}
	case *CMapType:
		types1 := []string{
			v.AType1(),
			v.AType2(),
			v.AType3(),
			v.AType4(),
			v.AType5(),
			v.AType6(),
			v.AType7(),
			v.AType8(),
		}
		types2 := []string{
			v.AType1(),
			v.AType2(),
			v.AType3(),
			v.AType4(),
			v.AType8(),
		}

		return []string{
			strings.Join(types1, HASH_KEY_SEP),
			strings.Join(types2, HASH_KEY_SEP),
		}
	case *Fragment:
		return []string{
			v.Name(),
		}

	}

	return []string{}
}

/**********************************************************
* GMXSetup
**********************************************************/

type GMXSetup struct {
	nbFunc   int
	combRule int
	genPairs bool
	fudgeLJ  float64
	fudgeQQ  float64
}

func NewGMXSetup(nbFunc, combRule int, genPairs bool, flj, fqq float64) *GMXSetup {
	g := &GMXSetup{
		nbFunc:   nbFunc,
		combRule: combRule,
		genPairs: genPairs,
		fudgeLJ:  flj,
		fudgeQQ:  fqq,
	}

	return g
}

func (g *GMXSetup) NbFunc() int {
	return g.nbFunc
}

func (g *GMXSetup) CombinationRule() int {
	return g.combRule
}

func (g *GMXSetup) GeneratePairs() bool {
	return g.genPairs
}

func (g *GMXSetup) FudgeLJ() float64 {
	return g.fudgeLJ
}

func (g *GMXSetup) FudgeQQ() float64 {
	return g.fudgeQQ
}

/**********************************************************
* ForceField
**********************************************************/

type FFSetting int64

const (
	FF_TYPE_CHM FFSetting = 1 << iota
	FF_TYPE_GMX
	FF_TYPE_AMB
)

type ForceField struct {
	atomTypes      map[string]*AtomType
	bondTypes      map[string]*BondType
	angleTypes     map[string]*AngleType
	dihedralTypes  map[string]*DihedralType
	improperTypes  map[string]*ImproperType
	nonBondedTypes map[string]*PairType
	oneFourTypes   map[string]*PairType
	cMapTypes      map[string]*CMapType
	fragments      map[string]*Fragment
	gmxSetup       *GMXSetup
	setting        FFSetting
	errors         []error
}

func NewForceField(ft FFSetting) *ForceField {
	return &ForceField{
		atomTypes:      make(map[string]*AtomType),
		bondTypes:      make(map[string]*BondType),
		angleTypes:     make(map[string]*AngleType),
		dihedralTypes:  make(map[string]*DihedralType),
		improperTypes:  make(map[string]*ImproperType),
		nonBondedTypes: make(map[string]*PairType),
		oneFourTypes:   make(map[string]*PairType),
		cMapTypes:      make(map[string]*CMapType),
		fragments:      make(map[string]*Fragment),
		setting:        ft,
	}
}

/* ff setups */

func (ff *ForceField) SetGMXSetup(g *GMXSetup) {
	ff.gmxSetup = g
}

func (ff *ForceField) GMXSetup() *GMXSetup {
	return ff.gmxSetup
}

/* atom types */

func (ff *ForceField) AddAtomType(v *AtomType) {
	for _, el := range GenerateHashKey(v) {
		ff.atomTypes[el] = v
	}
}

func (ff *ForceField) AtomTypes() map[string]*AtomType {
	return ff.atomTypes
}

/* bond types */

func (ff *ForceField) AddBondType(v *BondType) {
	for _, el := range GenerateHashKey(v) {
		ff.bondTypes[el] = v
	}
}

func (ff *ForceField) BondTypes() map[string]*BondType {
	return ff.bondTypes
}

/* angle types */

func (ff *ForceField) AddAngleType(v *AngleType) {
	for _, el := range GenerateHashKey(v) {
		ff.angleTypes[el] = v
	}
}

func (ff *ForceField) AngleTypes() map[string]*AngleType {
	return ff.angleTypes
}

/* dihedral types */

func (ff *ForceField) AddDihedralType(v *DihedralType) {
	for _, el := range GenerateHashKey(v) {
		ff.dihedralTypes[el] = v
	}
}

func (ff *ForceField) DihedralTypes() map[string]*DihedralType {
	return ff.dihedralTypes
}

/* improper types */

func (ff *ForceField) AddImproperType(v *ImproperType) {
	for _, el := range GenerateHashKey(v) {
		ff.improperTypes[el] = v
	}
}

func (ff *ForceField) ImproperTypes() map[string]*ImproperType {
	return ff.improperTypes
}

/* nonbonded types */

func (ff *ForceField) AddNonBondedType(v *PairType) {
	for _, el := range GenerateHashKey(v) {
		ff.nonBondedTypes[el] = v
	}
}

func (ff *ForceField) NonBondedTypes() map[string]*PairType {
	return ff.nonBondedTypes
}

/* one four types */

func (ff *ForceField) AddOneFourType(v *PairType) {
	for _, el := range GenerateHashKey(v) {
		ff.oneFourTypes[el] = v
	}
}

func (ff *ForceField) OneFourTypes() map[string]*PairType {
	return ff.oneFourTypes
}

/* cmap types */

func (ff *ForceField) AddCMapType(v *CMapType) {
	for _, el := range GenerateHashKey(v) {
		ff.cMapTypes[el] = v
	}
}

func (ff *ForceField) CMapTypes() map[string]*CMapType {
	return ff.cMapTypes
}

/* fragments */

func (ff *ForceField) AddFragment(v *Fragment) {
	for _, el := range GenerateHashKey(v) {
		ff.fragments[el] = v
	}
}

func (ff *ForceField) Fragments() map[string]*Fragment {
	return ff.fragments
}

/* setting */

func (ff *ForceField) Setting() FFSetting {
	return ff.setting
}
