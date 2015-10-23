package atom

import (
	"github.com/resal81/molkit/blocks/atomtype"
)

// ***********************************************************************//
// Atom struct
// ***********************************************************************//

type Atom struct {
	name   string
	serial int64
	coords [][3]float64

	pdb struct {
		occupancy float64
		beta      float64
		altloc    string
	}

	atomType *atomtype.AtomType
}

func NewAtom(name string, serial int64) *Atom {
	if name == "" {
		panic("Atom: name cannot be empty")
	}

	return &Atom{
		name:   name,
		serial: serial,
		coords: make([][3]float64, 0),
	}
}

func (a *Atom) Name() string {
	return a.name
}

func (a *Atom) Serial() int64 {
	return a.serial
}

func (a *Atom) AddCoord(c [3]float64) {
	a.coords = append(a.coords, c)
}

func (a *Atom) Coords() [][3]float64 {
	return a.coords
}

func (a *Atom) CoordsAtFrame(i int) [3]float64 {
	return a.coords[i]
}

func (a *Atom) PdbOccupancy() float64 {
	return a.pdb.occupancy
}

func (a *Atom) SetPdbOccupancy(occ float64) {
	a.pdb.occupancy = occ
}

func (a *Atom) PdbBeta() float64 {
	return a.pdb.beta
}

func (a *Atom) SetPdbBeta(beta float64) {
	a.pdb.beta = beta
}

func (a *Atom) PdbAltLoc() string {
	return a.pdb.altloc
}

func (a *Atom) SetPdbAltLoc(altloc string) {
	a.pdb.altloc = altloc
}
