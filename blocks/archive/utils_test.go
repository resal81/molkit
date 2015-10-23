package blocks

import (
	"testing"
)

func TestAtomNameToProtons(t *testing.T) {
	var vals = []struct {
		name    string
		protons int
	}{
		// basic
		{"H", 1}, {"Li", 3}, {"C", 6}, {"N", 7}, {"O", 8}, {"F", 9}, {"Na", 11}, {"Mg", 12}, {"AL", 13},
		{"P", 15}, {"S", 16}, {"Cl", 17}, {"K", 19}, {"Fe", 26}, {"Zn", 30}, {"Br", 35}, {"I", 53},

		// hydrogens
		{"HE2", 1}, {"HA1", 1}, {"HGA1", 1}, {"HGAAM0", 1}, {"22H21", 1},
		{"HGTIP3", 1},

		// carbons
		{"CT", 6}, {"CT1", 6}, {"CPH1", 6}, {"CAI", 6}, {"CG2R67", 6}, {"CG3AM0", 6},
		{"CG25C1", 6},

		// N
		{"NP", 7}, {"NH2", 7}, {"NC2", 7},

		// O
		{"OH1", 8}, {"OB", 8}, {"OG25C1", 8},

		// F
		{"FGA1", 9},

		// P
		{"PG0", 15},

		// S

		// Cl

		// Br
		{"BRGA3", 35},

		// I
		{"IGR1", 53},
	}

	for _, el := range vals {
		p, err := AtomNameToProtons(el.name)
		if err != nil {
			t.Errorf("%s", err)
			continue
		}
		if p != el.protons {
			t.Errorf("wrong protons => for name %s is %d, expected %d", el.name, p, el.protons)
		}
	}

}
