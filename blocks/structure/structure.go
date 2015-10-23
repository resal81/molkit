package structure

import (
	"github.com/resal81/molkit/blocks/chain"
)

/*
An Structure is made of one or more chains.
*/
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
