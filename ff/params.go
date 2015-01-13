package ff

/**********************************************************
* AtomType
**********************************************************/

type atSetting int32

const (
	at_sett_PROTONS_SET atSetting = 1 << iota
	at_sett_MASS_SET
	at_sett_CHARGE_SET
	at_sett_SIGMA_SET
	at_sett_SIGMA14_SET
	at_sett_EPSILON_SET
	at_sett_EPSILON14_SET
)

type AtomType struct {
	atype     string
	protons   int8
	mass      float64
	sigma     float64
	epsilon   float64
	sigma14   float64
	epsilon14 float64
	charge    float64

	setting atSetting
	ffType  ffTypes
}

// Constructor
func NewAtomType(atype string, ffType ffTypes) *AtomType {
	return &AtomType{
		atype:  atype,
		ffType: ffType,
	}
}

//
func (a *AtomType) SetProtons(u int8) {
	a.setting |= at_sett_PROTONS_SET
	a.protons = u
}

func (a *AtomType) HasProtonsSet() bool {
	return a.setting&at_sett_PROTONS_SET != 0
}

func (a *AtomType) Protons() int8 {
	return a.protons
}

//
func (a *AtomType) SetMass(v float64) {
	a.setting |= at_sett_MASS_SET
	a.mass = v
}

func (a *AtomType) HasMassSet() bool {
	return a.setting&at_sett_MASS_SET != 0
}

func (a *AtomType) Mass() float64 {
	return a.mass
}

//
func (a *AtomType) SetCharge(v float64) {
	a.setting |= at_sett_CHARGE_SET
	a.charge = v
}

func (a *AtomType) HasChargeSet() bool {
	return a.setting&at_sett_CHARGE_SET != 0
}

func (a *AtomType) Charge() float64 {
	return a.charge
}

//
func (a *AtomType) SetSigma(v float64) {
	a.setting |= at_sett_SIGMA_SET
	a.sigma = v
}

func (a *AtomType) HasSigmaSet() bool {
	return a.setting&at_sett_SIGMA_SET != 0
}

func (a *AtomType) Sigma(to ffTypes) float64 {
	return convertSigma(a.sigma, a.ffType, to)
}

//
func (a *AtomType) SetEpsilon(v float64) {
	a.setting |= at_sett_EPSILON_SET
	a.epsilon = v
}

func (a *AtomType) HasEpsilonSet() bool {
	return a.setting&at_sett_EPSILON_SET != 0
}

func (a *AtomType) Epsilon(to ffTypes) float64 {
	return convertEpsilon(a.epsilon, a.ffType, to)
}

//
func (a *AtomType) SetSigma14(v float64) {
	a.setting |= at_sett_SIGMA14_SET
	a.sigma14 = v
}

func (a *AtomType) HasSigma14Set() bool {
	return a.setting&at_sett_SIGMA14_SET != 0
}

func (a *AtomType) Sigma14(to ffTypes) float64 {
	return convertSigma(a.sigma14, a.ffType, to)
}

//
func (a *AtomType) SetEpsilon14(v float64) {
	a.setting |= at_sett_EPSILON14_SET
	a.epsilon14 = v
}

func (a *AtomType) HasEpsilon14Set() bool {
	return a.setting&at_sett_EPSILON14_SET != 0
}

func (a *AtomType) Epsilon14(to ffTypes) float64 {
	return convertEpsilon(a.epsilon14, a.ffType, to)
}

/**********************************************************
* NonBondedType
**********************************************************/

type nonbondedtypeSetting int32

const (
	nbt_sett_SIGMA_SET nonbondedtypeSetting = 1 << iota
	nbt_sett_EPSILON_SET
)

type NonBondedType struct {
	atype1 string
	atype2 string

	nbtype nonbondedtypeSetting

	sigma   float64
	epsilon float64

	setting nonbondedtypeSetting
	kind    prTypes
	ffType  ffTypes
}

func NewNonBondedType(atype1, atype2 string, nbt prTypes, ffType ffTypes) *NonBondedType {

	if nbt&FF_NON_BONDED_TYPE_1 == 0 {
		panic("unsupported nonbonded-param")
	}
	return &NonBondedType{
		atype1: atype1,
		atype2: atype2,
		kind:   nbt,
		ffType: ffType,
	}
}

//
func (n *NonBondedType) AType1() string {
	return n.atype1
}

