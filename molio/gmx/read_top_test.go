package gmx

import (
	frc "github.com/resal81/molkit/ff"
	"testing"
)

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

func compareString(t *testing.T, v1, v2, header string) {
	if v1 != v2 {
		t.Errorf("%s : %d != %d", header, v1, v2)
	}
}

func TestDefaults(t *testing.T) {
	s := `
	   [ defaults ]
	   ; nbfunc    comb-rule   gen-pairs   fudgeLJ fudgeQQ
	   1   2   yes 1.0 1.0
	   `

	if _, ff, err := ReadTOPString(s); err != nil {
		t.Fatalf("could not parse the string")
	} else {
		prop := ff.PropsGMX()
		compareInt8(t, prop.NbFunc(), 1, "NbFunc()")
		compareInt8(t, prop.CombRule(), 2, "CombRule()")
		compareF64(t, prop.FudgeLJ(), 1.0, "FudgeLJ()")
		compareF64(t, prop.FudgeQQ(), 1.0, "FudgeQQ()")
	}
}

func TestAtomTypes(t *testing.T) {
	s := `
    [ defaults ]
    [ atomtypes ]
    ; name  at.num  mass      charge  ptype   sigma                epsilon 
    CEL1     6    12.0110     -0.150     A    3.72395664183e-01    2.845120e-01 
    CRL1     6    12.0110      0.140     A    3.58141284692e-01    1.506240e-01 
    `

	if _, ff, err := ReadTOPString(s); err != nil {
		t.Fatalf("could not parse the string")
	} else {
		ats := ff.AtomTypes()
		checkLength(t, len(ats), 2, "AtomTypes()")
		compareF64(t, ats[0].LJDist(frc.FF_GROMACS), 3.72395664183e-01, "LJDist()")
		compareF64(t, ats[1].LJEnergy(frc.FF_GROMACS), 1.506240e-01, "LJEnergy()")
	}
}

func TestPairTypes(t *testing.T) {
	s := `
    [ defaults ]
    [ pairtypes ]
    ; i j   func    sigma1-4    epsilon1-4
    CEL1    CRL1     1  3.55468588538e-01  1.09105371453e-01 
    CEL1    CRL2     1  3.55468588538e-01  1.09105371453e-01 
    `

	if _, ff, err := ReadTOPString(s); err != nil {
		t.Fatalf("could not parse the string")
	} else {
		pts := ff.PairTypes()
		checkLength(t, len(pts), 2, "PairTypes()")
		compareF64(t, pts[0].LJDist14(frc.FF_GROMACS), 3.55468588538e-01, "LJDist14()")
		compareF64(t, pts[1].LJEnergy14(frc.FF_GROMACS), 1.09105371453e-01, "LJEnergy14()")
	}
}

func TestNonBondedTypes(t *testing.T) {
	s := `
    [ defaults ]
    [ nonbond_params ]
    ; i j   func    sigma   epsilon
    CRL1    CTL1     1  3.57250385974e-01  1.15060000000e-01 
    CRL1    HAL1     1  2.98451070577e-01  1.22591200000e-01 
    `

	if _, ff, err := ReadTOPString(s); err != nil {
		t.Fatalf("could not parse the string")
	} else {
		nbs := ff.NonBondedTypes()
		checkLength(t, len(nbs), 2, "NondBondedTypes()")
		compareF64(t, nbs[0].LJDist(frc.FF_GROMACS), 3.57250385974e-01, "LJDist()")
		compareF64(t, nbs[1].LJEnergy(frc.FF_GROMACS), 1.22591200000e-01, "LJEnergy()")

	}
}

func TestBondTypes(t *testing.T) {
	s := `
    [ defaults ]
    [ bondtypes ]
    ; i j   func    b0  Kb
    CEL1    CEL1     1  1.340000e-01  3.681920e+05
    CEL1    CRL1     1  1.502000e-01  2.008320e+05
    `

	if _, ff, err := ReadTOPString(s); err != nil {
		t.Fatalf("could not parse the string")
	} else {
		bts := ff.BondTypes()
		checkLength(t, len(bts), 2, "BondTypes()")
		compareF64(t, bts[0].HarmonicConstant(frc.FF_GROMACS), 3.681920e+05, "HarmonicConstant()")
		compareF64(t, bts[1].HarmonicDistance(frc.FF_GROMACS), 1.502000e-01, "HarmonicDistance()")
	}
}

func TestAngleTypes(t *testing.T) {
	s := `
    [ defaults ]
    [ angletypes ]
    ; i j   k   func    th0 cth S0  Kub
    CEL1    CEL1    CEL1     5  1.2350000e+02  4.0166400e+02  0.0000000e+00  0.0000000e+00
    CEL1    CEL1    CRL1     5  1.2350000e+02  4.0166400e+02  0.0000001e+00  0.0000002e+00
    `

	if _, ff, err := ReadTOPString(s); err != nil {
		t.Fatalf("could not parse the string")
	} else {
		ags := ff.AngleTypes()
		checkLength(t, len(ags), 2, "AngleTypes()")
		compareF64(t, ags[0].ThetaConstant(frc.FF_GROMACS), 4.0166400e+02, "ThetaConstant()")
		compareF64(t, ags[0].Theta(frc.FF_GROMACS), 1.2350000e+02, "Theta()")
		compareF64(t, ags[1].UBConstant(frc.FF_GROMACS), 0.0000002e+00, "UBConstant()")
		compareF64(t, ags[1].R13(frc.FF_GROMACS), 0.0000001e+00, "R13()")

	}
}

