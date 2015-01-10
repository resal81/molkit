package pdb

import (
	"fmt"
	"testing"
)

type pdbData struct {
	path      string
	natoms    int
	nresidues int
	nchains   int
	nmodels   int
}

var testData []pdbData = []pdbData{
	pdbData{
		path:      "../../testdata/pdb/2zta.pdb",
		natoms:    576,
		nresidues: 116,
		nchains:   2,
		nmodels:   1,
	},
	pdbData{
		path:      "../../testdata/pdb/2mjx.pdb",
		natoms:    884,
		nresidues: 28,
		nchains:   2,
		nmodels:   8,
	},
}

func TestParsePDBFile(t *testing.T) {

	for _, pd := range testData {
		sys, err := ReadPDBFile(pd.path)
		if err != nil {
			t.Error(fmt.Sprintf("%s", err))
		}

		atoms := sys.Atoms()
		if len(atoms) != pd.natoms {
			t.Errorf("natoms -> should be %d, is %d", pd.natoms, len(atoms))
		}

		if len(sys.Fragments()) != pd.nresidues {
			t.Errorf("nresidues -> should be %d, is %d", pd.nresidues, len(sys.Fragments()))
		}

		if len(sys.Polymers()) == pd.nmodels {
			t.Errorf("nmodels -> should be %d, is %d", pd.nmodels, len(sys.Polymers()))
		}

	}

}