func (n *NonBondedType) AType2() string {
	return n.atype2
}

//
func (n *NonBondedType) SetSigma(v float64) {
	n.setting |= nbt_sett_SIGMA_SET
	n.sigma = v
}

func (n *NonBondedType) HasSigmaSet() bool {
	return n.setting&nbt_sett_SIGMA_SET != 0
}

func (n *NonBondedType) Sigma(to ffTypes) float64 {
	return convertNBSigma(n.sigma, n.ffType, to)
}

//
func (n *NonBondedType) SetEpsilon(v float64) {
	n.setting |= nbt_sett_EPSILON_SET
	n.epsilon = v
}

func (n *NonBondedType) HasEpsilonSet() bool {
	return n.setting&nbt_sett_EPSILON_SET != 0
}

func (n *NonBondedType) Epsilon(to ffTypes) float64 {
	return convertEpsilon(n.epsilon, n.ffType, to)
}

/**********************************************************
* PairType
**********************************************************/

type pairtypeSetting int32

const (
	pt_sett_SIGMA14_SET pairtypeSetting = 1 << iota
	pt_sett_EPSILON14_SET
)

type PairType struct {
	atype1    string
	atype2    string
	sigma14   float64
	epsilon14 float64

	setting pairtypeSetting
	kind    prTypes
	ffType  ffTypes
}

func NewPairType(atype1, atype2 string, pt prTypes, ffType ffTypes) *PairType {
	if pt&FF_PAIR_TYPE_1 == 0 {
		panic("unsupported pairtype")
	}

	return &PairType{
		atype1: atype1,
		atype2: atype2,
		kind:   pt,
		ffType: ffType,
	}
}

//
func (p *PairType) AType1() string {
	return p.atype1
}

func (p *PairType) AType2() string {
	return p.atype2
}

//
func (p *PairType) SetSigma14(v float64) {
	p.setting |= pt_sett_SIGMA14_SET
	p.sigma14 = v
}

func (p *PairType) HasSigma14Set() bool {
	return p.setting&pt_sett_SIGMA14_SET != 0
}

func (p *PairType) Sigma14(to ffTypes) float64 {
	return convertSigma(p.sigma14, p.ffType, to)
}

//
func (p *PairType) SetEpsilon14(v float64) {
	p.setting |= pt_sett_EPSILON14_SET
	p.epsilon14 = v
}

func (p *PairType) HasEpsilon14Set() bool {
	return p.setting&pt_sett_EPSILON14_SET != 0
}

func (p *PairType) Epsilon14(to ffTypes) float64 {
	return convertEpsilon(p.epsilon14, p.ffType, to)
}

/**********************************************************
* BondType
**********************************************************/

type bondtypeSetting int32

const (
	bt_sett_HARMONIC_CONSTANT_SET bondtypeSetting = 1 << iota
	bt_sett_HARMONIC_DISTANCE_SET
)

type BondType struct {
	atype1 string
	atype2 string

	kr float64
	r0 float64

	setting bondtypeSetting
	kind    prTypes
	ffType  ffTypes
}

func NewBondType(atype1, atype2 string, bt prTypes, ffType ffTypes) *BondType {
	if bt&FF_BOND_TYPE_1 == 0 {
		panic("unsupported bondtype")
	}

	return &BondType{
		atype1: atype1,
		atype2: atype2,
		kind:   bt,
		ffType: ffType,
	}
}

//
func (b *BondType) AType1() string {
	return b.atype1
}

func (b *BondType) AType2() string {
	return b.atype2
}

//
func (b *BondType) SetHarmonicConstant(v float64) {
	b.setting |= bt_sett_HARMONIC_CONSTANT_SET
	b.kr = v
}

func (b *BondType) HasHarmonicConstantSet() bool {
	return b.setting&bt_sett_HARMONIC_CONSTANT_SET != 0
}

func (b *BondType) HarmonicConstant(to ffTypes) float64 {
	return convertHarmonicConstant(b.kr, b.ffType, to)
}

//
func (b *BondType) SetHarmonicDistance(v float64) {
	b.setting |= bt_sett_HARMONIC_DISTANCE_SET
	b.r0 = v
}

func (b *BondType) HasHarmonicDistanceSet() bool {
	return b.setting&bt_sett_HARMONIC_DISTANCE_SET != 0
}

