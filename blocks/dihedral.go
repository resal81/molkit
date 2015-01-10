package blocks

import (
	"sync/atomic"
)

var (
	dihedralid_counter int64 = 0
)

type Dihedral struct {
	id int64

	atom1, atom2, atom3, atom4 *Atom
}

func NewDihedral() *Dihedral {
	id := atomic.AddInt64(&dihedralid_counter, 1)
	return &Dihedral{
		id: id,
	}
}

func (d *Dihedral) Id() int64 {
	return d.id
}
