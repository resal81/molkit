package residue

import (
	"github.com/resal81/molkit/blocks/atom"
)

// ***********************************************************************//
// Residue struct
// ***********************************************************************//

type Residue struct {
	name   string
	serial int64
	atoms  []*atom.Atom
}

func NewResidue(name string, serial int64) *Residue {
	return &Residue{
		name:   name,
		serial: serial,
		atoms:  make([]*atom.Atom, 0),
	}
}

func (rs *Residue) AddAtom(at *atom.Atom) {
	rs.atoms = append(rs.atoms, at)
}

func (rs *Residue) Atoms() []*atom.Atom {
	return rs.atoms
}

func (rs *Residue) Name() string {
	return rs.name
}

func (rs *Residue) Serial() int64 {
	return rs.serial
}
