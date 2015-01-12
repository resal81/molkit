package gmx

import (
	"testing"
)

var testCase1_string string = `
;; CHARMM36 FF in GROMACS format

[ defaults ]
; nbfunc    comb-rule   gen-pairs   fudgeLJ fudgeQQ
1   2   yes 1.0 1.0

[ atomtypes ]
; name  at.num  mass    charge  ptype   sigma   epsilon ;   sigma_14    epsilon_14
    CEL1     6    12.0110     -0.150     A    3.72395664183e-01    2.845120e-01 
    CRL1     6    12.0110      0.140     A    3.58141284692e-01    1.506240e-01 ;   3.38541512893e-01    4.184000e-02 

[ nonbond_params ]
; i j   func    sigma   epsilon
   CRL1    CTL1     1  3.57250385974e-01  1.15060000000e-01 
   CRL1    HAL1     1  2.98451070577e-01  1.22591200000e-01 

[ bondtypes ]
; i j   func    b0  Kb
   CEL1    CEL1     1  1.340000e-01  3.681920e+05
   CEL1    CRL1     1  1.502000e-01  2.008320e+05

[ pairtypes ]
; i j   func    sigma1-4    epsilon1-4
   CEL1    CRL1     1  3.55468588538e-01  1.09105371453e-01 
   CEL1    CRL2     1  3.55468588538e-01  1.09105371453e-01 

[ angletypes ]
; i j   k   func    th0 cth S0  Kub
   CEL1    CEL1    CEL1     5  1.2350000e+02  4.0166400e+02  0.0000000e+00  0.0000000e+00
   CEL1    CEL1    CRL1     5  1.2350000e+02  4.0166400e+02  0.0000001e+00  0.0000002e+00

[ dihedraltypes ]
; i j   k   l   func    phi0    cp  mult
   CEL1    CEL1    CRL1    CRL1     9  1.800000e+02  2.092000e+00      1
   CEL1    CEL1    CRL1    CRL1     9  1.800000e+02  5.439200e+00      3

[ dihedraltypes ]
; i j   k   l   func    q0  cq
    OBL       X       X      CL     2  0.000001e+00  8.368000e+02
`

func checkLength(t *testing.T, length, expectedLength int, header string) {
	if length != expectedLength {
		t.Errorf("%s - length should be %d, is %d", header, expectedLength, length)
	}
}

func compareF64(t *testing.T, v1, v2 float64, header string) {
	if v1 != v2 {
		t.Errorf("%s : %f != %f", header, v1, v2)
	}
}

func compareInt8(t *testing.T, v1, v2 int8, header string) {
	if v1 != v2 {
		t.Errorf("%s : %d != %d", header, v1, v2)
	}
}

func TestCase1(t *testing.T) {
	_, ff, err := ReadTOPString(testCase1_string)
	if err != nil {
		t.Fatalf("error parsing testCase1_string -> %s", err)
	}

	checkLength(t, len(ff.AtomTypes()), 2, "AtomTypes()")
	checkLength(t, len(ff.BondTypes()), 2, "BondTypes()")
	checkLength(t, len(ff.PairTypes()), 2, "PairTypes()")
	checkLength(t, len(ff.NonBondedTypes()), 2, "NondBondedTypes()")
	checkLength(t, len(ff.AngleTypes()), 2, "AngleTypes()")
	checkLength(t, len(ff.DihedralTypes()), 2, "DihedralTypes()")
	checkLength(t, len(ff.ImproperTypes()), 1, "ImproperTypes()")

	// atomtypes
	ats := ff.AtomTypes()
	compareF64(t, ats[0].Sigma(), 3.72395664183e-01, "Sigma()")
	compareF64(t, ats[1].Epsilon(), 1.506240e-01, "Sigma()")

	// bondtypes
	bts := ff.BondTypes()
	compareF64(t, bts[0].HarmonicConstant(), 3.681920e+05, "HarmonicConstant()")
	compareF64(t, bts[1].HarmonicDistance(), 1.502000e-01, "HarmonicDistance()")

	// nonbonded
	nbs := ff.NonBondedTypes()
	compareF64(t, nbs[0].Sigma(), 3.57250385974e-01, "Sigma()")
	compareF64(t, nbs[1].Epsilon(), 1.22591200000e-01, "Epsilon()")

	// pairtypes
	pts := ff.PairTypes()
	compareF64(t, pts[0].Sigma14(), 3.55468588538e-01, "Sigma14()")
	compareF64(t, pts[1].Epsilon14(), 1.09105371453e-01, "Epsilon14()")

	// angletypes
	ags := ff.AngleTypes()
	compareF64(t, ags[0].ThetaConstant(), 4.0166400e+02, "ThetaConstant()")
	compareF64(t, ags[0].Theta(), 1.2350000e+02, "Theta()")
	compareF64(t, ags[1].UBConstant(), 0.0000002e+00, "UBConstant()")
	compareF64(t, ags[1].R13(), 0.0000001e+00, "R13()")

	// dihedraltypes
	dhs := ff.DihedralTypes()
	compareF64(t, dhs[0].PhiConstant(), 2.092000e+00, "PhiConstant()")
	compareF64(t, dhs[1].Phi(), 1.800000e+02, "Phi()")
	compareInt8(t, dhs[1].Mult(), 3, "Mult()")

	// impropertypes
	ims := ff.ImproperTypes()
	compareF64(t, ims[0].PsiConstant(), 8.368000e+02, "PsiConstant()")
	compareF64(t, ims[0].Psi(), 0.000001e+00, "Psi()")

}

// func TestReadTop(t *testing.T) {
// 	_, _, err := ReadTOPFile("../../testdata/tmp.top")
// 	if err != nil {
// 		t.Errorf("could not read top file: %s", err)
// 	}
// }
