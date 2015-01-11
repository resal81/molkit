package ff

type PARAMS_DB_TYPE int8

const (
	P_DB_GROMACS PARAMS_DB_TYPE = 1 << iota
	P_DB_CHARMM
	P_DB_AMBER
)

type ParamsDB struct {
	dbtype PARAMS_DB_TYPE

	gmxdefaults *GMXDefaults

	atomTypes     []*AtomType            // ANY; normal atom types
	nonbondTypes  []*ParamsNonBondedType // GROMACS; nb interactions that don't obey combination rules
	pairTypes     []*ParamsPairType      // GROMACS; data for 1-4 interactions
	bondTypes     []*ParamsBondType      // ANY;
	angleTypes    []*ParamsAngleType
	dihedralTypes []*ParamsDihedralType
	improperTypes []*ParamsDihedralType
}

func NewParamsDB(dbtype PARAMS_DB_TYPE) *ParamsDB {
	return &ParamsDB{
		dbtype: dbtype,
	}
}

func (p *ParamsDB) Type() PARAMS_DB_TYPE {
	return p.dbtype
}

func (p *ParamsDB) SetGMXDefaults(gd *GMXDefaults) {
	p.gmxdefaults = gd
}
