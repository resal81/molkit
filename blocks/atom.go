package blocks

import (
	"sync"
	"sync/atomic"
)

var (
	atomid_counter int64 = 0
)

type Atom struct {
	id     int64
	name   string
	serial int64

	mass          float32
	protons       int8
	valence       int8
	formal_charge int8

	pdbprops struct {
		bfactor   float32
		occupancy float32
		altloc    string
		isHetero  bool
	}

	pqrprops struct {
		charge float32
		radius float32
	}

	ffprops struct {
		atomtype string
	}

	coords [][3]float64

	mu *sync.Mutex

	frag *Fragment
}

func NewAtom() *Atom {
	id := atomic.AddInt64(&atomid_counter, 1)
	return &Atom{
		id: id,
		mu: new(sync.Mutex),
	}
}

func (a *Atom) Delete() {
	a.Fragment().deleteAtom(a)
}

func (a *Atom) Id() int64 {
	return a.id
}

func (a *Atom) SetName(name string) {
	a.name = name
}

func (a *Atom) Name() string {
	return a.name
}

func (a *Atom) SetSerial(ser int64) {
	a.serial = ser
}

func (a *Atom) Serial() int64 {
	return a.serial
}

func (a *Atom) SetAtomicNumber(n int8) {
	a.protons = n
}

func (a *Atom) AtomicNumber() int8 {
	return a.protons
}

func (a *Atom) setFragment(f *Fragment) {
	a.frag = f
}

func (a *Atom) Fragment() *Fragment {
	return a.frag
}

func (a *Atom) AddCoord(x, y, z float64) {
	a.coords = append(a.coords, [3]float64{x, y, z})
}

func (a *Atom) Coords() [][3]float64 {
	return a.coords
}

func (a *Atom) SetPropBFactor(val float32) {
	a.pdbprops.bfactor = val
}

func (a *Atom) PropBFactor() float32 {
	return a.pdbprops.bfactor
}

func (a *Atom) SetPropOccupancy(val float32) {
	a.pdbprops.occupancy = val
}

func (a *Atom) PropOccupancy() float32 {
	return a.pdbprops.occupancy
}

func (a *Atom) SetPropAltloc(flag string) {
	a.pdbprops.altloc = flag
}

func (a *Atom) PropAltloc() string {
	return a.pdbprops.altloc
}

func (a *Atom) SetPropIsHetero(flag bool) {
	a.pdbprops.isHetero = flag
}

func (a *Atom) PropIsHetero() bool {
	return a.pdbprops.isHetero
}
