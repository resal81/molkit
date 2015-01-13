/*
Package ff provides structs useful to read forcefields and topologies.
*/
package ff

type ffTypes int64

const (
	FF_GROMACS ffTypes = 1 << iota
	FF_CHARMM
	FF_AMBER
)

type prTypes int64

const (
	FF_BOND_TYPE_1       prTypes = 1 << iota // harmonic bond
	FF_NON_BONDED_TYPE_1                     //
	FF_PAIR_TYPE_1                           //
	FF_ANGLE_TYPE_1                          // harmonic
	FF_ANGLE_TYPE_5                          // UB
	FF_DIHEDRAL_TYPE_1                       // proper dihedral
	FF_DIHEDRAL_TYPE_9                       // proper multiple
	FF_IMPROPER_TYPE_1                       // improper
)

/**********************************************************
* GMXProps
**********************************************************/

type GMXProps struct {
	nbfunc   int8
	combrule int8
	genpairs bool
	fudgeLJ  float64
	fudgeQQ  float64
}

func NewGMXProps(nbfunc, combrule int8, genpairs bool, fudgeLJ, fudgeQQ float64) *GMXProps {
	gd := GMXProps{
		nbfunc:   nbfunc,
		combrule: combrule,
		genpairs: genpairs,
		fudgeLJ:  fudgeLJ,
		fudgeQQ:  fudgeQQ,
	}

	return &gd
}

func (p *GMXProps) NbFunc() int8 {
	return p.nbfunc
}

func (p *GMXProps) CombRule() int8 {
	return p.combrule
}

func (p *GMXProps) GenPairs() bool {
	return p.genpairs
}

func (p *GMXProps) FudgeLJ() float64 {
	return p.fudgeLJ
}

func (p *GMXProps) FudgeQQ() float64 {
	return p.fudgeQQ
}

/**********************************************************
* ForceField
**********************************************************/

type ForceField struct {
	kind ffTypes

	gmxdefaults *GMXProps

	props struct {
		gmx *GMXProps
	}

	atomTypes       []*AtomType      // ANY; normal atom types
	nonbondTypes    []*NonBondedType // GROMACS; nb interactions that don't obey combination rules
	pairTypes       []*PairType      // GROMACS; data for 1-4 interactions
	bondTypes       []*BondType      // ANY;
	angleTypes      []*AngleType
	dihedralTypes   []*DihedralType
	improperTypes   []*ImproperType
	constraintTypes []*ConstraintType
	cmapTypes       []*CMapType

	fragments []*TopFragment
}

func NewForceField(kind ffTypes) *ForceField {
	return &ForceField{
		kind: kind,
	}
}

func (f *ForceField) Kind() ffTypes {
	return f.kind
}

//
func (f *ForceField) SetPropsGMX(gd *GMXProps) {
	f.props.gmx = gd
}

func (f *ForceField) PropsGMX() *GMXProps {
	return f.props.gmx
}

//
func (f *ForceField) AddAtomType(at *AtomType) {
	f.atomTypes = append(f.atomTypes, at)
}

func (f *ForceField) AtomTypes() []*AtomType {
	return f.atomTypes
}

//
func (f *ForceField) AddNonBondedType(nb *NonBondedType) {
	f.nonbondTypes = append(f.nonbondTypes, nb)
}

func (f *ForceField) NonBondedTypes() []*NonBondedType {
	return f.nonbondTypes
}

//
func (f *ForceField) AddPairType(pt *PairType) {
	f.pairTypes = append(f.pairTypes, pt)
}

func (f *ForceField) PairTypes() []*PairType {
	return f.pairTypes
}

//
func (f *ForceField) AddBondType(bt *BondType) {
	f.bondTypes = append(f.bondTypes, bt)
}

func (f *ForceField) BondTypes() []*BondType {
	return f.bondTypes
}

//
func (f *ForceField) AddAngleType(ag *AngleType) {
	f.angleTypes = append(f.angleTypes, ag)
}

func (f *ForceField) AngleTypes() []*AngleType {
	return f.angleTypes
}

//
func (f *ForceField) AddDihedralType(dt *DihedralType) {
	if dt.kind&FF_DIHEDRAL_TYPE_1 == 0 && dt.kind&FF_DIHEDRAL_TYPE_9 == 0 {
		panic("cannot add a dihedral with a bad type")
	}
	f.dihedralTypes = append(f.dihedralTypes, dt)
}

func (f *ForceField) DihedralTypes() []*DihedralType {
	return f.dihedralTypes
}

//
func (f *ForceField) AddImproperType(im *ImproperType) {
	if im.kind&FF_IMPROPER_TYPE_1 == 0 {
		panic("cannot add a imporper with a bad type")
	}
	f.improperTypes = append(f.improperTypes, im)
}

func (f *ForceField) ImproperTypes() []*ImproperType {
	return f.improperTypes
}

//
func (f *ForceField) AddCMapType(cm *CMapType) {
	f.cmapTypes = append(f.cmapTypes, cm)
}

func (f *ForceField) CMapTypes() []*CMapType {
	return f.cmapTypes
}
