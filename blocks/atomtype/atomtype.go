package atomtype

import (
	"github.com/resal81/molkit/blocks/element"
)

type AtomType struct {
	name string

	pqr struct {
		charge float64
		radius float64
	}
	vdw struct {
		eps, eps14, sig, sig14 float64
	}

	element *element.Element
}

func NewAtomType(name string) *AtomType {
	return &AtomType{
		name: name,
	}
}

func (at *AtomType) Name() string {
	return at.name
}

func (at *AtomType) PqrCharge() float64 {
	return at.pqr.charge
}

func (at *AtomType) SetPqrCharge(charge float64) {
	at.pqr.charge = charge
}

func (at *AtomType) PqrRadius() float64 {
	return at.pqr.radius
}

func (at *AtomType) SetPqrRadius(radius float64) {
	at.pqr.radius = radius
}
