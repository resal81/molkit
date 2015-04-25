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
	id         int64
	name       string
	serial     int64
	atoms      []*Atom
	bonds      []*Bond
	angles     []*Angle
	dihedrals  []*Dihedral
	impropers  []*Improper
	cmaps      []*CMap
	nextLinker *Linker
	prevLinker *Linker
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

func (f *Fragment) AddAtom(a *Atom) {
	f.atoms = append(f.atoms, a)
}

func (f *Fragment) Atoms() []*Atom {
	return f.atoms
}

/* bonds */

func (f *Fragment) AddBond(b *Bond) {
	f.bonds = append(f.bonds, b)
}

func (f *Fragment) Bonds() []*Bond {
	return f.bonds
}

/* angles */

func (f *Fragment) AddAngle(a *Angle) {
	f.angles = append(f.angles, a)
}

func (f *Fragment) Angles() []*Angle {
	return f.angles
}

/* dihedrals */

func (f *Fragment) AddDihedral(d *Dihedral) {
	f.dihedrals = append(f.dihedrals, d)
}

func (f *Fragment) Dihedrals() []*Dihedral {
	return f.dihedrals
}

/* impropers */

func (f *Fragment) AddImproper(b *Improper) {
	f.impropers = append(f.impropers, b)
}

func (f *Fragment) Impropers() []*Improper {
	return f.impropers
}

/* cmaps */

func (f *Fragment) AddCMap(c *CMap) {
	f.cmaps = append(f.cmaps, c)
}

func (f *Fragment) CMaps() []*CMap {
	return f.cmaps
}

/* linkers */

func (f *Fragment) SetLinkerNext(l *Linker) {
	f.nextLinker = l
}

func (f *Fragment) LinkerNext() *Linker {
	return f.nextLinker
}

func (f *Fragment) HasLinkerNext() bool {
	return f.nextLinker != nil
}

func (f *Fragment) SetLinkerPrev(l *Linker) {
	f.prevLinker = l
}

func (f *Fragment) LinkerPrev() *Linker {
	return f.prevLinker
}

func (f *Fragment) HasLinkerPrev() bool {
	return f.prevLinker != nil
}
