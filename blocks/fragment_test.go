package blocks

import (
	"testing"
)

func TestFragment(t *testing.T) {
	var frgs = []struct {
		name   string
		serial int64
	}{
		{"ALA", 1},
	}

	for _, el := range frgs {
		f := NewFragment(el.name)
		if v := f.Name(); v != el.name {
			t.Errorf("name is not right => %q, expected %q", v, el.name)
		}

		f.SetSerial(el.serial)
		if v := f.Serial(); v != el.serial {
			t.Errorf("serial is not rigth => %q, expected %q", v, el.serial)
		}
	}

	f := NewFragment("ALA")

	if v := len(f.Atoms()); v != 0 {
		t.Errorf("# of atoms is not zero => %q", v)
	}
	if v := len(f.Bonds()); v != 0 {
		t.Errorf("# of bonds is not zero => %q", v)
	}
	if v := len(f.Angles()); v != 0 {
		t.Errorf("# of angles is not zero => %q", v)
	}
	if v := len(f.Dihedrals()); v != 0 {
		t.Errorf("# of dihedrals is not zero => %q", v)
	}
	if v := len(f.Impropers()); v != 0 {
		t.Errorf("# of impropers is not zero => %q", v)
	}

	a1 := NewAtom("CA")
	a2 := NewAtom("N")
	a3 := NewAtom("C")
	a4 := NewAtom("CA")
	f.AddAtom(a1)
	f.AddAtom(a2)
	f.AddAtom(a3)
	f.AddAtom(a4)

	b1 := NewBond(a1, a2)
	b2 := NewBond(a2, a3)
	b3 := NewBond(a3, a4)
	f.AddBond(b1)
	f.AddBond(b2)
	f.AddBond(b3)

	g1 := NewAngle(a1, a2, a3)
	g2 := NewAngle(a2, a3, a4)
	f.AddAngle(g1)
	f.AddAngle(g2)

	d1 := NewDihedral(a1, a2, a3, a4)
	i1 := NewImproper(a1, a2, a3, a4)
	f.AddDihedral(d1)
	f.AddImproper(i1)

	if v := len(f.Atoms()); v != 4 {
		t.Errorf("# of atoms is not 4 => %q", v)
	}
	if v := len(f.Bonds()); v != 3 {
		t.Errorf("# of bonds is not 3 => %q", v)
	}
	if v := len(f.Angles()); v != 2 {
		t.Errorf("# of angles is not 2 => %q", v)
	}
	if v := len(f.Dihedrals()); v != 1 {
		t.Errorf("# of dihedrals is not 1 => %q", v)
	}
	if v := len(f.Impropers()); v != 1 {
		t.Errorf("# of impropers is not 1 => %q", v)
	}

	// id
	f1 := NewFragment("ALA")
	f2 := NewFragment("LYS")
	if f1.Id() == f2.Id() {
		t.Errorf("fragment ids are identical => %q", f1.Id())
	}

}
