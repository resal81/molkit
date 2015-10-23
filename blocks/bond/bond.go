package bond

import (
	"github.com/resal81/molkit/atom"
)

type Bond struct {
	a1 *atom.Atom
	a2 *atom.Atom
}

func NewBond(a1, a2 *atom.Atom) *Bond {
	return &Bond{
		a1: a1,
		a2: a2,
	}
}

func (bn *Bond) Atom1() *atom.Atom {
	return bn.a1
}

func (bn *Bond) Atom2() *atom.Atom {
	return bn.a2
}
