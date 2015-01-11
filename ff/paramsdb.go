package ff

type PARAMS_DB_TYPE int8

const (
	P_DB_GROMACS PARAMS_DB_TYPE = 1 << iota
	P_DB_CHARMM
	P_DB_AMBER
)

type ParamsDB struct {
	dbtype PARAMS_DB_TYPE

	gmxprops struct {
		defaults struct {
		}
	}

	atomTypes     map[string]*ParamsAtomType                 // ANY; normal atom types
	atomTypes14   map[string]*ParamsAtomType                 // CHARMM; data for 1-4 interactions
	nonbondTypes  map[string]map[string]*ParamsNonBondedType // GROMACS; nb interactions that don't obey combination rules
	pairTypes     map[string]map[string]*ParamsPairType      // GROMACS; data for 1-4 interactions
	bondTypes     map[string]map[string]*ParamsBondType      // ANY;
	angleTypes    map[string]map[string]map[string]*ParamsAngleType
	dihedralTypes map[string]map[string]map[string]map[string]*ParamsDihedralType
	improperTypes map[string]map[string]map[string]map[string]*ParamsDihedralType
}

func NewParamsDB(dbtype PARAMS_DB_TYPE) *ParamsDB {
	return &ParamsDB{
		dbtype: dbtype,
	}
}

func (p *ParamsDB) Type() PARAMS_DB_TYPE {
	return p.dbtype
}

func (p *ParamsDB) GMXAddAtomType(atype string, protons int8, mass, charge, sigma, epsilon float32) {
	at := ParamsAtomType{
		atype:   atype,
		protons: protons,
		mass:    mass,
		charge:  charge,
		sigma:   sigma,
		epsilon: epsilon,
	}

	p.atomTypes[atype] = &at
}
