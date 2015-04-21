package blocks

import (
	"github.com/resal81/molkit/utils"
)

var (
	fragmentHash = utils.NewComponentHash()
)

type Fragment struct {
	id      int64
	Name    string
	Serial  int64
	Atoms   []*Atom
	Polymer *Polymer
}

func NewFragment() *Fragment {
	frag := &Fragment{}
	frag.id = fragmentHash.Add(frag)
	return frag
}

func (f *Fragment) Delete() {
	f.Polymer.deleteFragment(f)
}

func (f *Fragment) deleteAtom(a1 *Atom) {

	// ask Polymer to update bonds, ...
	f.Polymer.atomDeleted(a1)

	for i, a2 := range f.Atoms {
		if a1.Id() == a2.Id() {
			f.Atoms = append(f.Atoms[:i], f.Atoms[i+1:]...)
			return
		}
	}
}

func (f *Fragment) Id() int64 {
	return f.id
}
