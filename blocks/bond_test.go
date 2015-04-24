package blocks

import (
	"testing"
)

func TestBondType(t *testing.T) {
	var bts = []struct {
		atype1 string
		atype2 string
		hconst float64
		hdist  float64
	}{
		{"CA", "N", 10, 12},
	}

	for _, el := range bts {
		b := NewBondType(el.atype1, el.atype2, BT_TYPE_GMX_1)

		if v := b.Setting(); v&BT_TYPE_GMX_1 == 0 {
			t.Errorf("BT_TYPE is wrong => %q, expected %q", v, BT_TYPE_GMX_1)
		}

		// types
		if b.AType1() != el.atype1 {
			t.Errorf("atype1 is not correct")
		}
		if b.AType2() != el.atype2 {
			t.Errorf("atype2 is not correct")
		}

		// harmonic constant
		b.SetHarmonicConstant(el.hconst)

		if b.HarmonicConstant() != el.hconst {
			t.Errorf("harmonic const is not right => %q, want %q", b.HarmonicConstant(), el.hconst)
		}

		if !b.HasHarmonicConstantSet() {
			t.Errorf("harmonic constant was not set")
		}

		// harmonic distance
		b.SetHarmonicDistance(el.hdist)

		if b.HarmonicDistance() != el.hdist {
			t.Errorf("harmonic distance is not right")
		}

		if !b.HasHarmonicDistanceSet() {
			t.Errorf("harmonic distance was not set")
		}

	}
}

func TestBond(t *testing.T) {
	var bonds = []struct {
		aname1 string
		aname2 string
		atype1 string
		atype2 string
	}{
		{"CA", "N", "C3", "NH"},
	}

	for _, el := range bonds {
		a1 := NewAtom(el.aname1)
		a2 := NewAtom(el.aname2)

		b := NewBond(a1, a2)

		// atom names
		if an := b.Atom1().Name(); an != el.aname1 {
			t.Errorf("aname1 is not correct => %q, wanted %q", an, el.aname1)
		}

		if an := b.Atom2().Name(); an != el.aname2 {
			t.Errorf("aname2 is not correct => %q, wanted %q", an, el.aname2)
		}

		// bond type
		bt := NewBondType(el.atype1, el.atype2, BT_TYPE_CHM_1)
		b.SetType(bt)

		if lb := b.Type().AType1(); lb != el.atype1 {
			t.Errorf("bondtype.atype1 is not correct => %q, wanted %q", lb, el.atype1)
		}

		if lb := b.Type().AType2(); lb != el.atype2 {
			t.Errorf("bondtype.atype2 is not correct => %q, wanted %q", lb, el.atype2)
		}
	}

	b1 := NewBond(NewAtom("C"), NewAtom("N"))
	b2 := NewBond(NewAtom("P"), NewAtom("O"))
	if b1.Id() == b2.Id() {
		t.Errorf("bond ids are identical => %q", b1.Id())
	}
}