func (b *BondType) HarmonicDistance(to ffTypes) float64 {
	return convertHarmonicDistance(b.r0, b.ffType, to)
}

/**********************************************************
* AngleType
**********************************************************/

type angletypeSetting int32

const (
	ang_sett_THETA_CONSTANT_SET angletypeSetting = 1 << iota
	ang_sett_THETA_SET
	ang_sett_UB_CONSTANT_SET
	ang_sett_R13_SET
)

type AngleType struct {
	atype1 string
	atype2 string
	atype3 string

	k_theta float64
	theta   float64
	r13     float64
	k_ub    float64

	setting angletypeSetting
	kind    prTypes
	ffType  ffTypes
}

func NewAngleType(atype1, atype2, atype3 string, angtype prTypes, ffType ffTypes) *AngleType {

	if angtype&FF_ANGLE_TYPE_1 == 0 && angtype&FF_ANGLE_TYPE_5 == 0 {
		panic("unsupported angle type")
	}

	return &AngleType{
		atype1: atype1,
		atype2: atype2,
		atype3: atype3,
		kind:   angtype,
		ffType: ffType,
	}
}

//
func (a *AngleType) AType1() string {
	return a.atype1
}

func (a *AngleType) AType2() string {
	return a.atype2
}

func (a *AngleType) AType3() string {
	return a.atype3
}

//
func (a *AngleType) SetThetaConstant(v float64) {
	a.setting |= ang_sett_THETA_CONSTANT_SET
	a.k_theta = v
}

func (a *AngleType) HasThetaConstantSet() bool {
	return a.setting&ang_sett_THETA_CONSTANT_SET != 0
}

func (a *AngleType) ThetaConstant(to ffTypes) float64 {
	return convertThetaConstant(a.k_theta, a.ffType, to)
}

//
func (a *AngleType) SetTheta(v float64) {
	a.setting |= ang_sett_THETA_SET
	a.theta = v
}

func (a *AngleType) HasThetaSet() bool {
	return a.setting&ang_sett_THETA_SET != 0
}

func (a *AngleType) Theta(to ffTypes) float64 {
	return convertTheta(a.theta, a.ffType, to)
}

//
func (a *AngleType) SetUBConstant(v float64) {
	a.setting |= ang_sett_UB_CONSTANT_SET
	a.k_ub = v
}

func (a *AngleType) HasUBConstantSet() bool {
	return a.setting&ang_sett_UB_CONSTANT_SET != 0
}

func (a *AngleType) UBConstant(to ffTypes) float64 {
	return convertUBConstant(a.k_ub, a.ffType, to)
}

//
func (a *AngleType) SetR13(v float64) {
	a.setting |= ang_sett_R13_SET
	a.r13 = v
}

func (a *AngleType) HasR13Set() bool {
	return a.setting&ang_sett_R13_SET != 0
}

func (a *AngleType) R13(to ffTypes) float64 {
	return convertR13(a.r13, a.ffType, to)
}

/**********************************************************
* DihedralType
**********************************************************/

type dihedraltypeSetting int32

const (
	dih_sett_PHI_CONSTANT_SET dihedraltypeSetting = 1 << iota
	dih_sett_PHI_SET
	dih_sett_MULT_SET
)

type DihedralType struct {
	atype1 string
	atype2 string
	atype3 string
	atype4 string

	k_phi float64
	phi   float64
	mult  int8

	setting dihedraltypeSetting
	kind    prTypes
	ffType  ffTypes
}

func NewDihedralType(atype1, atype2, atype3, atype4 string, dht prTypes, ffType ffTypes) *DihedralType {
	if dht&FF_DIHEDRAL_TYPE_1 == 0 && dht&FF_DIHEDRAL_TYPE_9 == 0 {
		panic("unsupported dihedraltype")
	}

	return &DihedralType{
		atype1: atype1,
		atype2: atype2,
		atype3: atype3,
		atype4: atype4,
		kind:   dht,
		ffType: ffType,
	}
}

//
func (d *DihedralType) Kind() prTypes {
	return d.kind
}

//
func (d *DihedralType) AType1() string {
	return d.atype1
}

func (d *DihedralType) AType2() string {
	return d.atype2
}

func (d *DihedralType) AType3() string {
	return d.atype3
}

func (d *DihedralType) AType4() string {
	return d.atype4
}

