package blocks

import (
	"sync/atomic"
)

var (
	bondid_counter     int64 = 0
	angleid_counter    int64 = 0
	dihedralid_counter int64 = 0
)

type Bond struct {
	id int64

	atom1 *Atom
	atom2 *Atom
}

func NewBond(atom1, atom2 *Atom) *Bond {
	id := atomic.AddInt64(&bondid_counter, 1)
	return &Bond{
		id:    id,
		atom1: atom1,
		atom2: atom2,
	}
}

func (b *Bond) Id() int64 {
	return b.id
}

type Angle struct {
	id int64

	atom1, atom2, atom3 *Atom
}

func NewAngle(atom1, atom2, atom3 *Atom) *Angle {
	id := atomic.AddInt64(&angleid_counter, 1)
	return &Angle{
		id:    id,
		atom1: atom1,
		atom2: atom2,
		atom3: atom3,
	}
}

func (a *Angle) Id() int64 {
	return a.id
}

type Dihedral struct {
	id int64

	atom1, atom2, atom3, atom4 *Atom
}

func NewDihedral(atom1, atom2, atom3, atom4 *Atom) *Dihedral {
	id := atomic.AddInt64(&dihedralid_counter, 1)
	return &Dihedral{
		id:    id,
		atom1: atom1,
		atom2: atom2,
		atom3: atom3,
		atom4: atom4,
	}
}

func (d *Dihedral) Id() int64 {
	return d.id
}
