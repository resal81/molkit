package blocks

import (
	"github.com/resal81/molkit/utils"
)

var (
	polymerHash = utils.NewComponentHash()
)

type Polymer struct {
	id   int64
	name string

	fragments []*Fragment
	system    *System

	bonds     []*Bond
	angles    []*Angle
	dihedrals []*Dihedral
	impropers []*Dihedral
}

func NewPolymer() *Polymer {
	pol := &Polymer{}
	pol.id = polymerHash.Add(pol)
	return pol
}

func (p *Polymer) Delete() {
	p.System().deletePolymer(p)
}

func (p *Polymer) deleteFragment(f1 *Fragment) {
	for i, f2 := range p.Fragments() {
		if f1.Id() == f2.Id() {
			for _, a1 := range f1.Atoms() {
				p.atomDeleted(a1)
			}
			p.fragments = append(p.fragments[:i], p.fragments[i+1:]...)
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

// Name
func (p *Polymer) SetName(name string) {
	p.name = name
}

func (p *Polymer) Name() string {
	return p.name
}

// Atoms
func (p *Polymer) Atoms() []*Atom {
	n := 0
	for _, frag := range p.Fragments() {
		n += len(frag.Atoms())
	}

	out := make([]*Atom, n)
	var i int = 0
	for _, frag := range p.Fragments() {
		for _, a := range frag.Atoms() {
			out[i] = a
			i++
		}
	}
	return out
}

// Fragment
func (p *Polymer) AddFragment(f *Fragment) {
	f.setPolymer(p)
	p.fragments = append(p.fragments, f)
}

func (p *Polymer) Fragments() []*Fragment {
	return p.fragments
}

// System
func (p *Polymer) setSystem(s *System) {
	p.system = s
}

func (p *Polymer) System() *System {
	return p.system
}