//
func (d *DihedralType) SetPhiConstant(v float64) {
	d.setting |= dih_sett_PHI_CONSTANT_SET
	d.k_phi = v
}

func (d *DihedralType) HasPhiConstantSet() bool {
	return d.setting&dih_sett_PHI_CONSTANT_SET != 0
}

func (d *DihedralType) PhiConstant(to ffTypes) float64 {
	return convertPhiConstant(d.k_phi, d.ffType, to)
}

//
func (d *DihedralType) SetPhi(v float64) {
	d.setting |= dih_sett_PHI_SET
	d.phi = v
}

func (d *DihedralType) HasPhiSet() bool {
	return d.setting&dih_sett_PHI_SET != 0
}

func (d *DihedralType) Phi(to ffTypes) float64 {
	return convertPhi(d.phi, d.ffType, to)
}

//
func (d *DihedralType) SetMult(v int8) {
	d.setting |= dih_sett_MULT_SET
	d.mult = v
}

func (d *DihedralType) HasMultSet() bool {
	return d.setting&dih_sett_MULT_SET != 0
}

func (d *DihedralType) Mult(to ffTypes) int8 {
	return convertMutl(d.mult, d.ffType, to)
}

func (d *DihedralType) Setting() dihedraltypeSetting {
	return d.setting
}

/**********************************************************
* ImproperType
**********************************************************/

type impropertypeSetting int32

const (
	imp_sett_PSI_CONSTANT_SET impropertypeSetting = 1 << iota
	imp_sett_PSI_SET
)

type ImproperType struct {
	atype1 string
	atype2 string
	atype3 string
	atype4 string

	// for improper
	k_psi float64
	psi   float64

	setting impropertypeSetting
	kind    prTypes
	ffType  ffTypes
}

func NewImproperType(atype1, atype2, atype3, atype4 string, dht prTypes, ffType ffTypes) *ImproperType {
	if dht&FF_IMPROPER_TYPE_1 == 0 {
		panic("unsupported dihedraltype")
	}

	return &ImproperType{
		atype1: atype1,
		atype2: atype2,
		atype3: atype3,
		atype4: atype4,
		kind:   dht,
		ffType: ffType,
	}
}

//
func (d *ImproperType) Kind() prTypes {
	return d.kind
}

//
func (d *ImproperType) AType1() string {
	return d.atype1
}

func (d *ImproperType) AType2() string {
	return d.atype2
}

func (d *ImproperType) AType3() string {
	return d.atype3
}

func (d *ImproperType) AType4() string {
	return d.atype4
}

//
func (d *ImproperType) SetPsiConstant(v float64) {
	d.setting |= imp_sett_PSI_CONSTANT_SET
	d.k_psi = v
}

func (d *ImproperType) HasPsiConstantSet() bool {
	return d.setting&imp_sett_PSI_CONSTANT_SET != 0
}

func (d *ImproperType) PsiConstant(to ffTypes) float64 {
	return convertPsiConstant(d.k_psi, d.ffType, to)
}

//
func (d *ImproperType) SetPsi(v float64) {
	d.setting |= imp_sett_PSI_SET
	d.psi = v
}

func (d *ImproperType) HasPsiSet() bool {
	return d.setting&imp_sett_PSI_SET != 0
}

func (d *ImproperType) Psi(to ffTypes) float64 {
	return convertPsi(d.psi, d.ffType, to)
}

func (d *ImproperType) Setting() impropertypeSetting {
	return d.setting
}

/**********************************************************
* ConstraintType
**********************************************************/

type ConstraintType struct {
	atype1 string
	atype2 string

	b0 float64
}

/**********************************************************
* CMAP
**********************************************************/

type CMapType struct {
	nx     int
	ny     int
	values []float64
	atypes []string
	ffType ffTypes
}

func NewCMapType(nx, ny int, ffType ffTypes) *CMapType {
	return &CMapType{
		nx:     nx,
		ny:     ny,
		ffType: ffType,
	}
}

//
func (c *CMapType) SetAtomTypes(atypes ...string) {
	if len(atypes) != 5 || len(atypes) != 8 {
		panic("cmap number of atypes must be 5 or 8")
	}

	c.atypes = atypes
}

func (c *CMapType) AtomTypes() []string {
	return c.atypes
}

//
func (c *CMapType) SetValues(v []float64) {
	c.values = v
}

func (c *CMapType) Values() []float64 {
	return c.values
}

//
