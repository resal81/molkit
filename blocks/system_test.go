package blocks

import (
	"testing"
)

func TestSystem(t *testing.T) {
	// id
	s1 := NewSystem()
	s2 := NewSystem()
	if s1.Id() == s2.Id() {
		t.Errorf("system ids are identical => %q", s1.Id())
	}

	//
	f := NewSystem()

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
	if v := len(f.Links()); v != 0 {
		t.Errorf("# of links is not zero => %q", v)
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

	l1 := NewLink()
	l1.AddConnection(NewConnection(a1, a2))
	l1.AddConnection(NewConnection(a2, a3))
	f.AddLink(l1)

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
	if v := len(f.Links()); v != 1 {
		t.Errorf("# of links is not 1 => %q", v)
	}
}
