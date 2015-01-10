package blocks

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

	atomTypes     map[string]*AtomType            // ANY; normal atom types
	atomTypes14   map[string]*AtomType            // CHARMM; data for 1-4 interactions
	nonbondTypes  map[string]map[string]*PairType // GROMACS; nb interactions that don't obey combination rules
	pairTypess    map[string]map[string]*PairType // GROMACS; data for 1-4 interactions
	bondTypes     map[string]map[string]*BondType // ANY;
	angleTypes    map[string]map[string]map[string]*AngleType
	dihedralTypes map[string]map[string]map[string]map[string]*DihedralType
	improperTypes map[string]map[string]map[string]map[string]*DihedralType
}

func NewParamsDB(dbtype PARAMS_DB_TYPE) *ParamsDB {
	return &ParamsDB{dbtype: dbtype}
}

func (p *ParamsDB) GMXAddAtomType(atype string, protons int8, mass, charge, sigma, epsilon float32) {
	at := AtomType{
		atype:   atype,
		protons: protons,
		mass:    mass,
		charge:  charge,
		sigma:   sigma,
		epsilon: epsilon,
	}

	p.atomTypes[atype] = &at
}

func (p *ParamsDB) GMXAddNBType(atype1, atyp2 string, fn int8, sigma, epsilon float32) {

}

func (p *ParamsDB) GMXAddPairType(atype1, atyp2 string, fn int8, sigma, epsilon float32) {

}

func (p *ParamsDB) GMXAddBondType(atype1, atyp2 string, fn int8, k_r, r0 float32) {

}

func (p *ParamsDB) GMXAddAngleType(atyp1, atyp2, atype3 string, fn int8, k_theta, theta, r13, k_ub float32) {

}

func (p *ParamsDB) GMXAddDihedralType(atyp1, atyp2, atype3, atype4 string, fn int8, k_phi, phi float32, mult int8) {

}
