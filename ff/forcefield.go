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

	atomTypes     []*AtomType      // ANY; normal atom types
	nonbondTypes  []*NonBondedType // GROMACS; nb interactions that don't obey combination rules
	pairTypes     []*PairType      // GROMACS; data for 1-4 interactions
	bondTypes     []*BondType      // ANY;
	angleTypes    []*AngleType
	dihedralTypes []*DihedralType
	improperTypes []*DihedralType

	fragments []*TopFragment
}

func NewForceField(fftype FORCEFIELD_TYPE) *ForceField {
	return &ForceField{
		fftype: fftype,
	}
}

func (p *ForceField) Type() FORCEFIELD_TYPE {
	return p.fftype
}

func (p *ForceField) SetGMXDefaults(gd *GMXDefaults) {
	p.gmxdefaults = gd
}

func (p *ForceField) AddAtomType(at *AtomType) {
	p.atomTypes = append(p.atomTypes, at)
}
