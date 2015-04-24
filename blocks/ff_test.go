package blocks

import (
	"testing"
)

func TestGMXSetup(t *testing.T) {
	var gs = []struct {
		nbf      int
		combrule int
		genpairs bool
		flj      float64
		fqq      float64
	}{
		{1, 2, true, 0.5, 0.8},
	}

	for _, el := range gs {
		g := NewGMXSetup(el.nbf, el.combrule, el.genpairs, el.flj, el.fqq)

		if v := g.NbFunc(); v != el.nbf {
			t.Errorf("wrong nbFunc => %q, expected %q", v, el.nbf)
		}
		if v := g.CombinationRule(); v != el.combrule {
			t.Errorf("wrong combRule => %q, expected %q", v, el.combrule)
		}
		if v := g.GeneratePairs(); v != el.genpairs {
			t.Errorf("wrong genPairs => %q, expected %q", v, el.genpairs)
		}
		if v := g.FudgeLJ(); v != el.flj {
			t.Errorf("wrong fudgeLJ => %q, expected %q", v, el.flj)
		}
		if v := g.FudgeQQ(); v != el.fqq {
			t.Errorf("wrong fudgeQQ => %q, expected %q", v, el.fqq)
		}
	}

}

func TestForceField(t *testing.T) {
	var parms = []struct {
		at1, at2, at3, at4 string
	}{
		{"CA", "N", "O", "C"},
	}

	for _, el := range parms {
		ff := NewForceField(FF_TYPE_GMX)
		if v := ff.Setting(); v&FF_TYPE_GMX == 0 {
			t.Errorf("ff type is not right => %q, expected %q", v, FF_TYPE_GMX)
		}

		ff.AddAtomType(NewAtomType(el.at1, AT_TYPE_GMX_1))
		ff.AddAtomType(NewAtomType(el.at2, AT_TYPE_GMX_1))
		ff.AddAtomType(NewAtomType(el.at3, AT_TYPE_GMX_1))
		ff.AddAtomType(NewAtomType(el.at4, AT_TYPE_GMX_1))

		if v := len(ff.AtomTypes()); v != 4 {
			t.Errorf("wrong # of atom types => %d, expected %d", v, 4)
		}

		ff.AddBondType(NewBondType(el.at1, el.at2, BT_TYPE_GMX_1))
		ff.AddBondType(NewBondType(el.at2, el.at3, BT_TYPE_GMX_1))
		ff.AddBondType(NewBondType(el.at3, el.at4, BT_TYPE_GMX_1))

		if v := len(ff.BondTypes()); v != 3 {
			t.Errorf("wrong # of bond types => %d, expected %d", v, 3)
		}

		ff.AddAngleType(NewAngleType(el.at1, el.at2, el.at3, NT_TYPE_CHM_2))
		ff.AddAngleType(NewAngleType(el.at2, el.at3, el.at4, NT_TYPE_CHM_2))

		if v := len(ff.AngleTypes()); v != 2 {
			t.Errorf("wrong # of angle types => %d, expected %d", v, 2)
		}

		ff.AddDihedralType(NewDihedralType(el.at1, el.at2, el.at3, el.at4, DT_TYPE_GMX_1))
		ff.AddImproperType(NewImproprtType(el.at1, el.at2, el.at3, el.at4, IT_TYPE_GMX_1))

		if v := len(ff.DihedralTypes()); v != 1 {
			t.Errorf("wrong # of dihedral types => %d, expected %d", v, 1)
		}
		if v := len(ff.ImproperTypes()); v != 1 {
			t.Errorf("wrong # of improper types => %d, expected %d", v, 1)
		}

		ff.AddNonBondedType(NewPairType(el.at1, el.at2, PT_TYPE_GMX_1))
		ff.AddOneFourType(NewPairType(el.at1, el.at2, PT_TYPE_GMX_1))

		if v := len(ff.NonBondedTypes()); v != 1 {
			t.Errorf("wrong # of nonbonded types => %d, expected %d", v, 1)
		}
		if v := len(ff.OneFourTypes()); v != 1 {
			t.Errorf("wrong # of onefour types => %d, expected %d", v, 1)
		}

		// two entries per cmap
		ff.AddCMapType(NewCMapType(el.at1, el.at2, el.at3, el.at4, el.at1, el.at2, el.at3, el.at4, CM_TYPE_GMX_1))
		if v := len(ff.CMapTypes()); v != 2 {
			t.Errorf("wrong # of cmap types => %d, expected %d", v, 1)
		}

	}

	// fragments

	f1 := NewFragment("ALA")
	f2 := NewFragment("LYS")

	ff := NewForceField(FF_TYPE_CHM)

	ff.AddFragment(f1)
	ff.AddFragment(f2)

	if v := len(ff.Fragments()); v != 2 {
		t.Errorf("wrong number of fragments => %d, %d", v, 2)
	}
}
