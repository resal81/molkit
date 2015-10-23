package pdb

import (
	"testing"
)

func TestReadPdb(t *testing.T) {
	var data = []struct {
		path      string
		natoms    int
		nresidues int
		nchains   int
		nframes   int
	}{
		{"./testdata/nmr.pdb", 1340, 87, 1, 3},
		{"./testdata/dna.pdb", 1332, 380, 4, 1},
		{"./testdata/vdac.pdb", 2207, 288, 1, 1},
	}

	for _, d := range data {
		st, err := ReadPdbFile(d.path)
		if err != nil {
			t.Fatalf("Could not read pdb file: %s", err)
		}

		atoms := st.Atoms()
		if n := len(atoms); n != d.natoms {
			t.Errorf("Wrong # of atoms for '%s': got '%d', expected '%d'.", d.path, n, d.natoms)
		}

		at0 := atoms[0]
		if n := len(at0.Coords()); n != d.nframes {
			t.Errorf("Wrong # of frames for '%s': got '%d', expected '%d'.", d.path, n, d.nframes)
		}

		residues := st.Residues()
		if n := len(residues); n != d.nresidues {
			t.Errorf("Wrong # of residues for '%s': got '%d', expected '%d'.", d.path, n, d.nresidues)
		}

		chains := st.Chains()
		if n := len(chains); n != d.nchains {
			t.Errorf("Wrong # of chains for '%s': got '%d', expected '%d'.", d.path, n, d.nchains)
		}
	}

}
