package blocks

import (
	"math"
	"testing"
)

func TestAtomTypeConversion(t *testing.T) {

	var vals = []struct {
		a1_type                            ATSetting
		a1_label                           string
		a1_lje, a1_ljd, a1_lje14, a1_ljd14 float64

		a2_type                            ATSetting
		a2_label                           string
		a2_lje, a2_ljd, a2_lje14, a2_ljd14 float64

		mass, parchg, rad float64
		prot, chg         int
	}{
		{
			AT_TYPE_CHM_1,
			"C",
			-1,
			-2,
			-3,
			-4,
			AT_TYPE_GMX_1,
			"C",
			1 * 4.184,
			-2 * 2 * 0.1 / math.Pow(2.0, 1.0/6.0),
			3 * 4.184,
			-4 * 2 * 0.1 / math.Pow(2.0, 1.0/6.0),
			10, 11, 12, 13, 14, // mass ...
		},
	}

	for _, el := range vals {

		a1 := NewAtomType(el.a1_label, el.a1_type)
		a1.SetLJEnergy(el.a1_lje)
		a1.SetLJEnergy14(el.a1_lje14)
		a1.SetLJDistance(el.a1_ljd)
		a1.SetLJDistance14(el.a1_ljd14)
		a1.SetMass(el.mass)
		a1.SetCharge(el.chg)
		a1.SetPartialCharge(el.parchg)
		a1.SetProtons(el.prot)
		a1.SetRadius(el.rad)

		a2, err := a1.ConvertTo(el.a2_type)
		if err != nil {
			t.Errorf("%s", err)
		}

		if v := a2.LJEnergy(); v != el.a2_lje {
			t.Errorf("atomtype lje conversion is wrong => %f, expected %f", v, el.a2_lje)
		}

		if v := a2.LJEnergy14(); v != el.a2_lje14 {
			t.Errorf("atomtype lje14 conversion is wrong => %f, expected %f", v, el.a2_lje14)
		}

		if v := a2.LJDistance(); v != el.a2_ljd {
			t.Errorf("atomtype ljd conversion is wrong => %f, expected %f", v, el.a2_ljd)
		}

		if v := a2.LJDistance14(); v != el.a2_ljd14 {
			t.Errorf("atomtype ljd14 conversion is wrong => %f, expected %f", v, el.a2_ljd14)
		}

		if v := a2.Mass(); v != el.mass {
			t.Errorf("mass is not right => %f, expected %f", v, el.mass)
		}

		if v := a2.PartialCharge(); v != el.parchg {
			t.Errorf("parchg is not right => %f, expected %f", v, el.parchg)

		}
		if v := a2.Radius(); v != el.rad {
			t.Errorf("radius is not right => %f, expected %f", v, el.rad)

		}
		if v := a2.Charge(); v != el.chg {
			t.Errorf("charge is not right => %d, expected %d", v, el.chg)

		}
		if v := a2.Protons(); v != el.prot {
			t.Errorf("protons is not right => %d, expected %d", v, el.prot)
		}
	}
}

func TestBondTypeConversion(t *testing.T) {

}

func TestAngleTypeConversion(t *testing.T) {

}

func TestDihedralTypeConversion(t *testing.T) {

}

func TestImproperTypeConversion(t *testing.T) {

}

func TestPairTypeConversion(t *testing.T) {

}

func TestCMapTypeConversion(t *testing.T) {

}
