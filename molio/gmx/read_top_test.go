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

var testCase2_string = `
[ moleculetype ]
; name  nrexcl
LIG1         3

[ atoms ]
; nr    type    resnr   residu  atom    cgnr    charge  mass
     1       CRL1      1     LIG1     C3      1      0.140    12.0110   ; qtot  0.140
     2        OHL      1     LIG1     O3      2     -0.660    15.9994   ; qtot -0.520
     3        HOL      1     LIG1    H3'      3      0.430     1.0080   ; qtot -0.090
     4       HGA1      1     LIG1     H3      4      0.090     1.0080   ; qtot -0.000
     5       CRL2      1     CHL1     C4      5     -0.180    12.0110   ; qtot -0.180
     6       HGA2      1     CHL1    H4A      6      0.090     1.0080   ; qtot -0.090
     7       HGA2      1     CHL1    H4B      7      0.090     1.0080   ; qtot -0.000
     8       CEL1      1     CHL1     C5      8      0.000    12.0110   ; qtot -0.000
    
[ bonds ]
; ai    aj  funct   b0  Kb
    1     2     1
    1     4     1
    2     3     1

[ pairs ]
; ai    aj  funct   c6  c12
    3     4     1 
    1     4     1 
    
[ angles ] 
; ai    aj  ak  funct   th0 cth S0  Kub
    2     1     3     5
    2     1     6     5
   
[ dihedrals ]
; ai    aj  ak  al  funct   phi0    cp  mult
    2     1     6     7     9
    5     1     8     3     9

[ dihedrals ]
; ai    aj  ak  al  funct   q0  cq
   5    6    1    4     2

[ position_restraints ]
   8     1      0.0      0.0  1000.0

[ dihedral_restraints ]
   5    1    8    3     1    120.0      2.5  1000.0
    
[ moleculetype ]
; name  nrexcl
TIP3         2

[ atoms ]
; nr    type    resnr   residu  atom    cgnr    charge  mass
     1         OT      1     TIP3    OH2      1     -0.834    15.9994   ; qtot -0.834
     2         HT      1     TIP3     H1      2      0.417     1.0080   ; qtot -0.417
     3         HT      1     TIP3     H2      3      0.417     1.0080   ; qtot  0.000

[ settles ]
; i j   funct   length
1   1    9.572000e-02    1.513900e-01

[ exclusions ]
1 2 3
2 1 3
3 1 2

[ system ]
; Name
Title

[ molecules ]
; Compound  #mols
LIG1              3
TIP3            100
`

func TestCase2(t *testing.T) {
	top, _, err := ReadTOPString(testCase2_string)
	if err != nil {
		t.Fatalf("%s", err)
	}

	rp := top.RegisteredTopPolymers()
	if len(rp) != 2 {
		t.Errorf("shoud have registered 2 polymers, but has %d", len(rp))
	}
}

func TestReadTop(t *testing.T) {
	_, _, err := ReadTOPFile("../../testdata/tmp.top")
	if err != nil {
		t.Errorf("could not read top file: %s", err)
	}

}
