package blocks

import (
	"fmt"
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

/* errors */

func (ff *ForceField) Errors() []error {
	return ff.errors
}

/* ff setups */

func (ff *ForceField) SetGMXSetup(g *GMXSetup) {
	ff.gmxSetup = g
}

func (ff *ForceField) GMXSetup() *GMXSetup {
	return ff.gmxSetup
}

/* atom types */

func (ff *ForceField) AddAtomType(at *AtomType) {
	key := GenerateHashKey(at, HK_MODE_NORMAL)

	// check if already exist
	if _, ok := ff.atomTypes[key]; ok {
		err := fmt.Errorf("Overwriting atomtype => %v", at)
		ff.errors = append(ff.errors, err)
	}

	ff.atomTypes[key] = at
}

func (ff *ForceField) AtomTypes() map[HashKey]*AtomType {
	return ff.atomTypes
}

func (ff *ForceField) AtomType(key HashKey) *AtomType {
	return ff.AtomTypes()[key]
}

/* bond types */

func (ff *ForceField) AddBondType(bt *BondType) {
	key := GenerateHashKey(bt, HK_MODE_NORMAL)

	// check if already exist
	if _, ok := ff.bondTypes[key]; ok {
		err := fmt.Errorf("Overwriting bondtype => %v", bt)
		ff.errors = append(ff.errors, err)
	}

	ff.bondTypes[key] = bt
}

func (ff *ForceField) BondTypes() map[HashKey]*BondType {
	return ff.bondTypes
}

func (ff *ForceField) BondType(key HashKey) *BondType {
	return ff.bondTypes[key]
}

/* angle types */

func (ff *ForceField) AddAngleType(at *AngleType) {
	key := GenerateHashKey(at, HK_MODE_NORMAL)

	// check if already exist
	if _, ok := ff.angleTypes[key]; ok {
		err := fmt.Errorf("Overwriting angletype => %v", at)
		ff.errors = append(ff.errors, err)
	}

	ff.angleTypes[key] = at
}

func (ff *ForceField) AngleTypes() map[HashKey]*AngleType {
	return ff.angleTypes
}

func (ff *ForceField) AngleType(key HashKey) *AngleType {
	return ff.angleTypes[key]
}

/* dihedral types */

func (ff *ForceField) AddDihedralType(dt *DihedralType) {
	key := GenerateHashKey(dt, HK_MODE_NORMAL)

	// check if already exist
	if _, ok := ff.dihedralTypes[key]; ok {
		err := fmt.Errorf("Overwriting dihedraltype => %v", dt)
		ff.errors = append(ff.errors, err)
	}

	ff.dihedralTypes[key] = dt
}

func (ff *ForceField) DihedralTypes() map[HashKey]*DihedralType {
	return ff.dihedralTypes
}

func (ff *ForceField) DihedralType(key HashKey) *DihedralType {
	return ff.dihedralTypes[key]
}

/* improper types */

func (ff *ForceField) AddImproperType(it *ImproperType) {
	key := GenerateHashKey(it, HK_MODE_NORMAL)

	// check if already exist
	if _, ok := ff.improperTypes[key]; ok {
		err := fmt.Errorf("Overwriting impropertype => %v", it)
		ff.errors = append(ff.errors, err)
	}

	ff.improperTypes[key] = it
}

func (ff *ForceField) ImproperTypes() map[HashKey]*ImproperType {
	return ff.improperTypes
}

func (ff *ForceField) ImproperType(key HashKey) *ImproperType {
	return ff.improperTypes[key]
}

/* nonbonded types */

func (ff *ForceField) AddNonBondedType(nb *PairType) {
	key := GenerateHashKey(nb, HK_MODE_NORMAL)

	// check if already exist
	if _, ok := ff.nonBondedTypes[key]; ok {
		err := fmt.Errorf("Overwriting nonbondedtype => %v", nb)
		ff.errors = append(ff.errors, err)
	}

	ff.nonBondedTypes[key] = nb
}

func (ff *ForceField) NonBondedTypes() map[HashKey]*PairType {
	return ff.nonBondedTypes
}

func (ff *ForceField) NonBondedType(key HashKey) *PairType {
	return ff.nonBondedTypes[key]
}

/* one four types */

func (ff *ForceField) AddOneFourType(nb *PairType) {
	key := GenerateHashKey(nb, HK_MODE_NORMAL)

	// check if already exist
	if _, ok := ff.oneFourTypes[key]; ok {
		err := fmt.Errorf("Overwriting onefourtype => %v", nb)
		ff.errors = append(ff.errors, err)
	}

	ff.oneFourTypes[key] = nb
}

func (ff *ForceField) OneFourTypes() map[HashKey]*PairType {
	return ff.oneFourTypes
}

func (ff *ForceField) OneFourType(key HashKey) *PairType {
	return ff.oneFourTypes[key]
}

/* cmap types */

func (ff *ForceField) AddCMapType(v *CMapType) {
	key := GenerateHashKey(v, HK_MODE_NORMAL)

	// check if already exist
	if _, ok := ff.cMapTypes[key]; ok {
		err := fmt.Errorf("Overwriting cmaptype => %v", v)
		ff.errors = append(ff.errors, err)
	}

	ff.cMapTypes[key] = v
}

func (ff *ForceField) CMapTypes() map[HashKey]*CMapType {
	return ff.cMapTypes
}

func (ff *ForceField) CMapType(key HashKey) *CMapType {
	return ff.cMapTypes[key]
}

/* fragments */

func (ff *ForceField) AddFragment(v *Fragment) {
	key := GenerateHashKey(v, HK_MODE_NORMAL)

	// check if already exist
	if _, ok := ff.fragments[key]; ok {
		err := fmt.Errorf("Overwriting fragment => %v", v)
		ff.errors = append(ff.errors, err)
	}

	ff.fragments[key] = v
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