func TestDihedralTypes(t *testing.T) {
	s := `
    [ defaults ]
    [ dihedraltypes ]
    ; i j   k   l   func    phi0    cp  mult
    CEL1    CEL1    CRL1    CRL1     9  1.800000e+02  2.092000e+00      1
    CEL1    CEL1    CRL1    CRL1     9  1.800000e+02  5.439200e+00      3

    [ dihedraltypes ]
    ; i j   k   l   func    q0  cq
    OBL       X       X      CL     2  0.000001e+00  8.368000e+02
    `

	if _, ff, err := ReadTOPString(s); err != nil {
		t.Fatalf("could not parse the string")
	} else {
		dhs := ff.DihedralTypes()
		checkLength(t, len(dhs), 2, "DihedralTypes()")
		compareF64(t, dhs[0].PhiConstant(frc.FF_GROMACS), 2.092000e+00, "PhiConstant()")
		compareF64(t, dhs[1].Phi(frc.FF_GROMACS), 1.800000e+02, "Phi()")
		compareInt8(t, dhs[1].Mult(frc.FF_GROMACS), 3, "Mult()")

		// impropertypes
		ims := ff.ImproperTypes()
		checkLength(t, len(ims), 1, "ImproperTypes()")
		compareF64(t, ims[0].PsiConstant(frc.FF_GROMACS), 8.368000e+02, "PsiConstant()")
		compareF64(t, ims[0].Psi(frc.FF_GROMACS), 0.000001e+00, "Psi()")
	}
}

func TestTopPolymer(t *testing.T) {
	s := `
    [ moleculetype ]
    ; name  nrexcl
    LIG1         3

    [ moleculetype ]
    ; name  nrexcl
    LIG2         3
    `

	if top, _, err := ReadTOPString(s); err != nil {
		t.Fatalf("could not parse the string")
	} else {
		pols := top.RegisteredTopPolymers()
		checkLength(t, len(pols), 2, "RegisteredTopPolymers()")
		compareString(t, pols["LIG1"].Name(), "LIG1", "Name()")

	}
}

func TestTopAtoms(t *testing.T) {
	if top, _, err := ReadTOPString(topString1); err != nil {
		t.Fatalf("could not parse the string")
	} else {
		pol := top.TopPolymerByName("LIG1")
		checkLength(t, len(pol.TopAtoms()), 8, "TopAtoms()")

	}
}

func TestTopPairs(t *testing.T) {
	if top, _, err := ReadTOPString(topString1); err != nil {
		t.Fatalf("could not parse the string")
	} else {
		pol := top.TopPolymerByName("LIG1")
		checkLength(t, len(pol.TopPairs()), 2, "LIG1 TopPairs()")
	}
}

func TestTopBonds(t *testing.T) {
	if top, _, err := ReadTOPString(topString1); err != nil {
		t.Fatalf("could not parse the string")
	} else {
		pol := top.TopPolymerByName("LIG1")
		checkLength(t, len(pol.TopBonds()), 3, "LIG1 TopBonds()")
	}
}

func TestTopAngles(t *testing.T) {
	if top, _, err := ReadTOPString(topString1); err != nil {
		t.Fatalf("could not parse the string")
	} else {
		pol := top.TopPolymerByName("LIG1")
		checkLength(t, len(pol.TopAngles()), 2, "LIG1 TopAngles()")
	}
}

func TestTopDihedrals(t *testing.T) {
	if top, _, err := ReadTOPString(topString1); err != nil {
		t.Fatalf("could not parse the string")
	} else {
		pol := top.TopPolymerByName("LIG1")
		checkLength(t, len(pol.TopDihedrals()), 2, "LIG1 TopDihedrals()")
	}
}

func TestTopImpropers(t *testing.T) {
	if top, _, err := ReadTOPString(topString1); err != nil {
		t.Fatalf("could not parse the string")
	} else {
		pol := top.TopPolymerByName("LIG1")
		checkLength(t, len(pol.TopImpropers()), 1, "LIG1 TopImpropers()")
	}
}

func TestReadTop(t *testing.T) {
	_, _, err := ReadTOPFile("../../testdata/tmp.top")
	if err != nil {
		t.Errorf("could not read top file: %s", err)
	}

}

var topString1 string = `

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

;[ position_restraints ]
;   8     1      0.0      0.0  1000.0

;[ dihedral_restraints ]
;   5    1    8    3     1    120.0      2.5  1000.0
    
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
