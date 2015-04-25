package blocks

/**********************************************************
* PairType
**********************************************************/

type PTSetting int64

const (
	PT_NULL PTSetting = 1 << iota
	PT_TYPE_CHM_1
	PT_TYPE_GMX_1
	PT_HAS_LJ_DISTANCE_SET
	PT_HAS_LJ_ENERGY_SET
	PT_HAS_LJ_DISTANCE_14_SET
	PT_HAS_LJ_ENERGY_14_SET
)

// PairType is used to describe custom parameters for non-bonding interactions
type PairType struct {
	aType1     string
	aType2     string
	ljDist     float64
	ljEnergy   float64
	ljDist14   float64
	ljEnergy14 float64
	setting    PTSetting
}

/* new pairtype */

func NewPairType(at1, at2 string, t PTSetting) *PairType {
	return &PairType{
		aType1:  at1,
		aType2:  at2,
		setting: t,
	}
}

/* types */

func (pt *PairType) AType1() string {
	return pt.aType1
}

func (pt *PairType) AType2() string {
	return pt.aType2
}

/* ljDist */

func (pt *PairType) SetLJDist(v float64) {
	pt.setting |= PT_HAS_LJ_DISTANCE_SET
	pt.ljDist = v
}
func (pt *PairType) HasLJDistSet() bool {
	return pt.setting&PT_HAS_LJ_DISTANCE_SET != 0
}
func (pt *PairType) LJDist() float64 {
	return pt.ljDist
}

/* ljDist14 */

func (pt *PairType) SetLJDist14(v float64) {
	pt.setting |= PT_HAS_LJ_DISTANCE_14_SET
	pt.ljDist14 = v
}
func (pt *PairType) HasLJDist14Set() bool {
	return pt.setting&PT_HAS_LJ_DISTANCE_14_SET != 0
}
func (pt *PairType) LJDist14() float64 {
	return pt.ljDist14
}

/* ljEnergy */

func (pt *PairType) SetLJEnergy(v float64) {
	pt.setting |= PT_HAS_LJ_ENERGY_SET
	pt.ljEnergy = v
}
func (pt *PairType) HasLJEnergySet() bool {
	return pt.setting&PT_HAS_LJ_ENERGY_SET != 0
}
func (pt *PairType) LJEnergy() float64 {
	return pt.ljEnergy
}

/* ljEnergy14 */

func (pt *PairType) SetLJEnergy14(v float64) {
	pt.setting |= PT_HAS_LJ_ENERGY_14_SET
	pt.ljEnergy14 = v
}
func (pt *PairType) HasLJEnergy14Set() bool {
	return pt.setting&PT_HAS_LJ_ENERGY_14_SET != 0
}
func (pt *PairType) LJEnergy14() float64 {
	return pt.ljEnergy14
}

/* setting */

func (pt *PairType) Setting() PTSetting {
	return pt.setting
}

/**********************************************************
* Pair
**********************************************************/

type Pair struct {
	atom1 *Atom
	atom2 *Atom
	tipe  *PairType
}

/* new pair */

func NewPair(a1, a2 *Atom) *Pair {
	return &Pair{
		atom1: a1,
		atom2: a2,
	}
}

/* atoms */

func (p *Pair) Atom1() *Atom {
	return p.atom1
}

func (p *Pair) Atom2() *Atom {
	return p.atom2
}

/* type */

func (p *Pair) SetType(pt *PairType) {
	p.tipe = pt
}

func (p *Pair) Type() *PairType {
	return p.tipe
}
