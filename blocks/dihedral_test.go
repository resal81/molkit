package blocks

import (
	"testing"
)

func TestDihedralType(t *testing.T) {
	var dhts = []struct {
		at1, at2, at3, at4 string
		phi, kp            float64
		mult               int
	}{
		{"C", "N", "O", "P", 11, 12, 1},
	}

	for _, el := range dhts {
		dt := NewDihedralType(el.at1, el.at2, el.at3, el.at4, DT_TYPE_CHM_1)
		if v := dt.Setting(); v&DT_TYPE_CHM_1 == 0 {
			t.Errorf("DT_TYPE is wrong => %q, expected %q", v, DT_TYPE_CHM_1)
		}

		// types
		if v := dt.AType1(); v != el.at1 {
			t.Errorf("wrong atomtype1 => %q, expected %q", v, el.at1)
		}
		if v := dt.AType2(); v != el.at2 {
			t.Errorf("wrong atomtype2 => %q, expected %q", v, el.at2)
		}
		if v := dt.AType3(); v != el.at3 {
			t.Errorf("wrong atomtype3 => %q, expected %q", v, el.at3)
		}
		if v := dt.AType4(); v != el.at4 {
			t.Errorf("wrong atomtype4 => %q, expected %q", v, el.at4)
		}

		// phi
		dt.SetPhi(el.phi)
		if v := dt.Phi(); v != el.phi {
			t.Errorf("wrong phi => %q, expected %q", v, el.phi)
		}
		if !dt.HasPhiSet() {
			t.Errorf("phi is not set")
		}

		// phi const
		dt.SetPhiConstant(el.kp)
		if v := dt.PhiConstant(); v != el.kp {
			t.Errorf("wrong phi const => %q, expected %q", v, el.kp)
		}
		if !dt.HasPhiConstantSet() {
			t.Errorf("phi const is not set")
		}

		// mult
		dt.SetMultiplicity(el.mult)
		if v := dt.Multiplicity(); v != el.mult {
			t.Errorf("wrong mult => %q, expected %q", v, el.mult)
		}
		if !dt.HasMultiplicitySet() {
			t.Errorf("mult is not set")
		}
	}

}

func TestDihedral(t *testing.T) {
	var dihedrals = []struct {
		an1, an2, an3, an4 string
		at1, at2, at3, at4 string
	}{
		{"C1", "C2", "C3", "C4", "N1", "N2", "N3", "N4"},
	}

	for _, el := range dihedrals {
		a1 := NewAtom(el.an1)
		a2 := NewAtom(el.an2)
		a3 := NewAtom(el.an3)
		a4 := NewAtom(el.an4)

		d := NewDihedral(a1, a2, a3, a4)
		if v := d.Atom1(); v.Name() != el.an1 {
			t.Errorf("an1 is wrong => %q, expected %q", v, el.an1)
		}
		if v := d.Atom2(); v.Name() != el.an2 {
			t.Errorf("an2 is wrong => %q, expected %q", v, el.an2)
		}
		if v := d.Atom3(); v.Name() != el.an3 {
			t.Errorf("an3 is wrong => %q, expected %q", v, el.an3)
		}
		if v := d.Atom4(); v.Name() != el.an4 {
			t.Errorf("an4 is wrong => %q, expected %q", v, el.an4)
		}

		dt := NewDihedralType(el.at1, el.at2, el.at3, el.at4, DT_TYPE_GMX_1)
		d.SetType(dt)
		if v := d.Type().AType1(); v != el.at1 {
			t.Errorf("dihtype atype1 is not right => %q, expected %q", v, el.at1)
		}
	}

	d1 := NewDihedral(NewAtom("C"), NewAtom("C"), NewAtom("C"), NewAtom("C"))
	d2 := NewDihedral(NewAtom("C"), NewAtom("C"), NewAtom("C"), NewAtom("C"))
	if d1.Id() == d2.Id() {
		t.Errorf("dihedral ids are identical => %q", d1.Id())
	}

}
