package mol

const (
	FragmentAminoAcid = iota
	FragmentNucleic
	FragmentLipid
	FragmentWater
	FragmentIon
)

type FragmentType struct {
	longName  string // LYSINE
	shortName string // LYS
	codeName  string // K -> single letter abbr
}

func NewFragmentType(longName, shortName, codeName string) *FragmentType {
	return &FragmentType{
		longName:  strings.ToUpper(longName),
		shortName: strings.ToUpper(shortName),
		codeName:  strings.ToUpper(codeName),
	}
}

func (fk *FragmentType) LongName() string {
	return fk.longName
}

func (fk *FragmentType) ShortName() string {
	return fk.shortName
}

func (fk *FragmentType) CodeName() string {
	return fk.codeName
}

type Fragment struct {
	fragSerial int
	fragKind   *FragmentType
	Atoms      []*Atom
}

func NewFragment(serial int) *Fragment {
	return &Fragment{
		fragSerial: serial,
		Atoms:      make([]*Atom),
	}
}
