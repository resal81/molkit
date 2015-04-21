package blocks

import (
	"github.com/resal81/molkit/utils"
)

var (
	fragmentHash = utils.NewComponentHash()
)

type Fragment struct {
	id int64

	name   string
	serial int64

	atoms []*Atom
	pol   *Polymer
}

func NewFragment() *Fragment {
	frag := &Fragment{}
	frag.id = fragmentHash.Add(frag)
	return frag
}

func (f *Fragment) Delete() {
	f.Polymer().deleteFragment(f)
}

func (f *Fragment) deleteAtom(a1 *Atom) {
	f.Polymer().atomDeleted(a1)
	for i, a2 := range f.Atoms() {
		if a1.Id() == a2.Id() {
			f.atoms = append(f.atoms[:i], f.atoms[i+1:]...)
			return
		}
	}
}

func (f *Fragment) Id() int64 {
	return f.id
}

func (f *Fragment) SetName(name string) {
	f.name = name
}

func (f *Fragment) Name() string {
	return f.name
}

func (f *Fragment) SetSerial(ser int64) {
	f.serial = ser
}

func (f *Fragment) Serial() int64 {
	return f.serial
}

func (f *Fragment) AddAtom(a *Atom) {
	a.setFragment(f)
	f.atoms = append(f.atoms, a)
}

func (f *Fragment) Atoms() []*Atom {
	return f.atoms
}

func (f *Fragment) setPolymer(p *Polymer) {
	f.pol = p
}

func (f *Fragment) Polymer() *Polymer {
	return f.pol
}
