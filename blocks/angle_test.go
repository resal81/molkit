package blocks

import (
	"testing"
)

func TestAngleType(t *testing.T) {
	var angts = []struct {
		atype1              string
		atype2              string
		atype3              string
		theta, kt, r13, kub float64
	}{
		{"CA", "N", "C", 110, 1, 2, 3},
	}

	for _, el := range angts {
		at := NewAngleType(el.atype1, el.atype2, el.atype3)

		// types
		if t1 := at.AType1(); t1 != el.atype1 {
			t.Errorf("atype1 is not correct => %q, wanted %q", t1, el.atype1)
		}
		if t2 := at.AType2(); t2 != el.atype2 {
			t.Errorf("atype2 is not correct => %q, wanted %q", t2, el.atype2)
		}
		if t3 := at.AType3(); t3 != el.atype3 {
			t.Errorf("atype3 is not correct => %q, wanted %q", t3, el.atype3)
		}

		// theta
		if at.HasThetaSet() {
			t.Errorf("theta is already set")
		}

		at.SetTheta(el.theta)

		if v := at.Theta(); v != el.theta {
			t.Errorf("theta is not right => %q, wanted %q", v, el.theta)
		}

		if !at.HasThetaSet() {
			t.Errorf("theta was not set")
		}

		// theta constant
		if at.HasThetaConstantSet() {
			t.Errorf("theta const is already set")
		}

		at.SetThetaConstant(el.kt)

		if v := at.ThetaConstant(); v != el.kt {
			t.Errorf("theta const is not right => %q, wanted %q", v, el.kt)
		}

		if !at.HasThetaConstantSet() {
			t.Errorf("theta const was not set")
		}

		// r13
		if at.HasR13Set() {
			t.Errorf("r13 is already set")
		}

		at.SetR13(el.r13)

		if v := at.R13(); v != el.r13 {
			t.Errorf("r13 is not right => %q, wanted %q", v, el.r13)
		}

		if !at.HasR13Set() {
			t.Errorf("r13 was not set")
		}

		// kub
		if at.HasUBConstantSet() {
			t.Errorf("kub is already set")
		}

		at.SetUBConstant(el.kub)

		if v := at.UBConstant(); v != el.kub {
			t.Errorf("kub is not right => %q, wanted %q", v, el.kub)
		}

		if !at.HasUBConstantSet() {
			t.Errorf("kub was not set")
		}

	}
}

func TestAngle(t *testing.T) {
	var angles = []struct {
		aname1 string
		aname2 string
		aname3 string
		atype1 string
		atype2 string
		atype3 string
	}{
		{"N", "O", "C", "NH", "OG", "C3"},
	}

	for _, el := range angles {
		a1 := NewAtom(el.aname1)
		a2 := NewAtom(el.aname2)
		a3 := NewAtom(el.aname3)

		ang := NewAngle(a1, a2, a3)

		// atom names
		if n := ang.Atom1().Name(); n != el.aname1 {
			t.Errorf("aname1 is not correct => %q, wanted %q", n, el.aname1)
		}

		if n := ang.Atom2().Name(); n != el.aname2 {
			t.Errorf("aname2 is not correct => %q, wanted %q", n, el.aname2)
		}

		if n := ang.Atom3().Name(); n != el.aname3 {
			t.Errorf("aname3 is not correct => %q, wanted %q", n, el.aname3)
		}

		// angle type
		at := NewAngleType(el.atype1, el.atype2, el.atype3)
		ang.SetType(at)

		if lb := ang.Type().AType1(); lb != el.atype1 {
			t.Errorf("angletype.atype1 is not correct => %q, wanted %q", lb, el.atype1)
		}

		if lb := ang.Type().AType2(); lb != el.atype2 {
			t.Errorf("angletype.atype2 is not correct => %q, wanted %q", lb, el.atype2)
		}

		if lb := ang.Type().AType3(); lb != el.atype3 {
			t.Errorf("angletype.atype3 is not correct => %q, wanted %q", lb, el.atype3)
		}
	}

}
