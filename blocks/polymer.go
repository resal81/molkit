package blocks

import (
	"github.com/resal81/molkit/utils"
)

var (
	polymerHash = utils.NewComponentHash()
)

type Polymer struct {
	id   int64
	Name string

	Fragments []*Fragment
	System    *System

	Bonds     []*Bond
	Angles    []*Angle
	Dihedrals []*Dihedral
	Impropers []*Dihedral
}

func NewPolymer() *Polymer {
	pol := &Polymer{}
	pol.id = polymerHash.Add(pol)
	return pol
}

func (p *Polymer) Delete() {
	p.System.deletePolymer(p)
}

func (p *Polymer) deleteFragment(f1 *Fragment) {
	for i, f2 := range p.Fragments {
		if f1.Id() == f2.Id() {
			for _, a1 := range f1.Atoms {
				p.atomDeleted(a1)
			}
			p.Fragments = append(p.Fragments[:i], p.Fragments[i+1:]...)
			return
		}
	}
}

func (p *Polymer) atomDeleted(a *Atom) {
	// update bonds, angles, dihedrals and impropers
	panic("updates for atom deletetion hasn't been implemented yet")
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
