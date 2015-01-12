package ff

type FORCEFIELD_TYPE int8

const (
	FF_GROMACS FORCEFIELD_TYPE = 1 << iota
	FF_CHARMM
	FF_AMBER
)

type ForceField struct {
	fftype FORCEFIELD_TYPE

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

func NewForceField(fftype FORCEFIELD_TYPE) *ForceField {
	return &ForceField{
		fftype: fftype,
	}
}

func (f *ForceField) Type() FORCEFIELD_TYPE {
	return f.fftype
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

func (f *ForceField) NondBondedTypes() []*NonBondedType {
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
	if dt.setting&DHT_TYPE_1 == 0 && dt.setting&DHT_TYPE_9 == 0 {
		panic("cannot add a dihedral with a bad type")
	}
	f.dihedralTypes = append(f.dihedralTypes, dt)
}

func (f *ForceField) DihedralTypes() []*DihedralType {
	return f.dihedralTypes
}

//
func (f *ForceField) AddImproperType(im *DihedralType) {
	if im.setting&DHT_TYPE_1 == 2 {
		panic("cannot add a imporper with a bad type")
	}
	f.improperTypes = append(f.improperTypes, im)
}

func (f *ForceField) ImproperTytpes() []*DihedralType {
	return f.improperTypes
}
