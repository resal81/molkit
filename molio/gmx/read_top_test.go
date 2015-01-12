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
   CEL1    CEL1    CRL1     5  1.2350000e+02  4.0166400e+02  0.0000000e+00  0.0000000e+00

[ dihedraltypes ]
; i j   k   l   func    phi0    cp  mult
   CEL1    CEL1    CRL1    CRL1     9  1.800000e+02  2.092000e+00      1
   CEL1    CEL1    CRL1    CRL1     9  1.800000e+02  5.439200e+00      3

[ dihedraltypes ]
; i j   k   l   func    q0  cq
    OBL       X       X      CL     2  0.000000e+00  8.368000e+02
`

func TestCase1(t *testing.T) {
	_, ff, err := ReadTOPString(testCase1_string)
	if err != nil {
		t.Error("could not parse testCase1_string")
	}

	if len(ff.AtomTypes()) != 2 {
		t.Errorf("atomtype length is not 2 and is %d", len(ff.AtomTypes()))
	}
}

func TestReadTop(t *testing.T) {
	_, _, err := ReadTOPFile("../../testdata/tmp.top")
	if err != nil {
		t.Errorf("could not read top file: %s", err)
	}
}
