package blocks

import (
	"math"
	"testing"
)

const DELTA = 1e-12

func TestAtomTypeConversion(t *testing.T) {

	var vals = []struct {
		label              string
		mass, parchg, rad  float64
		prot, chg          int
		a1_type, a2_type   ATSetting
		a1_lje, a2_lje     float64
		a1_ljd, a2_ljd     float64
		a1_lje14, a2_lje14 float64
		a1_ljd14, a2_ljd14 float64
	}{
		{
			"C",
			10, 11, 12, 13, 14,
			AT_TYPE_CHM_1, AT_TYPE_GMX_1,
			-1, 1 * 4.184,
			-2, -2 * 2 * 0.1 / math.Pow(2.0, 1.0/6.0),
			-3, 3 * 4.184,
			-4, -4 * 2 * 0.1 / math.Pow(2.0, 1.0/6.0),
		},
	}

	for _, el := range vals {

		a1 := NewAtomType(el.label, el.a1_type)
		a1.SetMass(el.mass)
		a1.SetCharge(el.chg)
		a1.SetPartialCharge(el.parchg)
		a1.SetProtons(el.prot)
		a1.SetRadius(el.rad)

		a1.SetLJEnergy(el.a1_lje)
		a1.SetLJEnergy14(el.a1_lje14)
		a1.SetLJDistance(el.a1_ljd)
		a1.SetLJDistance14(el.a1_ljd14)

		a2, err := a1.ConvertTo(el.a2_type)
		if err != nil {
			t.Errorf("%s", err)
		}

		if v := a2.Label(); v != el.label {
			t.Errorf("atomtype label is wrong => %s, expected %s", v, el.label)
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
	var vals = []struct {
		at1, at2         string
		b1_type, b2_type BTSetting
		b1_kb, b2_kb     float64
		b1_b0, b2_b0     float64
	}{
		{
			"A", "B",
			BT_TYPE_CHM_1, BT_TYPE_GMX_1,
			-1, -1 * 2 * 4.184 * 100,
			-2, -2 * 0.1,
		},
	}

	for _, el := range vals {
		bt1 := NewBondType(el.at1, el.at2, el.b1_type)
		bt1.SetHarmonicConstant(el.b1_kb)
		bt1.SetHarmonicDistance(el.b1_b0)

		bt2, err := bt1.ConvertTo(el.b2_type)
		if err != nil {
			t.Errorf("%s", err)
		}

		if v := bt2.AType1(); v != el.at1 {
			t.Errorf("wrong at1 => %s, expected %s", v, el.at1)
		}
		if v := bt2.AType2(); v != el.at2 {
			t.Errorf("wrong at2 => %s, expected %s", v, el.at2)
		}
		if v := bt2.HarmonicConstant(); math.Abs(v-el.b2_kb) > DELTA {
			t.Errorf("wrong kb => %f, expected %f", v, el.b2_kb)
		}
		if v := bt2.HarmonicDistance(); math.Abs(v-el.b2_b0) > DELTA {
			t.Errorf("wrong b0 => %f, expected %f", v, el.b2_b0)
		}
	}

}

func TestAngleTypeConversion(t *testing.T) {
	var vals = []struct {
		at1, at2, at3      string
		a1_type, a2_type   NTSetting
		a1_theta, a2_theta float64
		a1_kt, a2_kt       float64
		a1_r13, a2_r13     float64
		a1_kub, a2_kub     float64
	}{
		{
			"A", "B", "C",
			NT_TYPE_CHM_1, NT_TYPE_GMX_5,
			1.0, 1.0,
			2.0, 2.0 * 2 * 4.184,
			3.0, 3.0 * 0.1,
			4.0, 4.0 * 2 * 4.184 * 100,
		},
	}

	for _, el := range vals {
		at1 := NewAngleType(el.at1, el.at2, el.at3, el.a1_type)
		at1.SetTheta(el.a1_theta)
		at1.SetThetaConstant(el.a1_kt)
		at1.SetR13(el.a1_r13)
		at1.SetUBConstant(el.a1_kub)

		at2, err := at1.ConvertTo(el.a2_type)
		if err != nil {
			t.Errorf("%s", err)
		}

		if v := at2.AType1(); v != el.at1 {
			t.Errorf("wrong at1 => %s, expected %s", v, el.at1)
		}
		if v := at2.AType2(); v != el.at2 {
			t.Errorf("wrong at2 => %s, expected %s", v, el.at2)
		}
		if v := at2.AType3(); v != el.at3 {
			t.Errorf("wrong at3 => %s, expected %s", v, el.at3)
		}

		if v := at2.Theta(); math.Abs(v-el.a2_theta) > DELTA {
			t.Errorf("wrong theta => %f, expected %f", v, el.a2_theta)
		}
		if v := at2.ThetaConstant(); math.Abs(v-el.a2_kt) > DELTA {
			t.Errorf("wrong kt => %f, expected %f", v, el.a2_kt)
		}
		if v := at2.R13(); math.Abs(v-el.a2_r13) > DELTA {
			t.Errorf("wrong r13 => %f, expected %f, diff %v", v, el.a2_r13, v-el.a2_r13)
		}
		if v := at2.UBConstant(); math.Abs(v-el.a2_kub) > DELTA {
			t.Errorf("wrong kub => %f, expected %f, diff %v", v, el.a2_kub, v-el.a2_kub)
		}

	}

}

func TestDihedralTypeConversion(t *testing.T) {
	var vals = []struct {
		at1, at2, at3, at4 string
		d1_type, d2_type   DTSetting
		d1_phi, d2_phi     float64
		d1_kp, d2_kp       float64
		d1_m, d2_m         int
	}{
		{
			"A", "B", "C", "D",
			DT_TYPE_CHM_1, DT_TYPE_GMX_9,
			1, 1,
			2, 2 * 4.184,
			3, 3,
		},
	}

	for _, el := range vals {
		dt1 := NewDihedralType(el.at1, el.at2, el.at3, el.at4, el.d1_type)
		dt1.SetPhi(el.d1_phi)
		dt1.SetPhiConstant(el.d1_kp)
		dt1.SetMultiplicity(el.d1_m)

		dt2, err := dt1.ConvertTo(el.d2_type)
		if err != nil {
			t.Errorf("%s", err)
		}

		if v := dt2.AType1(); v != el.at1 {
			t.Errorf("wrong at1 => %s, expected %s", v, el.at1)
		}
		if v := dt2.AType2(); v != el.at2 {
			t.Errorf("wrong at2 => %s, expected %s", v, el.at2)
		}
		if v := dt2.AType3(); v != el.at3 {
			t.Errorf("wrong at3 => %s, expected %s", v, el.at3)
		}
		if v := dt2.AType4(); v != el.at4 {
			t.Errorf("wrong at4 => %s, expected %s", v, el.at4)
		}

		if v := dt2.Phi(); math.Abs(v-el.d2_phi) > DELTA {
			t.Errorf("wrong phi => %f, expected %f", v, el.d2_phi)
		}
		if v := dt2.PhiConstant(); math.Abs(v-el.d2_kp) > DELTA {
			t.Errorf("wrong kp => %f, expected %f", v, el.d2_kp)
		}
		if v := dt2.Multiplicity(); v != el.d2_m {
			t.Errorf("wrong mult => %d, expected %d", v, el.d2_m)
		}
	}

}

func TestImproperTypeConversion(t *testing.T) {
	var vals = []struct {
		at1, at2, at3, at4 string
		d1_type, d2_type   ITSetting
		d1_psi, d2_psi     float64
		d1_ks, d2_ks       float64
	}{
		{
			"A", "B", "C", "D",
			IT_TYPE_CHM_1, IT_TYPE_GMX_2,
			1, 1,
			2, 2 * 2 * 4.184,
		},
	}

	for _, el := range vals {
		dt1 := NewImproperType(el.at1, el.at2, el.at3, el.at4, el.d1_type)
		dt1.SetPsi(el.d1_psi)
		dt1.SetPsiConstant(el.d1_ks)

		dt2, err := dt1.ConvertTo(el.d2_type)
		if err != nil {
			t.Errorf("%s", err)
		}

		if v := dt2.AType1(); v != el.at1 {
			t.Errorf("wrong at1 => %s, expected %s", v, el.at1)
		}
		if v := dt2.AType2(); v != el.at2 {
			t.Errorf("wrong at2 => %s, expected %s", v, el.at2)
		}
		if v := dt2.AType3(); v != el.at3 {
			t.Errorf("wrong at3 => %s, expected %s", v, el.at3)
		}
		if v := dt2.AType4(); v != el.at4 {
			t.Errorf("wrong at4 => %s, expected %s", v, el.at4)
		}

		if v := dt2.Psi(); math.Abs(v-el.d2_psi) > DELTA {
			t.Errorf("wrong psi => %f, expected %f", v, el.d2_psi)
		}
		if v := dt2.PsiConstant(); math.Abs(v-el.d2_ks) > DELTA {
			t.Errorf("wrong ks => %f, expected %f", v, el.d2_ks)
		}
	}

}

func TestPairTypeConversion(t *testing.T) {
	var vals = []struct {
		at1, at2           string
		p1_type, p2_type   PTSetting
		p1_lje, p2_lje     float64
		p1_lje14, p2_lje14 float64
		p1_ljd, p2_ljd     float64
		p1_ljd14, p2_ljd14 float64
	}{
		{
			"A", "B",
			PT_TYPE_CHM_1, PT_TYPE_GMX_1,
			-1, 1 * 4.184,
			-2, 2 * 4.184,
			3, 3 * 0.1 / math.Pow(2, 1.0/6.0),
			4, 4 * 2 * 0.1 / math.Pow(2, 1.0/6.0),
		},
	}

	for _, el := range vals {
		pt1 := NewPairType(el.at1, el.at2, el.p1_type)
		pt1.SetLJEnergy(el.p1_lje)
		pt1.SetLJEnergy14(el.p1_lje14)
		pt1.SetLJDistance(el.p1_ljd)
		pt1.SetLJDistance14(el.p1_ljd14)

		pt2, err := pt1.ConvertTo(el.p2_type)
		if err != nil {
			t.Errorf("%s", err)
		}

		if v := pt2.AType1(); v != el.at1 {
			t.Errorf("wrong at1 => %s, expected %s", v, el.at1)
		}
		if v := pt2.AType2(); v != el.at2 {
			t.Errorf("wrong at2 => %s, expected %s", v, el.at2)
		}

		if v := pt2.LJEnergy(); math.Abs(v-el.p2_lje) > DELTA {
			t.Errorf("wrong lje => %f, expected %f", v, el.p2_lje)
		}
		if v := pt2.LJEnergy14(); math.Abs(v-el.p2_lje14) > DELTA {
			t.Errorf("wrong lje14 => %f, expected %f", v, el.p2_lje14)
		}
		if v := pt2.LJDistance(); math.Abs(v-el.p2_ljd) > DELTA {
			t.Errorf("wrong ljd => %f, expected %f", v, el.p2_ljd)
		}
		if v := pt2.LJDistance14(); math.Abs(v-el.p2_ljd14) > DELTA {
			t.Errorf("wrong ljd14 => %f, expected %f", v, el.p2_ljd14)
		}
	}
}

func TestCMapTypeConversion(t *testing.T) {
	var vals = []struct {
		at1, at2, at3, at4, at5, at6, at7, at8 string
		c1_type, c2_type                       CTSetting
		c1_values                              []float64
		c2_values                              []float64
	}{
		{
			"A", "B", "C", "D", "E", "F", "G", "H",
			CT_TYPE_CHM_1, CT_TYPE_GMX_1,
			[]float64{1, 2, 3},
			[]float64{1 * 4.184, 2 * 4.184, 3 * 4.184},
		},
	}

	for _, el := range vals {
		ct1 := NewCMapType(el.at1, el.at2, el.at3, el.at4, el.at5, el.at6, el.at7, el.at8, el.c1_type)
		ct1.SetValues(el.c1_values)

		ct2, err := ct1.ConvertTo(el.c2_type)
		if err != nil {
			t.Errorf("%s", err)
		}

		if v := ct2.AType1(); v != el.at1 {
			t.Errorf("wrong at1 => %s, expected %s", v, el.at1)
		}
		if v := ct2.AType2(); v != el.at2 {
			t.Errorf("wrong at2 => %s, expected %s", v, el.at2)
		}
		if v := ct2.AType3(); v != el.at3 {
			t.Errorf("wrong at3 => %s, expected %s", v, el.at3)
		}
		if v := ct2.AType4(); v != el.at4 {
			t.Errorf("wrong at4 => %s, expected %s", v, el.at4)
		}
		if v := ct2.AType5(); v != el.at5 {
			t.Errorf("wrong at5 => %s, expected %s", v, el.at5)
		}
		if v := ct2.AType6(); v != el.at6 {
			t.Errorf("wrong at6 => %s, expected %s", v, el.at6)
		}
		if v := ct2.AType7(); v != el.at7 {
			t.Errorf("wrong at7 => %s, expected %s", v, el.at7)
		}
		if v := ct2.AType8(); v != el.at8 {
			t.Errorf("wrong at8 => %s, expected %s", v, el.at8)
		}

		for i, v := range ct2.Values() {
			if math.Abs(v-el.c2_values[i]) > DELTA {
				t.Errorf("wrong value for cmap => %f, expected %f", v, el.c2_values[i])
			}
		}
	}

}
