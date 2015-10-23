package blocks

import (
	"testing"
)

func TestAtomType(t *testing.T) {
	var ats = []struct {
		label      string
		proton     int
		mass       float64
		charge     int
		par_charge float64
		radius     float64
		lje        float64
		lje14      float64
		ljd        float64
		ljd14      float64
	}{
		{"CA", 6, 12, 1, 0.15, 1.8, -1e10, -2e10, -3e10, -4e10},
	}

	for _, el := range ats {
		at := NewAtomType(el.label, AT_TYPE_GMX_1)
		if lb := at.Label(); lb != el.label {
			t.Errorf("wrong label => %q, wanted %q", lb, el.label)
		}
		if v := at.Setting(); v&AT_TYPE_GMX_1 == 0 {
			t.Errorf("AT_TYPE is wrong => %q, expected %q", v, AT_TYPE_GMX_1)
		}

		//
		var h string

		// protons
		if at.HasProtonsSet() {
			t.Errorf("%s was already set", h)
		}
		at.SetProtons(el.proton)
		if v := at.Protons(); v != el.proton {
			t.Errorf("wrong %s => %q, wanted %q", h, v, el.proton)
		}
		if !at.HasProtonsSet() {
			t.Errorf("%s was not set", h)
		}

		// mass
		if at.HasMassSet() {
			t.Errorf("%s was already set", h)
		}
		at.SetMass(el.mass)
		if v := at.Mass(); v != el.mass {
			t.Errorf("wrong %s => %q, wanted %q", h, v, el.mass)
		}
		if !at.HasMassSet() {
			t.Errorf("%s was not set", h)
		}

		// charge
		if at.HasChargeSet() {
			t.Errorf("%s was already set", h)
		}
		at.SetCharge(el.charge)
		if v := at.Charge(); v != el.charge {
			t.Errorf("wrong %s => %q, wanted %q", h, v, el.charge)
		}
		if !at.HasMassSet() {
			t.Errorf("%s was not set", h)
		}

		// par charge
		if at.HasPartialChargeSet() {
			t.Errorf("%s was already set", h)
		}
		at.SetPartialCharge(el.par_charge)
		if v := at.PartialCharge(); v != el.par_charge {
			t.Errorf("wrong %s => %q, wanted %q", h, v, el.par_charge)
		}
		if !at.HasPartialChargeSet() {
			t.Errorf("%s was not set", h)
		}

		// radius
		if at.HasRadiusSet() {
			t.Errorf("%s was already set", h)
		}
		at.SetRadius(el.radius)
		if v := at.Radius(); v != el.radius {
			t.Errorf("wrong %s => %q, wanted %q", h, v, el.radius)
		}
		if !at.HasRadiusSet() {
			t.Errorf("%s was not set", h)
		}

		// lje
		if at.HasLJEnergySet() {
			t.Errorf("%s was already set", h)
		}
		at.SetLJEnergy(el.lje)
		if v := at.LJEnergy(); v != el.lje {
			t.Errorf("wrong %s => %q, wanted %q", h, v, el.lje)
		}
		if !at.HasLJEnergySet() {
			t.Errorf("%s was not set", h)
		}

		// lje14
		if at.HasLJEnergy14Set() {
			t.Errorf("%s was already set", h)
		}
		at.SetLJEnergy14(el.lje14)
		if v := at.LJEnergy14(); v != el.lje14 {
			t.Errorf("wrong %s => %q, wanted %q", h, v, el.lje14)
		}
		if !at.HasLJEnergy14Set() {
			t.Errorf("%s was not set", h)
		}

		// ljd
		if at.HasLJDistanceSet() {
			t.Errorf("%s was already set", h)
		}
		at.SetLJDistance(el.ljd)
		if v := at.LJDistance(); v != el.ljd {
			t.Errorf("wrong %s => %q, wanted %q", h, v, el.ljd)
		}
		if !at.HasLJDistanceSet() {
			t.Errorf("%s was not set", h)
		}

		// ljd14
		if at.HasLJDistance14Set() {
			t.Errorf("%s was already set", h)
		}
		at.SetLJDistance14(el.ljd14)
		if v := at.LJDistance14(); v != el.ljd14 {
			t.Errorf("wrong %s => %q, wanted %q", h, v, el.ljd14)
		}
		if !at.HasLJDistance14Set() {
			t.Errorf("%s was not set", h)
		}

	}

}

func TestAtom(t *testing.T) {
	var atoms = []struct {
		name      string
		serial    int64
		bfact     float64
		occ       float64
		alt       string
		is_hetero bool
		coord     [3]float64
		atype     string
	}{
		{"CA", 1, 0.1, 0.2, "A", true, [3]float64{0.4, 0.5, 0.6}, "C"},
	}

	for _, el := range atoms {
		a := NewAtom(el.name)
		if v := a.Name(); v != el.name {
			t.Errorf("name is not right => %q, expected %q", v, el.name)
		}

		a.SetSerial(el.serial)
		if v := a.Serial(); v != el.serial {
			t.Errorf("serial is not right => %q, expected %q", v, el.serial)
		}

		a.SetBFactor(el.bfact)
		if v := a.BFactor(); v != el.bfact {
			t.Errorf("bfact is not right => %q, expected %q", v, el.bfact)
		}

		a.SetOccupancy(el.occ)
		if v := a.Occupancy(); v != el.occ {
			t.Errorf("occ is not right => %q, expected %q", v, el.occ)
		}

		a.SetAltLoc(el.alt)
		if v := a.AltLoc(); v != el.alt {
			t.Errorf("alt is not right => %q, expected %q", v, el.alt)
		}

		a.SetHetero(el.is_hetero)
		if v := a.IsHetero(); v != el.is_hetero {
			t.Errorf("is_hetero is not right => %q, expected %q", v, el.is_hetero)
		}

		a.AddCoord(el.coord)
		if v := len(a.Coords()); v != 1 {
			t.Errorf("len(a.Coords()) is not right => %q, expected 1", v)
		}

		c := a.Coords()[0]
		if c[0] != el.coord[0] || c[1] != el.coord[1] || c[2] != el.coord[2] {
			t.Errorf("first coord is not right => %q, expected %q", c, el.coord)
		}

		at := NewAtomType(el.atype, AT_TYPE_GMX_1)
		a.SetType(at)
		if v := a.Type().Label(); v != el.atype {
			t.Errorf("wrong atom type => %q, expected %q", v, el.atype)
		}
	}

	// id
	a1 := NewAtom("C")
	a2 := NewAtom("N")
	if a1.Id() == a2.Id() {
		t.Errorf("atom ids are identical => %q", a1.Id())
	}

	// bonds
	n1 := len(a1.Bonds())
	n2 := len(a2.Bonds())
	if n1 != 0 || n2 != 0 {
		t.Errorf("initial # bonds is not zero => %q, %q", n1, n2)
	}

	NewBond(a1, a2)

	n3 := len(a1.Bonds())
	n4 := len(a2.Bonds())
	if n3 != 1 || n4 != 1 {
		t.Errorf("# bonds is not 1 => %q, %q", n3, n4)
	}

}
