package blocks

import (
	"github.com/resal81/molkit/utils"
)

var (
	atomHash = utils.NewComponentHash()
)

type ATSetting int64

const (
	AT_TYPE_CHM ATSetting = 1 << iota
	AT_TYPE_GMX
	AT_HAS_PRATONS_SET
	AT_HAS_MASS_SET
	AT_HAS_LJ_DIST_SET
	AT_HAS_LJ_ENERGY_SET
	AT_HAS_LJ_DIST14_SET
	AT_HAS_LJ_ENERGY14_SET
	AT_HAS_CHARGE_SET
	AT_HAS_PAR_CHARGE_SET
	AT_HAS_RADIUS_SET
)

type AtomType struct {
	Label      string
	Protons    int
	Mass       float64
	LJDist     float64
	LJEnergy   float64
	LJDist14   float64
	LJEnergy14 float64
	Charge     int
	ParCharge  float64
	Radius     float64
	Setting    ATSetting
}

type Atom struct {
	id        int64
	Name      string
	Serial    int
	BFactor   float64
	Occupancy float64
	AltLoc    string
	IsHetero  bool
	Coords    [][3]float64
	Type      *AtomType
	Fragment  *Fragment
}

func NewAtom() *Atom {
	at := &Atom{}
	id := atomHash.Add(at)
	at.id = id
	return at
}

func (a *Atom) Id() int64 {
	return a.id
}

func (a *Atom) Delete() {
	a.Fragment.deleteAtom(a)
	atomHash.Drop(a.Id())
}
