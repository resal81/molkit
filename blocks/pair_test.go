package blocks

import (
	"testing"
)

func TestPairType(t *testing.T) {
	var pts = []struct {
		at1, at2               string
		lje, lje14, ljd, ljd14 float64
	}{
		{"C", "N", 1e-4, 2e-4, 3e-4, 4e-4},
	}

	for _, el := range pts {
		pt := NewPairType(el.at1, el.at2, PT_TYPE_GMX_1)

		if v := pt.Setting(); v&PT_TYPE_GMX_1 == 0 {
			t.Errorf("PT_TYPE is not right => %q, expected %q", v, PT_TYPE_GMX_1)
		}

		// lje
		pt.SetLJEnergy(el.lje)
		if v := pt.LJEnergy(); v != el.lje {
			t.Errorf("wrong lje => %q, expected %q", v, el.lje)
		}
		if !pt.HasLJEnergySet() {
			t.Errorf("lje is not set")
		}

		// lje14
		pt.SetLJEnergy14(el.lje14)
		if v := pt.LJEnergy14(); v != el.lje14 {
			t.Errorf("wrong lje14 => %q, expected %q", v, el.lje14)
		}
		if !pt.HasLJEnergy14Set() {
			t.Errorf("lje is not set")
		}

		// ljd
		pt.SetLJDist(el.ljd)
		if v := pt.LJDist(); v != el.ljd {
			t.Errorf("wrong ljd => %q, expected %q", v, el.ljd)
		}
		if !pt.HasLJDistSet() {
			t.Errorf("ljd is not set")
		}

		// ljd14
		pt.SetLJDist14(el.ljd14)
		if v := pt.LJDist14(); v != el.ljd14 {
			t.Errorf("wrong ljd14 => %q, expected %q", v, el.ljd14)
		}
		if !pt.HasLJDist14Set() {
			t.Errorf("ljd14 is not set")
		}
	}

}

func TestPair(t *testing.T) {
	var pairs = []struct {
		aname1 string
		aname2 string
		atype1 string
		atype2 string
	}{
		{"CA", "N", "C3", "NH"},
	}

	for _, el := range pairs {
		a1 := NewAtom(el.aname1)
		a2 := NewAtom(el.aname2)

		p := NewPair(a1, a2)

		// atom names
		if an := p.Atom1().Name(); an != el.aname1 {
			t.Errorf("aname1 is not correct => %q, wanted %q", an, el.aname1)
		}

		if an := p.Atom2().Name(); an != el.aname2 {
			t.Errorf("aname2 is not correct => %q, wanted %q", an, el.aname2)
		}

		// bond type
		pt := NewPairType(el.atype1, el.atype2, PT_TYPE_CHM_1)
		p.SetType(pt)

		if lb := p.Type().AType1(); lb != el.atype1 {
			t.Errorf("pairtype.atype1 is not correct => %q, wanted %q", lb, el.atype1)
		}

		if lb := p.Type().AType2(); lb != el.atype2 {
			t.Errorf("pairtype.atype2 is not correct => %q, wanted %q", lb, el.atype2)
		}
	}
}
