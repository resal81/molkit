package structure

import (
	"github.com/resal81/molkit/blocks/atom"
	"github.com/resal81/molkit/blocks/chain"
	"github.com/resal81/molkit/blocks/residue"
)

// ***********************************************************************//
// Structure struct
// ***********************************************************************//

type Structure struct {
	chains []*chain.Chain
}

// NewStructure returns a new, empty Structure
func NewStructure() *Structure {
	return &Structure{
		chains: make([]*chain.Chain, 0),
	}
}

// Chains return the chain slice
func (st *Structure) Chains() []*chain.Chain {
	return st.chains
}

// AddChain adds a chain to the structure
func (st *Structure) AddChain(chain *chain.Chain) {
	st.chains = append(st.chains, chain)
}

// Atoms returns an up-to-date list of atoms for the structure
func (st *Structure) Atoms() []*atom.Atom {
	var atoms = []*atom.Atom{}
	for _, ch := range st.Chains() {
		atoms = append(atoms, ch.Atoms()...)
	}
	return atoms
}

// Residues returns an up-to-date list of residues for the structure
func (st *Structure) Residues() []*residue.Residue {
	var residues = []*residue.Residue{}
	for _, ch := range st.Chains() {
		residues = append(residues, ch.Residues()...)
	}
	return residues
}
