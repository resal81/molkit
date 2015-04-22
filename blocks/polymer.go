package blocks

import (
	"github.com/resal81/molkit/utils"
)

var (
	polymerHash = utils.NewComponentHash()
)

type Polymer struct {
	id        int64
	Name      string
	Bonds     []*Bond
	Angles    []*Angle
	Dihedrals []*Dihedral
	Impropers []*Dihedral
	Fragments []*Fragment
	Links     []*Link
	System    *System
}

func NewPolymer() *Polymer {
	pol := &Polymer{}
	pol.id = polymerHash.Add(pol)
	return pol
}

func (s *Polymer) Id() int64 {
	return s.id
}

// Atoms
func (p *Polymer) Atoms() []*Atom {
	n := 0
	for _, frag := range p.Fragments {
		n += len(frag.Atoms)
	}

	out := make([]*Atom, n)
	var i int = 0
	for _, frag := range p.Fragments {
		for _, a := range frag.Atoms {
			out[i] = a
			i++
		}
	}
	return out
}
