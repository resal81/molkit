package blocks

import (
	"strings"
)

const HASH_KEY_SEP = "_"

/**********************************************************
* Hash Keys
**********************************************************/

type HashKeySetting int64

const (
	HK_MODE_NORMAL HashKeySetting = 1 << iota
	HK_MODE_REVERSE_BOND
	HK_MODE_REVERSE_ANGLE
	HK_MODE_REVERSE_DIHEDRAL
	HK_MODE_SHORT_CMAP // 5 element cmap
)

type HashKey string

func GenerateHashKey(v interface{}, hks HashKeySetting) HashKey {

	switch v := v.(type) {

	case *AtomType:
		return HashKey(v.Label())

	case *BondType:
		if hks&HK_MODE_REVERSE_BOND != 0 {
			return HashKey(v.AType2() + HASH_KEY_SEP + v.AType1())
		}
		return HashKey(v.AType1() + HASH_KEY_SEP + v.AType2())

	case *AngleType:
		if hks&HK_MODE_REVERSE_ANGLE != 0 {
			return HashKey(v.AType3() + HASH_KEY_SEP + v.AType2() + HASH_KEY_SEP + v.AType1())
		}
		return HashKey(v.AType1() + HASH_KEY_SEP + v.AType2() + HASH_KEY_SEP + v.AType3())

	case *DihedralType:
		if hks&HK_MODE_REVERSE_DIHEDRAL != 0 {
			return HashKey(v.AType4() + HASH_KEY_SEP + v.AType3() + HASH_KEY_SEP + v.AType2() + HASH_KEY_SEP + v.AType1())
		}
		return HashKey(v.AType1() + HASH_KEY_SEP + v.AType2() + HASH_KEY_SEP + v.AType3() + HASH_KEY_SEP + v.AType4())

	case *ImproperType:
		return HashKey(v.AType1() + HASH_KEY_SEP + v.AType2() + HASH_KEY_SEP + v.AType3() + HASH_KEY_SEP + v.AType4())

	case *PairType:
		return HashKey(v.AType1() + HASH_KEY_SEP + v.AType2())

	case *CMapType:
		if hks&HK_MODE_SHORT_CMAP != 0 {
			types1 := []string{
				v.AType1(),
				v.AType2(),
				v.AType3(),
				v.AType4(),
				v.AType8(),
			}
			return HashKey(strings.Join(types1, HASH_KEY_SEP))

		}
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
		return HashKey(strings.Join(types1, HASH_KEY_SEP))

	case *Fragment:
		return HashKey(v.Name())

	}

	return ""
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
	atomTypes      map[HashKey]*AtomType
	bondTypes      map[HashKey]*BondType
	angleTypes     map[HashKey]*AngleType
	dihedralTypes  map[HashKey]*DihedralType
	improperTypes  map[HashKey]*ImproperType
	nonBondedTypes map[HashKey]*PairType
	oneFourTypes   map[HashKey]*PairType
	cMapTypes      map[HashKey]*CMapType
	fragments      map[HashKey]*Fragment
	gmxSetup       *GMXSetup
	setting        FFSetting
	errors         []error
}

func NewForceField(ft FFSetting) *ForceField {
	return &ForceField{
		atomTypes:      make(map[HashKey]*AtomType),
		bondTypes:      make(map[HashKey]*BondType),
		angleTypes:     make(map[HashKey]*AngleType),
		dihedralTypes:  make(map[HashKey]*DihedralType),
		improperTypes:  make(map[HashKey]*ImproperType),
		nonBondedTypes: make(map[HashKey]*PairType),
		oneFourTypes:   make(map[HashKey]*PairType),
		cMapTypes:      make(map[HashKey]*CMapType),
		fragments:      make(map[HashKey]*Fragment),
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
	el := GenerateHashKey(v, HK_MODE_NORMAL)
	ff.atomTypes[el] = v
}

func (ff *ForceField) AtomTypes() map[HashKey]*AtomType {
	return ff.atomTypes
}

func (ff *ForceField) AtomType(key HashKey) *AtomType {
	return ff.AtomTypes()[key]
}

/* bond types */

func (ff *ForceField) AddBondType(v *BondType) {
	el := GenerateHashKey(v, HK_MODE_NORMAL)
	ff.bondTypes[el] = v
}

func (ff *ForceField) BondTypes() map[HashKey]*BondType {
	return ff.bondTypes
}

func (ff *ForceField) BondType(key HashKey) *BondType {
	return ff.bondTypes[key]
}

/* angle types */

func (ff *ForceField) AddAngleType(v *AngleType) {
	el := GenerateHashKey(v, HK_MODE_NORMAL)
	ff.angleTypes[el] = v
}

func (ff *ForceField) AngleTypes() map[HashKey]*AngleType {
	return ff.angleTypes
}

func (ff *ForceField) AngleType(key HashKey) *AngleType {
	return ff.angleTypes[key]
}

/* dihedral types */

func (ff *ForceField) AddDihedralType(v *DihedralType) {
	el := GenerateHashKey(v, HK_MODE_NORMAL)
	ff.dihedralTypes[el] = v
}

func (ff *ForceField) DihedralTypes() map[HashKey]*DihedralType {
	return ff.dihedralTypes
}

func (ff *ForceField) DihedralType(key HashKey) *DihedralType {
	return ff.dihedralTypes[key]
}

/* improper types */

func (ff *ForceField) AddImproperType(v *ImproperType) {
	el := GenerateHashKey(v, HK_MODE_NORMAL)
	ff.improperTypes[el] = v
}

func (ff *ForceField) ImproperTypes() map[HashKey]*ImproperType {
	return ff.improperTypes
}

func (ff *ForceField) ImproperType(key HashKey) *ImproperType {
	return ff.improperTypes[key]
}

/* nonbonded types */

func (ff *ForceField) AddNonBondedType(v *PairType) {
	el := GenerateHashKey(v, HK_MODE_NORMAL)
	ff.nonBondedTypes[el] = v
}

func (ff *ForceField) NonBondedTypes() map[HashKey]*PairType {
	return ff.nonBondedTypes
}

func (ff *ForceField) NonBondedType(key HashKey) *PairType {
	return ff.nonBondedTypes[key]
}

/* one four types */

func (ff *ForceField) AddOneFourType(v *PairType) {
	el := GenerateHashKey(v, HK_MODE_NORMAL)
	ff.oneFourTypes[el] = v
}

func (ff *ForceField) OneFourTypes() map[HashKey]*PairType {
	return ff.oneFourTypes
}

func (ff *ForceField) OneFourType(key HashKey) *PairType {
	return ff.oneFourTypes[key]
}

/* cmap types */

func (ff *ForceField) AddCMapType(v *CMapType) {
	el := GenerateHashKey(v, HK_MODE_NORMAL)
	ff.cMapTypes[el] = v
}

func (ff *ForceField) CMapTypes() map[HashKey]*CMapType {
	return ff.cMapTypes
}

func (ff *ForceField) CMapType(key HashKey) *CMapType {
	return ff.cMapTypes[key]
}

/* fragments */

func (ff *ForceField) AddFragment(v *Fragment) {
	el := GenerateHashKey(v, HK_MODE_NORMAL)
	ff.fragments[el] = v
}

func (ff *ForceField) Fragments() map[HashKey]*Fragment {
	return ff.fragments
}

func (ff *ForceField) Fragment(key HashKey) *Fragment {
	return ff.fragments[key]
}

/* setting */

func (ff *ForceField) Setting() FFSetting {
	return ff.setting
}
