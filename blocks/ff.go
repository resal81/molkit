package blocks

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
	FF_CHM FFSetting = 1 << iota
	FF_GMX
	FF_AMB
)

type ForceField struct {
	atomTypes      []*AtomType
	bondTypes      []*BondType
	angleTypes     []*AngleType
	dihedralTypes  []*DihedralType
	improperTypes  []*ImproperType
	nonBondedTypes []*PairType
	oneFourTypes   []*PairType
	cMapTypes      []*CMapType
	fragments      []*Fragment
	gmxSetup       *GMXSetup
	setting        FFSetting
}

func NewForceField(ft FFSetting) *ForceField {
	return &ForceField{
		setting: ft,
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
	ff.atomTypes = append(ff.atomTypes, v)
}

func (ff *ForceField) AtomTypes() []*AtomType {
	return ff.atomTypes
}

/* angle types */

func (ff *ForceField) AddAngleType(v *AngleType) {
	ff.angleTypes = append(ff.angleTypes, v)
}

func (ff *ForceField) AngleTypes() []*AngleType {
	return ff.angleTypes
}

/* dihedral types */

func (ff *ForceField) AddDihedralType(v *DihedralType) {
	ff.dihedralTypes = append(ff.dihedralTypes, v)
}

func (ff *ForceField) DihedralTypes() []*DihedralType {
	return ff.dihedralTypes
}

/* improper types */

func (ff *ForceField) AddImproperType(v *ImproperType) {
	ff.improperTypes = append(ff.improperTypes, v)
}

func (ff *ForceField) ImproperTypes() []*ImproperType {
	return ff.improperTypes
}

/* nonbonded types */

func (ff *ForceField) AddNonBondedType(v *PairType) {
	ff.nonBondedTypes = append(ff.nonBondedTypes, v)
}

func (ff *ForceField) NonBondedTypes() []*PairType {
	return ff.nonBondedTypes
}

/* one four types */

func (ff *ForceField) AddOneFourType(v *PairType) {
	ff.oneFourTypes = append(ff.oneFourTypes, v)
}

func (ff *ForceField) OneFourTypes() []*PairType {
	return ff.oneFourTypes
}

/* cmap types */

func (ff *ForceField) AddCMapType(v *CMapType) {
	ff.cMapTypes = append(ff.cMapTypes, v)
}

func (ff *ForceField) CMapTypes() []*CMapType {
	return ff.cMapTypes
}

/* fragments */

func (ff *ForceField) AddFragment(v *Fragment) {
	ff.fragments = append(ff.fragments, v)
}

func (ff *ForceField) Fragments() []*Fragment {
	return ff.fragments
}

/* setting */
