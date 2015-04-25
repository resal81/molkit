package chm

import (
	"testing"

	"github.com/resal81/molkit/blocks"
)

func TestReadTOPFiles(t *testing.T) {
	fnames := []string{
		"../../testdata/chm_top/top_all22_prot.rtf",
		"../../testdata/chm_top/top_all35_ethers.rtf",
		"../../testdata/chm_top/top_all36_prot.rtf",
		"../../testdata/chm_top/top_all36_na.rtf",
		"../../testdata/chm_top/top_all36_lipid.rtf",
		"../../testdata/chm_top/top_all36_carb.rtf",
	}
	_, err := ReadTOPFiles(fnames...)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestReadTOP(t *testing.T) {
	var vals = []struct {
		resname                            string
		natoms, nbonds, nimpropers, ncmaps int
		hasln_next, hasln_prev             bool
	}{
		{"ASN", 14, 13, 4, 0, true, true},
		{"OTH", 17, 16, 4, 1, false, false},
	}

	ff, err := ReadTOPString(top_string)
	if err != nil {
		t.Fatalf("%s", err)
	}

	if v := len(ff.Fragments()); v != len(vals) {
		t.Errorf("wrong number of fragments => %d, expected %d", v, len(vals))
	}

	for _, el := range vals {
		f := ff.Fragment(blocks.HashKey(el.resname))
		if f == nil {
			t.Errorf("fragment is nil => %s", el.resname)
		}

		if v := len(f.Atoms()); v != el.natoms {
			t.Errorf("wrong number of atoms => %d, expected %d", v, el.natoms)
		}

		if v := len(f.Bonds()); v != el.nbonds {
			t.Errorf("wrong number of bonds => %d, expected %d", v, el.nbonds)
		}

		if v := len(f.Impropers()); v != el.nimpropers {
			t.Errorf("wrong number of impropers => %d, expected %d", v, el.nimpropers)
		}

		if v := len(f.CMaps()); v != el.ncmaps {
			t.Errorf("wrong number of cmaps => %v, expected %v", v, el.ncmaps)
		}

		if v := f.HasLinkerNext(); v != el.hasln_next {
			t.Errorf("mismatch linker_next => %v, expected %v", v, el.hasln_next)
		}

		if v := f.HasLinkerPrev(); v != el.hasln_prev {
			t.Errorf("mismatch linker_prev => %v, expected %v", v, el.hasln_prev)
		}
	}
}

var top_string = `
RESI ASN          0.00
GROUP   
ATOM N    NH1    -0.47  !     |       
ATOM HN   H       0.31  !  HN-N       
ATOM CA   CT1     0.07  !     |   HB1 OD1    HD21 (cis to OD1)
ATOM HA   HB      0.09  !     |   |   ||    /
GROUP                   !  HA-CA--CB--CG--ND2
ATOM CB   CT2    -0.18  !     |   |         \
ATOM HB1  HA      0.09  !     |   HB2        HD22 (trans to OD1)
ATOM HB2  HA      0.09  !   O=C           
GROUP                   !     |           
ATOM CG   CC      0.55
ATOM OD1  O      -0.55
GROUP   
ATOM ND2  NH2    -0.62
ATOM HD21 H       0.32
ATOM HD22 H       0.30
GROUP   
ATOM C    C       0.51
ATOM O    O      -0.51
BOND CB CA  CG CB   ND2 CG   
BOND N  HN  N  CA   C   CA    C +N   
BOND CA HA  CB HB1  CB  HB2  ND2 HD21  ND2 HD22 
DOUBLE C  O   CG  OD1  
IMPR N   -C  CA   HN    C   CA +N   O   
IMPR CG  ND2 CB   OD1   CG  CB ND2  OD1   
IMPR ND2 CG  HD21 HD22  ND2 CG HD22 HD21   
CMAP -C  N  CA  C   N  CA  C  +N
DONOR HN N   
DONOR HD21 ND2   
DONOR HD22 ND2   
ACCEPTOR OD1 CG   
ACCEPTOR O C   


RESI OTH          0.00
GROUP   
ATOM N    NH1    -0.47  
ATOM HN   H       0.31  
ATOM CA   CT1     0.07  
ATOM HA   HB      0.09  
GROUP                   
ATOM CB   CT2    -0.18  
ATOM HB1  HA      0.09  
ATOM HB2  HA      0.09  
GROUP                   
ATOM CG   CT2    -0.18
ATOM HG1  HA      0.09
ATOM HG2  HA      0.09
GROUP   
ATOM CD   CC      0.55
ATOM OE1  O      -0.55
GROUP   
ATOM NE2  NH2    -0.62
ATOM HE21 H       0.32
ATOM HE22 H       0.30
GROUP   
ATOM C    C       0.51
ATOM O    O      -0.51
BOND CB CA  CG  CB   CD  CG   NE2 CD   
BOND N  HN  N   CA   C   CA   
BOND CA  HA   CB  HB1  CB  HB2  CG HG1   
BOND CG HG2 NE2 HE21 NE2 HE22   
DOUBLE O  C    CD  OE1  
IMPR CD  NE2 CG   OE1   CD  CG NE2  OE1   
IMPR NE2 CD  HE21 HE22  NE2 CD HE22 HE21   
CMAP C  N  CA  C   N  CA  C  N
`
