package mol

// ==============================================

type AtomType struct {
	name   string
	radius float64
	charge float64
}

// ==============================================

type Atom struct {
	atomName   string
	atomSerial int
	atomType   *AtomType
	element    *Element

	pdb struct {
		occupancy float64
		bfactor   float64
		altloc    string
		isHetero  bool
	}
}

func NewAtom(name string, serial int) *Atom {
	return &Atom{
		atomName:   name,
		atomSerial: serial,
	}
}

func (at *Atom) Name() string {
	return at.atomName
}

func (at *Atom) Serial() int {
	return at.atomSerial
}

func (at *Atom) AtomType() *AtomType {
	return at.atomType
}

func (at *Atom) SetAtomType(kind *AtomType) {
	at.atomType = kind
}

func (at *Atom) Element() *Element {
	return at.element
}

func (at *Atom) SetElement(element *Element) {
	at.element = element
}
