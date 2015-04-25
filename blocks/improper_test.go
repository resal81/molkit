package blocks

import (
	"testing"
)

func TestImproperType(t *testing.T) {
	var impts = []struct {
		at1, at2, at3, at4 string
		psi, ks            float64
	}{
		{"C", "N", "O", "P", 11, 12},
	}

	for _, el := range impts {
		dt := NewImproperType(el.at1, el.at2, el.at3, el.at4, IT_TYPE_CHM_1)
		if v := dt.Setting(); v&IT_TYPE_CHM_1 == 0 {
			t.Errorf("DT_TYPE is wrong => %q, expected %q", v, IT_TYPE_CHM_1)
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

		// psi
		dt.SetPsi(el.psi)
		if v := dt.Psi(); v != el.psi {
			t.Errorf("wrong psi => %q, expected %q", v, el.psi)
		}
		if !dt.HasPsiSet() {
			t.Errorf("psi is not set")
		}

		// psi const
		dt.SetPsiConstant(el.ks)
		if v := dt.PsiConstant(); v != el.ks {
			t.Errorf("wrong psi const => %q, expected %q", v, el.ks)
		}
		if !dt.HasPsiConstantSet() {
			t.Errorf("psi const is not set")
		}
	}

}

func TestImproper(t *testing.T) {
	var impropers = []struct {
		an1, an2, an3, an4 string
		at1, at2, at3, at4 string
	}{
		{"C1", "C2", "C3", "C4", "N1", "N2", "N3", "N4"},
	}

	for _, el := range impropers {
		a1 := NewAtom(el.an1)
		a2 := NewAtom(el.an2)
		a3 := NewAtom(el.an3)
		a4 := NewAtom(el.an4)

		d := NewImproper(a1, a2, a3, a4)
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

		dt := NewImproperType(el.at1, el.at2, el.at3, el.at4, IT_TYPE_CHM_1)
		d.SetType(dt)
		if v := d.Type().AType1(); v != el.at1 {
			t.Errorf("dihtype atype1 is not right => %q, expected %q", v, el.at1)
		}
	}

	d1 := NewImproper(NewAtom("C"), NewAtom("C"), NewAtom("C"), NewAtom("C"))
	d2 := NewImproper(NewAtom("C"), NewAtom("C"), NewAtom("C"), NewAtom("C"))
	if d1.Id() == d2.Id() {
		t.Errorf("dihedral ids are identical => %q", d1.Id())
	}
}
