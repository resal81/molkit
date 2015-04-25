package blocks

import (
	"fmt"
	"github.com/resal81/molkit/utils"
)

var (
	systemHash = utils.NewComponentHash()
)

/**********************************************************
* System
**********************************************************/

type System struct {
	id        int64
	atoms     []*Atom
	bonds     []*Bond
	angles    []*Angle
	dihedrals []*Dihedral
	impropers []*Improper
	fragments []*Fragment
}

/* new system */

func NewSystem() *System {
	sys := &System{}
	sys.id = systemHash.Add(sys)
	return sys
}

/* id */

func (s *System) Id() int64 {
	return s.id
}

/* string */

func (s *System) String() string {
	return fmt.Sprintf("<system with %d atoms>", len(s.Atoms()))
}

/* atoms */

func (s *System) AddAtom(a *Atom) {
	s.atoms = append(s.atoms, a)
}

func (s *System) Atoms() []*Atom {
	return s.atoms
}

/* bonds */

func (s *System) AddBond(b *Bond) {
	s.bonds = append(s.bonds, b)
}

func (s *System) Bonds() []*Bond {
	return s.bonds
}

/* angles */

func (s *System) AddAngle(a *Angle) {
	s.angles = append(s.angles, a)
}

func (s *System) Angles() []*Angle {
	return s.angles
}

/* dihedrals */

func (s *System) AddDihedral(d *Dihedral) {
	s.dihedrals = append(s.dihedrals, d)
}

func (s *System) Dihedrals() []*Dihedral {
	return s.dihedrals
}

/* impropers */

func (s *System) AddImproper(b *Improper) {
	s.impropers = append(s.impropers, b)
}

func (s *System) Impropers() []*Improper {
	return s.impropers
}

/* fragments */

func (s *System) AddFragment(f *Fragment) {
	s.fragments = append(s.fragments, f)
}

func (s *System) Fragments() []*Fragment {
	return s.fragments
}
