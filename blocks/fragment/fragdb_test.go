package fragment

import (
	"testing"
)

// ***********************************************************************//
// Helpers Tests
// ***********************************************************************//

func TestStripComment(t *testing.T) {
	var data = []struct {
		lineIn  string
		lineOut string
	}{
		{"; a full line comment", ""},
		{";", ""},
		{"[ bond ] ; some info ", "[ bond ] "},
	}

	for _, d := range data {
		if out := stripComment(d.lineIn); out != d.lineOut {
			t.Errorf("Wrong stripComment result: got '%s', expected '%s'", out, d.lineOut)
		}
	}
}

func TestGetBracketField(t *testing.T) {
	var data = []struct {
		line   string
		result string
	}{
		{" [ bond ] ; with comment and extra [ bracket ] ", "bond"},
		{" no bracket ", " no bracket "},
	}

	for _, d := range data {
		if out := getBracketField(d.line); out != d.result {
			t.Errorf("Wrong getBracketField result: got '%s', expected '%s'", out, d.result)
		}
	}
}

// ***********************************************************************//
// RTP parse test
// ***********************************************************************//

func TestParseRtpContent(t *testing.T) {
	fdb := ParseRtpContent(sampleRtpContent)

	if fdb.Size() != 9 {
		t.Errorf("Wrong FragmentDatabase size: got `%d`, expected `%d`", fdb.Size(), 9)
	}

	if fdb.GetFragmentByName("XX") != nil {
		t.Error("Fragments that are not found must be nil")
	}

	URE := fdb.GetFragmentByName("URE")
	atoms := URE.Atoms()
	bonds := URE.Bonds()

	if len(atoms) != 8 {
		t.Errorf("Expected 8 atoms in fragment URE, got %d", len(atoms))
	}

	if len(bonds) != 7 {
		t.Errorf("Expected 7 bonds in fragment URE, got %d", len(bonds))
	}

	if ch := atoms[0].charge; ch != 0.880229 {
		t.Errorf("Atom charge is wrong: got %f", ch)
	}
}

// ***********************************************************************//
// Sample RTP
// ***********************************************************************//

const sampleRtpContent = `
[ LI ]
 [ atoms ]
   LI     Li           1.00000     1 

[ ZN ]
 [ atoms ]
   ZN     Zn           2.00000     1

[ URE ] ; urea added in by EJS, resp charges by Jim Caldwell
 [ atoms ]
    C      C            0.880229    1   
    O      O           -0.613359    2   
   N1      N           -0.923545    3   
  H11      H            0.395055    4   
  H12      H            0.395055    5   
   N2      N           -0.923545    6   
  H21      H            0.395055    7   
  H22      H            0.395055    8   
 [ bonds ]
    C     N1
    C     N2
    C      O
   N1    H11
   N1    H12
   N2    H21
   N2    H22
 [ impropers ]
    N1    N2     C     O
     C   H11    N1   H12
     C   H21    N2   H22    

[ ACE ]
 [ atoms ]
  HH31    HC           0.11230     1
   CH3    CT          -0.36620     2
  HH32    HC           0.11230     3
  HH33    HC           0.11230     4
     C    C            0.59720     5
     O    O           -0.56790     6
 [ bonds ]
  HH31   CH3
   CH3  HH32
   CH3  HH33
   CH3     C
     C     O
 [ impropers ]
   CH3    +N     C     O
                        
[ NME ] 
 [ atoms ]
     N    N           -0.41570     1
     H    H            0.27190     2
   CH3    CT          -0.14900     3
  HH31    H1           0.09760     4
  HH32    H1           0.09760     5
  HH33    H1           0.09760     6
 [ bonds ]
     N     H
     N   CH3
   CH3  HH31
   CH3  HH32
   CH3  HH33
    -C     N
 [ impropers ]
    -C   CH3     N     H
                        
[ NHE ]
 [ atoms ]
     N    N           -0.46300     1
    H1    H            0.23150     2
    H2    H            0.23150     3
 [ bonds ]
     N    H1
     N    H2
    -C     N
 [ impropers ]
    -C    H1     N    H2

[ NH2 ]
 [ atoms ]
     N    N           -0.46300     1
    H1    H            0.23150     2
    H2    H            0.23150     3
 [ bonds ]
     N    H1
     N    H2
    -C     N
 [ impropers ]
    -C    H1     N    H2

; Next are non-terminal AA's

[ ALA ]
 [ atoms ]
     N    N           -0.41570     1
     H    H            0.27190     2
    CA    CT           0.03370     3
    HA    H1           0.08230     4
    CB    CT          -0.18250     5
   HB1    HC           0.06030     6
   HB2    HC           0.06030     7
   HB3    HC           0.06030     8
     C    C            0.59730     9
     O    O           -0.56790    10
 [ bonds ]
     N     H
     N    CA
    CA    HA
    CA    CB
    CA     C
    CB   HB1
    CB   HB2
    CB   HB3
     C     O
    -C     N
 [ impropers ]
    -C    CA     N     H
    CA    +N     C     O
                        
[ GLY ]
 [ atoms ]
     N    N           -0.41570     1
     H    H            0.27190     2
    CA    CT          -0.02520     3
   HA1    H1           0.06980     4
   HA2    H1           0.06980     5
     C    C            0.59730     6
     O    O           -0.56790     7
 [ bonds ]
     N     H
     N    CA
    CA   HA1
    CA   HA2
    CA     C
     C     O
    -C     N
 [ impropers ]
    -C    CA     N     H
    CA    +N     C     O

`
