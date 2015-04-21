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

type BTSetting int64

const (
	BT_HARM_CONST_SET BTSetting = 1 << iota
	BT_HARM_DIST_SET
	BT_TYPE_GMX_1
	BT_TYPE_CHM_1
	BT_ORDER_SINGLE
	BT_ORDER_DOUBLE
	BT_ORDER_TRIPLE
)

type BondType struct {
	AType1    string
	AType2    string
	HarmConst float64
	HarmDist  float64
	Setting   BTSetting
}

type Bond struct {
	id int64

	atom1 *Atom
	atom2 *Atom
	order int
	btype *BondType
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
