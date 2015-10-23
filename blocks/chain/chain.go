package chain

import (
	"github.com/resal81/molkit/blocks/atom"
	"github.com/resal81/molkit/blocks/residue"
)

// ***********************************************************************//
// Chain struct
// ***********************************************************************//

type Chain struct {
	name     string
	residues []*residue.Residue
}

func NewChain(name string) *Chain {
	return &Chain{
		name:     name,
		residues: make([]*residue.Residue, 0),
	}
}

func (ch *Chain) AddResidue(rs *residue.Residue) {
	ch.residues = append(ch.residues, rs)
}

func (ch *Chain) Residues() []*residue.Residue {
	return ch.residues
}

func (ch *Chain) Name() string {
	return ch.name
}

func (ch *Chain) Atoms() []*atom.Atom {
	var atoms = []*atom.Atom{}
	for _, r := range ch.Residues() {
		atoms = append(atoms, r.Atoms()...)
	}
	return atoms
}
