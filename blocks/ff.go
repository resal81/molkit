package blocks

type FFSetting int64

const (
	FF_CHM FFSetting = 1 << iota
	FF_GMX
	FF_AMB
)

type ForceField struct {
	AtomTypes      []*AtomType
	BondTypes      []*BondType
	AngleTypes     []*AngleType
	DihedralTypes  []*DihedralType
	ImproperTypes  []*ImproperType
	NonBondedTypes []*PairType
	OneFourTypes   []*PairType
	Setting        FFSetting

	GMXNbFunc   int
	GMXCombRule int
	GMXGenPairs bool
	GMXFudgeLJ  float64
	GMXFudgeQQ  float64
}

func NewForceField() *ForceField {
	return &ForceField{}
}
