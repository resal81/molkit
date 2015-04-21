package blocks

import (
	"github.com/resal81/molkit/utils"
)

var (
	bondHash     = utils.NewComponentHash()
	angleHash    = utils.NewComponentHash()
	dihedralHash = utils.NewComponentHash()
)

// --------------------------------------------------------

type Bond struct {
	id int64

	atom1 *Atom
	atom2 *Atom
}

func NewBond(atom1, atom2 *Atom) *Bond {
	bnd := &Bond{
		atom1: atom1,
		atom2: atom2,
	}
	bnd.id = bondHash.Add(bnd)
	return bnd
}

func (b *Bond) Id() int64 {
	return b.id
}

// --------------------------------------------------------

type Angle struct {
	id int64

	atom1 *Atom
	atom2 *Atom
	atom3 *Atom
}

func NewAngle(atom1, atom2, atom3 *Atom) *Angle {
	ang := &Angle{
		atom1: atom1,
		atom2: atom2,
		atom3: atom3,
	}
	ang.id = angleHash.Add(ang)
	return ang
}

func (a *Angle) Id() int64 {
	return a.id
}

// --------------------------------------------------------

type Dihedral struct {
	id int64

	atom1 *Atom
	atom2 *Atom
	atom3 *Atom
	atom4 *Atom
}

func NewDihedral(atom1, atom2, atom3, atom4 *Atom) *Dihedral {
	dih := &Dihedral{
		atom1: atom1,
		atom2: atom2,
		atom3: atom3,
		atom4: atom4,
	}

	dih.id = dihedralHash.Add(dih)
	return dih
}

func (d *Dihedral) Id() int64 {
	return d.id
}
