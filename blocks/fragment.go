package blocks

import (
	"github.com/resal81/molkit/utils"
)

var (
	fragmentHash = utils.NewComponentHash()
)

/*
	Fragment
*/

type Fragment struct {
	id        int64
	name      string
	serial    int64
	atoms     []*Atom
	bonds     []*Bond
	angles    []*Angle
	dihedrals []*Dihedral
	impropers []*Improper
}

/* new fragment */

func NewFragment(name string) *Fragment {
	frag := &Fragment{
		name: name,
	}
	frag.id = fragmentHash.Add(frag)
	return frag
}

/* id */

func (f *Fragment) Id() int64 {
	return f.id
}

/* name */

func (f *Fragment) Name() string {
	return f.name
}

/* serial */

func (f *Fragment) SetSerial(s int64) {
	f.serial = s
}

func (f *Fragment) Serial() int64 {
	return f.serial
}

/* atoms */

func (s *Fragment) AddAtom(a *Atom) {
	s.atoms = append(s.atoms, a)
}

func (s *Fragment) Atoms() []*Atom {
	return s.atoms
}

/* bonds */

func (s *Fragment) AddBond(b *Bond) {
	s.bonds = append(s.bonds, b)
}

func (s *Fragment) Bonds() []*Bond {
	return s.bonds
}

/* angles */

func (s *Fragment) AddAngle(a *Angle) {
	s.angles = append(s.angles, a)
}

func (s *Fragment) Angles() []*Angle {
	return s.angles
}

/* dihedrals */

func (s *Fragment) AddDihedral(d *Dihedral) {
	s.dihedrals = append(s.dihedrals, d)
}

func (s *Fragment) Dihedrals() []*Dihedral {
	return s.dihedrals
}

/* impropers */

func (s *Fragment) AddImproper(b *Improper) {
	s.impropers = append(s.impropers, b)
}

func (s *Fragment) Impropers() []*Improper {
	return s.impropers
}
