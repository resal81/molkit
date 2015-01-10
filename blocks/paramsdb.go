package blocks

type ParamsDB struct {
	gmxprops struct {
		defaults struct {
		}
	}

	atomTypes     map[string]*AtomType
	atomTypes14   map[string]*AtomType
	bondTypes     map[string]map[string]*BondType
	angleTypes    map[string]map[string]map[string]*AngleType
	dihedralTypes map[string]map[string]map[string]map[string]*DihedralType
	improperTypes map[string]map[string]map[string]map[string]*DihedralType
}

func NewParamsDB() *ParamsDB {
	return &ParamsDB{}
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

func (p *ParamsDB) CHMAddAtomType() {

}

func (p *ParamsDB) GMXAddBondType(fn int8, k_r float32, r0 float32) {

}
