package mol

import (
	"strings"
)

/**********************************************\
    Bond
\**********************************************/

type Bond struct {
	bondOrder int
}

/**********************************************\
    Element
\**********************************************/

type Element struct {
	elementName   string
	elementNumber int
}

func NewElement(name string, number int) *Element {
	return &Element{
		elementName:   name,
		elementNumber: number,
	}
}

func (el *Element) Name() {
	return el.elementName
}

func (el *Element) Number() {
	return el.elementNumber
}

/**********************************************\
    AtomKind
\**********************************************/

type AtomKind struct {
	kindName string
	radius   float64
	charge   float64
}

/**********************************************\
    Atom
\**********************************************/

type Atom struct {
	atomName   string
	atomSerial int
	atomKind   *AtomKind
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

func (at *Atom) AtomKind() *AtomKind {
	return at.atomKind
}

func (at *Atom) SetAtomKind(kind *AtomKind) {
	at.atomKind = kind
}

func (at *Atom) Element() *Element {
	return at.element
}

func (at *Atom) SetElement(element *Element) {
	at.element = element
}

/**********************************************\
    FragmentKind
\**********************************************/

const (
	FragmentAminoAcid = iota
	FragmentNucleic
	FragmentLipid
	FragmentWater
	FragmentIon
)

type FragmentKind struct {
	longName  string // LYSINE
	shortName string // LYS
	codeName  string // K -> single letter abbr
}

func NewFragmentKind(longName, shortName, codeName string) *FragmentKind {
	return &FragmentKind{
		longName:  strings.ToUpper(longName),
		shortName: strings.ToUpper(shortName),
		codeName:  strings.ToUpper(codeName),
	}
}

func (fk *FragmentKind) LongName() string {
	return fk.longName
}

func (fk *FragmentKind) ShortName() string {
	return fk.shortName
}

func (fk *FragmentKind) CodeName() string {
	return fk.codeName
}

/**********************************************\
    Fragment
\**********************************************/

type Fragment struct {
	fragSerial int
	fragKind   *FragmentKind
	atoms      []*Atom
}

func NewFragment(serial int) *Fragment {
	return &Fragment{
		fragSerial: serial,
		atoms:      make([]*Atom),
	}
}

/**********************************************\
    Chain
\**********************************************/

type Chain struct {
	chainName string
	fragments []*Fragment
}

func NewChain(name string) *Chain {
	return &Chain{
		chainName: name,
		fragments: make([]*Fragment),
	}
}

/**********************************************\
    Complex
\**********************************************/

type Complex struct {
	complexName string
	chains      []*Chain
}

func NewComplex(name string) *Complex {
	return &Complex{
		complexName: name,
		chains:      make([]*Chain),
	}
}
