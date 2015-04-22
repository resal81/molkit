package blocks

import (
	"fmt"
	"github.com/resal81/molkit/utils"
)

var (
	atomHash     = utils.NewComponentHash()
	fragmentHash = utils.NewComponentHash()
	systemHash   = utils.NewComponentHash()
)

/*
	Atom
*/

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

/*
	Link
*/

type Link struct {
	Fragment1   *Fragment
	Fragment2   *Fragment
	Connections [][2]int64
}

/*
	System
*/

type System struct {
	id        int64
	Atoms     []*Atom
	Bonds     []*Bond
	Angles    []*Angle
	Dihedrals []*Dihedral
	Impropers []*Improper
	Links     []*Link
}

func NewSystem() *System {
	sys := &System{}
	sys.id = systemHash.Add(sys)
	return sys
}

func (s *System) Id() int64 {
	return s.id
}

func (s *System) String() string {
	return fmt.Sprintf("<system with %d atoms>", len(s.Atoms))
}
