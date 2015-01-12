/*
Package ff provides structs useful to read forcefields and topologies.
*/
package ff

type ffTypes int64

const (
	FF_SOURCE_GROMACS ffTypes = 1 << iota
	FF_SOURCE_CHARMM
	FF_SOURCE_AMBER
	FF_BOND_TYPE_1       // harmonic bond
	FF_NON_BONDED_TYPE_1 //
	FF_PAIR_TYPE_1       //
	FF_ANGLE_TYPE_1      // harmonic
	FF_ANGLE_TYPE_5      // UB
	FF_DIHEDRAL_TYPE_1   // proper dihedral
	FF_DIHEDRAL_TYPE_2   // improper
	FF_DIHEDRAL_TYPE_9   // proper multiple
)

type ForceField struct {
	kind ffTypes

	gmxdefaults *GMXDefaults

	atomTypes       []*AtomType      // ANY; normal atom types
	nonbondTypes    []*NonBondedType // GROMACS; nb interactions that don't obey combination rules
	pairTypes       []*PairType      // GROMACS; data for 1-4 interactions
	bondTypes       []*BondType      // ANY;
	angleTypes      []*AngleType
	dihedralTypes   []*DihedralType
	improperTypes   []*DihedralType
	constraintTypes []*ConstraintType

	fragments []*TopFragment
}

func NewForceField(kind ffTypes) *ForceField {
	if kind&FF_SOURCE_GROMACS == 0 && kind&FF_SOURCE_CHARMM == 0 && kind&FF_SOURCE_AMBER == 0 {
		panic("unsupported ff type")
	}
	return &ForceField{
		kind: kind,
	}
}

func (f *ForceField) Kind() ffTypes {
	return f.kind
}

func (f *ForceField) SetGMXDefaults(gd *GMXDefaults) {
	f.gmxdefaults = gd
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
func (f *ForceField) AddImproperType(im *DihedralType) {
	if im.kind&FF_DIHEDRAL_TYPE_2 == 0 {
		panic("cannot add a imporper with a bad type")
	}
	f.improperTypes = append(f.improperTypes, im)
}

func (f *ForceField) ImproperTypes() []*DihedralType {
	return f.improperTypes
}
