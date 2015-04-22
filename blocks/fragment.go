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
	Name      string
	Serial    int64
	Atoms     []*Atom
	Bonds     []*Bond
	Angles    []*Angle
	Dihedrals []*Dihedral
	Impropers []*Improper
}

func NewFragment() *Fragment {
	frag := &Fragment{}
	frag.id = fragmentHash.Add(frag)
	return frag
}

func (f *Fragment) Id() int64 {
	return f.id
}
