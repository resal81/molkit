package blocks

import (
	"sync/atomic"
)

var (
	bondid_counter     int64 = 0
	angleid_counter    int64 = 0
	dihedralid_counter int64 = 0
)

// --------------------------------------------------------

type Bond struct {
	id int64

	atom1    *Atom
	atom2    *Atom
	bondtype *BondType
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

func (b *Bond) SetType(bt *BondType) {
	b.bondtype = bt
}

func (b *Bond) Type() *BondType {
	return b.bondtype
}

// --------------------------------------------------------

type Angle struct {
	id int64

	atom1   *Atom
	atom2   *Atom
	atom3   *Atom
	angtype *AngleType
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

func (a *Angle) SetType(at *AngleType) {
	a.angtype = at
}

func (a *Angle) Type() *AngleType {
	return a.angtype
}

// --------------------------------------------------------

type Dihedral struct {
	id int64

	atom1   *Atom
	atom2   *Atom
	atom3   *Atom
	atom4   *Atom
	dihtype *DihedralType
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

func (d *Dihedral) SetType(dt *DihedralType) {
	d.dihtype = dt
}

func (d *Dihedral) Type() *DihedralType {
	return d.dihtype
}

// --------------------------------------------------------

type Constraint struct {
}

// --------------------------------------------------------

type Pair struct {
}

// --------------------------------------------------------

type Exclusion struct {
}

// --------------------------------------------------------

type Settle struct {
	d_OH float32
	d_HH float32
}

// --------------------------------------------------------

type PositionRestraint struct {
}

// --------------------------------------------------------

type DistanceRestraint struct {
}

// --------------------------------------------------------

type AngleRestraint struct {
}

// --------------------------------------------------------

type DihedralRestraint struct {
}

// --------------------------------------------------------

type OrientationRestraint struct {
}
